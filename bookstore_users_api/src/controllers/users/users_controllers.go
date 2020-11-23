package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../domain/users"
	"../../services/user_services"
	"../../utils/errors"
	// "fmt"
	"strconv"
	"../../services"
	"../login"
)

var (
	loginService    service.LoginService       = service.StaticLoginService()
	jwtService      service.JWTService         = service.JWTAuthService()
	loginController 						   = login.LoginHandler(loginService, jwtService)
)

func GetUser(c *gin.Context){
	key := "user_id"
	base := 10
	bitSize := 64
	userId, userErr := strconv.ParseInt(c.Param(key), base, bitSize)
	if userErr != nil {
		message := "User id should be a number"
		err := errors.NewBadrequestError(message)
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}	
	c.JSON(http.StatusOK , user)
}

func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil{
		message := "invalid json body"
		restErr := errors.NewBadrequestError(message)
		c.JSON(restErr.Status, restErr)
		//TODO: handle json error
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handleuser creation error
		return
	}	
	c.JSON(http.StatusCreated, result)
}

func LoginUser(ctx *gin.Context) {
	token := loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}
}


// func SearchUser(c *gin.Context){
// 	messages := "Implement data"	
// 	c.String(http.StatusNotImplemented, messages)
// }