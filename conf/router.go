package conf

import (
	. "nana/admin"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	adminH := router.Group("/admin")
	{
		adminH.GET("/login", Login)
	}

	//	router.POST("/person", AddPersonApi)

	//	router.GET("/persons", GetPersonsApi)

	//	router.GET("/person/:id", GetPersonApi)

	//	router.PUT("/person/:id", ModPersonApi)

	//	router.DELETE("/person/:id", DelPersonApi)

	return router
}
