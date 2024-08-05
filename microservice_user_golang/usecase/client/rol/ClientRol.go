package rol

import (
	"log"
	"net/http"
	"strconv"

	"microservice_user.com/infrastructure/utils"
	"microservice_user.com/usecase/client"
)

func MsvcRolById(rolId uint) string {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()
	data, err := http.NewRequest(utils.GET, utils.ROL_MSVC_URL+strconv.Itoa(int(rolId)), nil)
	msg := client.ReponseClient(data, err)
	return msg
}
func MsvcRolFindById(rolId uint) Rol {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()

	data, err := http.NewRequest(utils.GET, utils.ROL_MSVC_URL+strconv.Itoa(int(rolId)), nil)
	rol, _ := ReponseRolClient(data, err)
	return rol
}
