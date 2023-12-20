package main

import (
	irisSwagger "github.com/dingjc89/iris-swagger"
	"github.com/dingjc89/iris-swagger/example/api"
	_ "github.com/dingjc89/iris-swagger/example/docs"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	swaggerFiles "github.com/swaggo/files"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8000
// @BasePath /api/v1
func main() {
	router := iris.New()
	router.UseRouter(cors.New().ExtractOriginFunc(cors.DefaultOriginExtractor).
		ReferrerPolicy(cors.NoReferrerWhenDowngrade).
		AllowOriginFunc(cors.AllowAnyOrigin).
		Handler())
	user := api.UserController{}
	router.Get("/api/v1/detail", (&user).Detail)

	router.Get("/swagger/{*}", irisSwagger.IrisWrapHandler(swaggerFiles.Handler))

	router.Listen(":8000")
}
