package controller

import (
	"Go-Blog/internal/adapter/inbound/rest/controller/dto/request"
	"Go-Blog/internal/adapter/inbound/rest/controller/dto/response"
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/constant"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

type UserController struct {
	BaseController
	UserUseCase usecase.UserUseCase
}

// UserLogin 博客端 普通用户登陆
func (u *UserController) UserLogin(c *gin.Context) {
	req := request.UserLoginRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		response.InvalidParamWithMsg(http.StatusBadRequest, constant.InvalidParams, c)
	}
	user, err := u.UserUseCase.Login(&req)
	if err != nil {
		log.Println(err)
		response.Error(http.StatusUnauthorized, c)
		return
	}
	log.Println(req)
	response.SuccessWithData(user, c)
}

// AdminLogin 管理端 管理员登陆
func (u *UserController) AdminLogin(c *gin.Context) {
	req := request.AdminLoginRequest{}
	err := c.BindJSON(&req)
	if err != nil || req.Account == "" {
		response.InvalidParamWithMsg(http.StatusBadRequest, constant.LoginErrorInvalidParams.Message, c)
		return
	}
	user, err := u.UserUseCase.AdminLogin(&req)

	if err != nil {
		log.Println(err)
		response.InvalidParamWithMsg(http.StatusUnauthorized, err.Error(), c)
		return
	}

	log.Println(req)
	response.SuccessWithData(user, c)
}

// GetCurrentUser
func (u *UserController) GetCurrentUser(c *gin.Context) {
	response.SuccessWithData(nil, c)
}

// OutLogin
func (u *UserController) OutLogin(c *gin.Context) {
	response.SuccessWithData(nil, c)
}

func (u *UserController) UserRegister(c *gin.Context) {
	register := request.UserRegisterRequest{}
	err := c.BindJSON(&register)
	if err != nil {
		response.InvalidParamWithMsg(http.StatusBadRequest, constant.InvalidParams, c)
		log.Panicln(err)
	}
	log.Println("request from front end", c.Request.Body)
	res, err := u.UserUseCase.Register(&register)
	if err != nil {
		log.Println(err)
		response.Error(http.StatusInternalServerError, c)
		return
	}
	response.SuccessWithData(res, c)
}

// UserInfoById TODO: 后期删除该方法，通过UUID获取用户信息
func (u *UserController) UserInfoById(c *gin.Context) {
	userpo, err := u.UserUseCase.GetUserById(uuid.FromStringOrNil(c.Param("id")))
	if err != nil {
		response.Error(http.StatusInternalServerError, c)
		log.Panicln(err)
	}
	response.SuccessWithData(userpo, c)
}
