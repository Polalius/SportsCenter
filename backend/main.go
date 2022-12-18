package main

import (
	"github.com/Polalius/sa-65-example/controller"
	"github.com/Polalius/sa-65-example/entity"
	"github.com/Polalius/sa-65-example/middlewares"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

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

	//Route API

	//admin part
	adminApi := r.Group("/admin")
	{
		protected := adminApi.Use(middlewares.AuthorizedAdmin())
		{
			//role
			protected.GET("/roles", controller.ListRoles)
			protected.GET("/role/:id", controller.GetRole)
			protected.POST("/role", controller.CreateRole)
			protected.PATCH("/role", controller.UpdateRole)
			protected.DELETE("/role/:id", controller.DeleteRole)

			//login
			//Don't have post because we will create login when create employee (1 - 1 relations)
			protected.GET("/logins", controller.ListLogin)
			protected.GET("/login/:id", controller.GetLogin)
			protected.PATCH("/login", controller.UpdateLogin)
			protected.DELETE("/login/:id", controller.DeleteLogin)

			//employee
			protected.GET("/employees", controller.ListEmployee)
			protected.GET("/employee/:id", controller.GetEmployee)
			protected.POST("/employee", controller.CreateEmployee)
			protected.PATCH("/employee", controller.UpdateEmployee)
			protected.DELETE("/employee/:id", controller.DeleteEmployee)

		}
	}

	//trainer (roleName trainer)
	trainerApi := r.Group("/trainer")
	{
		protected := trainerApi.Use(middlewares.AuthorizedTrainer())
		{
			//เพิ่ม API ตรงส่วนนี้ ในกรณีเรียกใช้ ให้เรียกใช้จาก /explist/(...Route)
			protected.GET("/employee/:id", controller.GetEmployee)
			protected.GET("/employees", controller.ListEmployee)

			// WormUp Routes
			protected.GET("/wormup", controller.ListWormUp)
			protected.GET("/wormup/:id", controller.GetWormUp)
			protected.POST("/wormup", controller.CreateWormUp)
			protected.PATCH("/wormup", controller.UpdateWormUp)
			protected.DELETE("/wormup/:id", controller.DeleteWormUp)

			// Exercise Routes
			protected.GET("/exercise", controller.ListExercise)
			protected.GET("/exercise/:id", controller.GetExercise)
			protected.POST("/exercise", controller.CreateExercise)
			protected.PATCH("/exercise", controller.UpdateExercise)
			protected.DELETE("/exercise/:id", controller.DeleteExercise)

			// Stretch Routes
			protected.GET("/stretch", controller.ListStretch)
			protected.GET("/stretch/:id", controller.GetStretch)
			protected.POST("/stretch", controller.CreateStretch)
			protected.PATCH("/stretch", controller.UpdateStretch)
			protected.DELETE("/stretch/:id", controller.DeleteStretch)

			//ExerciseProgramList Routes
			protected.GET("/explist", controller.ListExPList)
			protected.GET("/explist/:id", controller.GetExPList)
			protected.POST("/explist", controller.CreateExerciseProgramList)
			protected.PATCH("/explist", controller.UpdateExPList)
			protected.DELETE("/explist/:id", controller.DeleteExPList)

		}
	}

	//all user login can use

	//For signin (Auth Route)
	r.POST("/signin", controller.Signin)
	r.GET("/employee/:id", controller.GetEmployeeByLoginID)

	//for check token
	r.GET("/valid", controller.Validation)

	r.Run()
}
