package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joserafaelSH/go-gorm-gin-api/controllers"
)


func HandleRequest(){
	r:= gin.Default()

	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Saudation)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.ShowStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.ShowStudentByCpf)
	r.Run()
}

