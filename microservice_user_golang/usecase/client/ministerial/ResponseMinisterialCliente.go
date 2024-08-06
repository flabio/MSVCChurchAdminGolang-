package ministerial

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"microservice_user.com/usecase/client"
)

func ResponseUserMinisterialClient(data *http.Request, err error) ([]Ministerial, string) {
	var msg string = ""
	var ministerial Ministerial
	var ministerials []Ministerial

	if err != nil {
		msg = err.Error()
	}

	clientHttp := &http.Client{}
	resp, err := clientHttp.Do(data)
	if err != nil {
		msg = err.Error()
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var sliceData []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &sliceData); err == nil {
		for _, item := range sliceData {
			client.MapToStruct(item, &ministerial)
			ministerials = append(ministerials, ministerial)
		}
	}
	if err != nil {
		msg = err.Error()
	}
	return ministerials, msg
}
