package jsonrpc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
)

const bodyType string = "application/json"

func Call(url string, params map[string]interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method": "call",
		"id": "id",
		"params": params,
	})
	if err != nil {
		log.Fatalf("Marshal: %v", err)
		return nil, err
	}

	resp, err := http.Post(url, bodyType, bytes.NewBuffer(data))

	if err != nil {
		log.Fatalf("Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Response: \n%s \nUnmarshal: %v", string(body),err)
		return nil, err
	}

	return result, nil
}