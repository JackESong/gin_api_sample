package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"gin_api_sample/common"
	"gin_api_sample/framework/auth"
	"gin_api_sample/project/dao"
)
type user struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := user{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := common.INVALID_PARAMS
	if ok {
		isExist := dao.CheckAuth(username, password)
		if isExist {
			token, err := auth.GenerateToken(username, password)
			if err != nil {
				code = common.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = common.SUCCESS
			}
		} else {
			code = common.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : common.GetMsg(code),
		"data" : data,
	})
}