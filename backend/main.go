package main

import (
	"github.com/gin-gonic/gin"

	"github.com/19church/se_dicipline/controller"
	"github.com/19church/se_dicipline/entity"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	/////////////////////////////////////////////////////////////

	//Admin
	r.POST("/CreateAdmin", controller.CreateAdmin)
	r.GET("/GetAdmin/:id", controller.GetAdmin)
	r.GET("/ListAdmin", controller.ListAdmin)
	r.DELETE("/DeleteAdmin/:id", controller.DeleteAdmin)
	r.PATCH("/UpdateAdmin", controller.UpdateAdmin)

	//Student
	r.POST("/CreateStudent", controller.CreateStudent)
	r.GET("/GetStudent/:id", controller.GetStudent)
	r.GET("/ListStudent", controller.ListStudent)
	r.DELETE("/DeleteStudent/:id", controller.DeleteAdmin)
	r.PATCH("/UpdateStudent", controller.UpdateStudent)

	//DisciplineType
	r.POST("/CreateDisciplineType", controller.CreateDisciplineType)
	r.GET("/GetDisciplineType/:id", controller.GetDisciplineType)
	r.GET("/ListDisciplineType", controller.ListDisciplineType)
	r.DELETE("/DeleteDisciplineType/:id", controller.DeleteDisciplineType)
	r.PATCH("/UpdateDisciplineType", controller.UpdateDisciplineType)

	//TODO....abcde

	//Discipline
	r.POST("/CreateDiscipline", controller.CreateDiscipline)
	r.GET("/GetDiscipline/:id", controller.GetDiscipline)
	r.GET("/ListDiscipline", controller.ListDiscipline)
	r.DELETE("/DeleteDiscipline", controller.DeleteDiscipline)
	r.PATCH("/UpdateDiscipline", controller.UpdateDiscipline)

	r.Run()
}
