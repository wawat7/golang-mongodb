package user

import (
	"github.com/gin-gonic/gin"
	"golang-mongodb/exception"
	"net/http"
)

type userController struct {
	service Service
}

func NewController(service Service) *userController {
	return &userController{service: service}
}

func (controller *userController) Route(app *gin.Engine) {
	route := app.Group("api/users")
	route.GET("/", controller.GetAll)
	route.POST("/", controller.Create)
	route.GET("/:id", controller.Get)
	route.PUT("/:id", controller.Update)
	route.DELETE("/:id", controller.Delete)
}

func (controller *userController) Create(c *gin.Context) {
	var input CreateUserRequest
	err := c.ShouldBindJSON(&input)
	exception.PanicIfNeeded(err)

	user := User{
		Name:  input.Name,
		Email: input.Email,
	}

	address := Address{
		Address:  input.Address,
		City:     input.City,
		Province: input.Province,
	}

	user = controller.service.Create(user, address)

	c.JSON(http.StatusOK, user)
	return
}

func (controller *userController) GetAll(c *gin.Context) {
	users := controller.service.FindAll()

	c.JSON(http.StatusOK, users)
	return
}

func (controller *userController) Get(c *gin.Context) {
	var param GetUserDetail
	err := c.ShouldBindUri(&param)
	exception.PanicIfNeeded(err)

	user := controller.service.FindById(param.Id)
	c.JSON(http.StatusOK, user)
	return
}

func (controller *userController) Update(c *gin.Context) {
	var param GetUserDetail
	err := c.ShouldBindUri(&param)
	exception.PanicIfNeeded(err)

	var input CreateUserRequest
	err = c.ShouldBindJSON(&input)
	exception.PanicIfNeeded(err)

	user := controller.service.Update(param.Id, input)
	c.JSON(http.StatusOK, user)
	return
}

func (controller *userController) Delete(c *gin.Context) {
	var param GetUserDetail
	err := c.ShouldBindUri(&param)
	exception.PanicIfNeeded(err)

	controller.service.Delete(param.Id)
	c.JSON(http.StatusOK, nil)
	return
}
