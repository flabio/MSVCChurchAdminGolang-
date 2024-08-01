package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReponseClient(data *http.Request, err error) string {
	var msg string = ""
	if err != nil {
		msg = err.Error()
	}
	client := &http.Client{}

	resp, err := client.Do(data)
	if err != nil {
		msg = err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var findByIdRol map[string]interface{}
	errJson := json.Unmarshal([]byte(body), &findByIdRol)

	if errJson != nil {
		msg = errJson.Error()
	}
	return msg
}
