package teams

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"microservice_user.com/usecase/client"
)

func ReponseTeamClient(data *http.Request, err error) (Team, string) {
	var team Team
	var msg string = ""
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
	body, err := ioutil.ReadAll(resp.Body)
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(body), &dataMap)

	if errJson != nil {
		msg = errJson.Error()
	}
	client.MapToStruct(dataMap, &team)
	return team, msg
}
