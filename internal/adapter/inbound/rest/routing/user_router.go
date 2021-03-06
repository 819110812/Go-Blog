package routing

import (
	"Go-Blog/internal/adapter/inbound/rest/controller"
	"Go-Blog/internal/app/usecase"
	"Go-Blog/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

var (
	baseController  = controller.BaseController{}
	loginController = controller.UserController{
		UserUseCase: usecase.UserUseCase{
			UserRepository: &service.UserService{},
		},
	}
)

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("")
	{
		// TODO: 后期改成通过uuid查询用户
		userRouter.GET("/user/:id", loginController.UserInfoById)
		userRouter.POST("/login", loginController.UserLogin)
		userRouter.POST("/loginAdmin", loginController.UserLogin)
		userRouter.POST("/register", loginController.UserRegister)

		//
		userRouter.POST("/login/account", loginController.AdminLogin)
		userRouter.GET("/currentUser", loginController.GetCurrentUser)
		userRouter.POST("/login/outLogin", loginController.OutLogin)
	}
}
