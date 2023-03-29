package pinecone

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Query(apiKey string, indexUrl string, queryReq *QueryRequest) (*QueryResult, error) {
	var queryRes QueryResult
	url := "https://" + indexUrl + "/query"

	queryReqJSON, _ := json.Marshal(queryReq)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(queryReqJSON))

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &queryRes)

	return &queryRes, err
}
