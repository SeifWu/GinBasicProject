package v1public

import (
	"net/http"
	"reflect"
	param "seifwu/app/params"
	service "seifwu/app/services"
	"seifwu/global/response"
	util "seifwu/utils"

	"github.com/gin-gonic/gin"
)

var signUpParam param.SignUpParam

// SignUp 注册
func SignUp(c *gin.Context) {

	if err := c.ShouldBind(&signUpParam); err != nil {
		reflects := reflect.TypeOf(signUpParam)
		util.UnifiedValidation(c, err, "40002", reflects)

		return
	}

	result := service.UserCreateService(c, signUpParam)
	if result["success"] == false {
		response.Response(c, http.StatusBadRequest, 400, nil, result["message"])
		return
	}

	response.Success(c, gin.H{"user": result["data"]}, result["message"])
}