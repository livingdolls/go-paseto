package http2

import (
	"github.com/livingdolls/go-paseto/internal/core/common/router"
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"
	"github.com/livingdolls/go-paseto/internal/core/port/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	gin         *gin.Engine
	userService service.UserPortService
}

func NewUserController(gin *gin.Engine, userService service.UserPortService) UserController {
	return UserController{
		gin:         gin,
		userService: userService,
	}
}

func (u UserController) InitRouter(v string) {
	api := u.gin.Group(v)

	router.Post(api, "/register", u.register)
	router.Get(api, "/users", u.getUsers)
	api.GET("/user/:id", u.getUserById)
	router.Get(api, "/login", u.login)
}

func (u UserController) register(c *gin.Context) {
	var req request.RegisterUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.RequestValidationError(c, err)
		return
	}

	resp, err := u.userService.Register(&req)

	if err != nil {
		response.HandleErrorResponse(c, err)
		return
	}

	response.HandleSuccessResponseCreated(c, resp)
}

func (u UserController) getUsers(c *gin.Context) {
	res, err := u.userService.ListUsers()

	if err != nil {
		response.HandleErrorResponse(c, err)
		return
	}

	response.HandleSuccessResponse(c, res)
}

func (u UserController) getUserById(c *gin.Context) {
	var getId request.GetUserByIdRequest

	if err := c.ShouldBindUri(&getId); err != nil {
		response.RequestValidationError(c, err)
		return
	}

	res, err := u.userService.GetUser(&getId)

	if err != nil {
		response.HandleErrorResponse(c, err)
		return
	}

	response.HandleSuccessResponse(c, res)
}

func (u UserController) login(c *gin.Context) {
	var user request.LoginUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		response.RequestValidationError(c, err)
		return
	}

	res, err := u.userService.Login(&user)

	if err != nil {
		response.HandleErrorResponse(c, err)
		return
	}

	response.HandleSuccessResponse(c, res)
}
