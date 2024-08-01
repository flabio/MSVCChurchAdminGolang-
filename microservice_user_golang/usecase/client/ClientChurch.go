package client

import (
	"net/http"
	"strconv"

	"microservice_user.com/infrastructure/utils"
)

func MsvcChurchById(churchId uint) string {
	data, err := http.NewRequest(utils.GET, utils.CHURCH_MSVC_URL+strconv.Itoa(int(churchId)), nil)
	msg := ReponseClient(data, err)
	return msg
}
