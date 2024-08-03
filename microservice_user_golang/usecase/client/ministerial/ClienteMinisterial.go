package ministerial

import (
	"net/http"
	"strconv"

	"microservice_user.com/infrastructure/utils"
)

func MsvcUserMInisterialFindById(id uint) string {

	data, err := http.NewRequest(utils.GET, utils.MINISTERIAL_USER_MSVC_URL+strconv.Itoa(int(id)), nil)
	result, _ := ResponseUserMinisterialClient(data, err)
	return result
}
