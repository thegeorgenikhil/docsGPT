package pinecone

const (
	Namespace = "docsGPT"
)

type Vector struct {
	ID       string            `json:"id"`
	Values   []float32         `json:"values"`
	Metadata map[string]string `json:"metadata"`
}

type Match struct {
	ID    string `json:"id"`
	Score string `json:"score"`
}

type UpsertRequest struct {
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace"`
}

type UpsertResult struct {
	UpsertedCount int `json:"upsertedCount"`
}

type QueryRequest struct {
	Filter          map[string]string `json:"filter"`
	IncludeValues   bool              `json:"includeValues"`
	IncludeMetadata bool              `json:"includeMetadata"`
	Vector          []float32         `json:"vector"`
	TopK            int               `json:"topK"`
	Namespace       string            `json:"namespace"`
}

type QueryResult struct {
	Matches   []Match `json:"matches"`
	Namespace string  `json:"namespace"`
}
