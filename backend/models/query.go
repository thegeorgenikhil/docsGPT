package models

type QueryRequest struct {
	DocumentID string `json:"document_id"`
	Query      string `json:"query"`
}
