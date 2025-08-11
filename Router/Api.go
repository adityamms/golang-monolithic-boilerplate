package Router

import (
	"github.com/gin-gonic/gin"
	Auth "github.com/mahdidl/golang_boilerplate/Components/Auth"
	User "github.com/mahdidl/golang_boilerplate/Components/User"
	"github.com/mahdidl/golang_boilerplate/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"net/http"
)

const (
	prefix                = "/api/v1"
	ingredientPostfix     = "/ingredient"
	usersPostfix          = "/user"
	ticketPostfix         = "/ticket"
	authenticationPostfix = "/authentication"
	permissionPostfix     = "/permission"
	rolePostfix           = "/role"
	rolePermissionPostfix = "/role-permission"
	userRolePostfix       = "/user-role"
)

func Routes(app *gin.Engine) {
	router := app.Group(prefix)
	//routerTicket := app.Group(prefix + ticketPostfix)
	//authUser := app.Group(prefix + authenticationPostfix)
	//authPermission := app.Group(prefix + permissionPostfix)
	//authRole := app.Group(prefix + rolePostfix)
	//RolePermissionRouter := app.Group(prefix + rolePermissionPostfix)
	//UserRoleRouter := app.Group(prefix + userRolePostfix)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Register all modules
	User.RegisterRoutes(router)
	Auth.RegisterRoutes(router)
	// AuthUser.RegisterRoutes(router)
	// Permission.RegisterRoutes(router)
	// Role.RegisterRoutes(router)
	// RolePermission.RegisterRoutes(router)
	// UserRole.RegisterRoutes(router)
}
