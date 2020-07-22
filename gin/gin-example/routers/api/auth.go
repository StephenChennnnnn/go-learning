package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/stephenchen/go-learning/gin/gin-example/models"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/e"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/logging"
	"github.com/stephenchen/go-learning/gin/gin-example/pkg/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok, _ := valid.Valid(&a); ok {
		if isExist := models.CheckAuth(username, password); isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
