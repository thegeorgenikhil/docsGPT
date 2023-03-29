package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ledongthuc/pdf"
	gogpt "github.com/sashabaranov/go-openai"
	"github.com/thegeorgenikhil/docsGPT/database"
	"github.com/thegeorgenikhil/docsGPT/models"
	"github.com/thegeorgenikhil/docsGPT/pkg/pinecone"
	"github.com/thegeorgenikhil/docsGPT/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 10 // 10MB
var DocumentCollection *mongo.Collection = database.Client.Database(database.Database).Collection("documents")
var ContentCollection *mongo.Collection = database.Client.Database(database.Database).Collection("content")

func GetDocuments() gin.HandlerFunc {
	return func(c *gin.Context) {
		var documents []models.Document

		cursor, err := DocumentCollection.Find(c, bson.M{})
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not able to fetch documents"})
			return
		}

		if err = cursor.All(c, &documents); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not able to fetch documents"})
			return
		}

		if len(documents) == 0 {
			c.JSON(http.StatusOK, gin.H{"documents": []string{}})
			return
		}

		c.JSON(http.StatusOK, gin.H{"documents": documents})
	}
}

func UploadDocumentToEmbed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
		PINECONE_API_KEY := os.Getenv("PINECONE_API_KEY")
		PINECONE_INDEX_URL := os.Getenv("PINECONE_INDEX_URL")

		if OPENAI_API_KEY == "" || PINECONE_API_KEY == "" || PINECONE_INDEX_URL == "" {
			log.Println("Environment variables are not set properly")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Server not configured properly"})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			log.Println("File not uploaded properly")
			c.JSON(http.StatusBadRequest, gin.H{"error": "File not uploaded properly"})
			return
		}

		if file.Size > MAX_UPLOAD_SIZE {
			log.Println("File is too big")
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is too big"})
			return
		}

		if path.Ext(file.Filename) != ".pdf" {
			log.Println("File is not a pdf")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not a PDF file"})
			return
		}

		document := models.Document{
			ID:        primitive.NewObjectID(),
			Name:      file.Filename,
			Size:      file.Size,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err = DocumentCollection.InsertOne(c, document)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while saving the document metadata"})
			return
		}

		err = c.SaveUploadedFile(file, "pdf/"+document.ID.String()+file.Filename)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		f, r, err := pdf.Open(path.Join("pdf", document.ID.String()+file.Filename))
		defer func() {
			_ = f.Close()
		}()
		totalPage := r.NumPage()
		pageCounter := 1
		var rows []string
		currText := ""

		for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
			p := r.Page(pageIndex)
			if p.V.IsNull() {
				continue
			}
			content, _ := p.GetPlainText(make(map[string]*pdf.Font))
			currText += content
			if pageCounter == 2 {
				rows = append(rows, currText)
				currText = ""
				pageCounter = 0
			}
			pageCounter++
		}

		if currText != "" {
			rows = append(rows, currText)
		}

		client := gogpt.NewClient(OPENAI_API_KEY)
		res, err := utils.GetEmbeddings(client, ctx, rows)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating embeddings"})
			return
		}

		// will be using this for insertMany operation
		var contentArr []interface{}
		var vectors []pinecone.Vector

		for index, row := range rows {
			content := models.Content{
				ID:         primitive.NewObjectID(),
				DocumentId: document.ID,
				Content:    row,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}

			vector := pinecone.Vector{
				ID:     content.ID.Hex(),
				Values: res.Data[index].Embedding,
				Metadata: map[string]string{
					"document_id": document.ID.Hex(),
				},
			}

			contentArr = append(contentArr, content)
			vectors = append(vectors, vector)

			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating embeddings"})
				return
			}

		}

		upsertReq := &pinecone.UpsertRequest{
			Vectors:   vectors,
			Namespace: pinecone.Namespace,
		}

		// All the content is saved in the contentArr, Now bulk insert to mongo
		_, err = pinecone.Upsert(PINECONE_API_KEY, PINECONE_INDEX_URL, upsertReq)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save embeddings"})
			return
		}

		_, err = ContentCollection.InsertMany(c, contentArr)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save text to database"})
			return
		}

		err = os.Remove(path.Join("pdf", document.ID.String()+file.Filename))

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
	}
}
