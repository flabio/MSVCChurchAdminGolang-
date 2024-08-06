package functionministerial

import (
	"net/http"
	"strconv"

	"microservice_user.com/infrastructure/utils"
)

func MsvcUserFunctionMInisterialFindById(id uint) []FunctionMinisterial {

	data, err := http.NewRequest(utils.GET, utils.FUNCTION_MINISTERIAL_USER_MSVC_URL+strconv.Itoa(int(id)), nil)
	result, _ := ResponseUserFunctionMinisterialClient(data, err)
	return result
}
