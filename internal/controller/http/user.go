package http2

import (
	"gopaseto/internal/core/common/router"
	"gopaseto/internal/core/model/request"
	"gopaseto/internal/core/model/response"
	"gopaseto/internal/core/port/service"

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

func (u UserController) InitRouter() {
	api := u.gin.Group("api/v1")

	router.Post(api, "/register", u.register)
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
