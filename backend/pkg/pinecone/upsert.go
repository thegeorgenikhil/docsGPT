package pinecone

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Upsert(apiKey string, indexUrl string, upsertReq *UpsertRequest) (*UpsertResult, error) {
	var upsertRes UpsertResult
	url := "https://" + indexUrl + "/vectors/upsert"

	upsertReqJSON, err := json.Marshal(upsertReq)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(upsertReqJSON))

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", apiKey)

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &upsertRes)

	return &upsertRes, err
}
