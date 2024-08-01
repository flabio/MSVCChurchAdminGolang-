package rol

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"microservice_user.com/usecase/client"
)

func ReponseRolClient(data *http.Request, err error) (Rol, string) {
	var msg string = ""
	var rol Rol

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
	client.MapToStruct(dataMap, &rol)
	return rol, msg
}
