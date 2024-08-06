package church

import (
	"log"
	"net/http"
	"strconv"

	"microservice_user.com/infrastructure/utils"
)

func MsvcChurchFindById(id uint) Church {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()
	data, err := http.NewRequest(utils.GET, utils.CHURCH_MSVC_URL+strconv.Itoa(int(id)), nil)
	result, _ := ReponseChurchClient(data, err)
	return result
}
