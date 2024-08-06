package teams

import (
	"log"
	"net/http"
	"strconv"

	"microservice_user.com/infrastructure/utils"
	"microservice_user.com/usecase/client"
)

func MsvcTeamById(id uint) string {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()
	data, err := http.NewRequest(utils.GET, utils.TEAM_MSVC_URL+strconv.Itoa(int(id)), nil)
	msg := client.ReponseClient(data, err)
	return msg
}
func MsvcTeamFindById(id uint) Team {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()
	data, err := http.NewRequest(utils.GET, utils.TEAM_MSVC_URL+strconv.Itoa(int(id)), nil)
	result, _ := ReponseTeamClient(data, err)
	return result
}
