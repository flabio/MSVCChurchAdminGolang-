package ministerial

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ResponseUserMinisterialClient(data *http.Request, err error) (string, string) {
	var msg string = ""
	//var userMinisterial UserMinisterial
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
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

	var jsonMap map[string]interface{}
	errJson := json.Unmarshal([]byte(body), &jsonMap)

	if errJson != nil {
		msg = errJson.Error()
	}
	log.Println(jsonMap)
	return string(body), msg
}
