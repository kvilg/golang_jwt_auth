package main

import (
	"fpProgect_norm/controllers"
	"fpProgect_norm/middlewares"
	"fpProgect_norm/models"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	models.ConnectDataBase()
	//err = models.DB.Migrator().DropTable(models.User{})
	//if err != nil {
	//	return
	//}

	//err = models.DB.Migrator().CreateTable(models.User{})
	//if err != nil {
	//	return
	//}

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/reg", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	err = r.Run(":8080")
	if err != nil {
		return
	}

}
