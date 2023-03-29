package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-openai"
	"github.com/thegeorgenikhil/docsGPT/models"
	"github.com/thegeorgenikhil/docsGPT/pkg/pinecone"
	"github.com/thegeorgenikhil/docsGPT/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func QueryFromEmbeddedDocument() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.QueryRequest
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
		PINECONE_API_KEY := os.Getenv("PINECONE_API_KEY")
		PINECONE_INDEX_URL := os.Getenv("PINECONE_INDEX_URL")

		if OPENAI_API_KEY == "" || PINECONE_API_KEY == "" || PINECONE_INDEX_URL == "" {
			log.Println("Environment variables are not set properly")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Server not configured properly"})
			return
		}

		// TODO:
		// 1. Get the Request Body
		err := c.BindJSON(&req)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while reading the request body"})
		}
		// 2. Check if a document exists with that ID
		documentId, err := primitive.ObjectIDFromHex(req.DocumentID)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Object Id"})
			return
		}

		count, err := DocumentCollection.CountDocuments(c, bson.M{"_id": documentId})

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not able to retrieve the documents"})
			return
		}

		if count != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No documents with the given id"})
			return
		}

		// 3. Embed the Request Query using embeddings api

		client := gogpt.NewClient(OPENAI_API_KEY)
		res, err := utils.GetEmbeddings(client, ctx, []string{req.Query})

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while embedding the user query"})
		}

		userQueryEmbedding := res.Data[0].Embedding

		// 4. Get Back the Response and query the vector db with the returned vector of user query along with document id as the filter (get the top 3 similar ones)
		queryRequest := &pinecone.QueryRequest{
			Filter:          map[string]string{},
			IncludeValues:   false,
			IncludeMetadata: false,
			Vector:          userQueryEmbedding,
			TopK:            3,
			Namespace:       pinecone.Namespace,
		}
		queryResult, err := pinecone.Query(PINECONE_API_KEY, PINECONE_INDEX_URL, queryRequest)

		// 5. Use the top 3 response and add it to chatgpt api along with user query and get back the result

		var results []primitive.ObjectID

		for _, result := range queryResult.Matches {
			contentId, _ := primitive.ObjectIDFromHex(result.ID)
			results = append(results, contentId)
		}

		curr, err := ContentCollection.Find(c, bson.M{"_id": bson.M{"$in": results}})

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting documents from database"})
		}

		var contents []models.Content

		for curr.Next(c) {
			var content models.Content
			err := curr.Decode(&content)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding the documents"})
				return
			}
			contents = append(contents, content)
		}

		// 6. Chatgpt api call
		prompt := utils.BuildPrompt(req.Query, contents)
		gptResponse, err := utils.GetGPTResponse(client, ctx, prompt)

		c.JSON(http.StatusOK, gin.H{"content": gptResponse})
	}
}
