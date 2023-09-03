package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joserafaelSH/go-gorm-gin-api/database"
	"github.com/joserafaelSH/go-gorm-gin-api/models"
)

func ShowAllStudents(c *gin.Context) {

	var alunos []models.Aluno
	database.Db.Find(&alunos)

	c.JSON(200, alunos)
}

func Saudation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

func CreateStudent(c *gin.Context){
	var recievedStudent models.Aluno
	if err := c.ShouldBindJSON(&recievedStudent); err != nil{
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}
	database.Db.Create(&recievedStudent)
	c.JSON(200, recievedStudent)
}

func ShowStudent(c *gin.Context){
	id := c.Params.ByName("id")
	var student models.Aluno
	database.Db.Find(&student, id)
	if student.ID <= 0 {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}
	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context){
	id := c.Params.ByName("id")
	var student models.Aluno
	database.Db.Find(&student, id)
	if student.ID <= 0 {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}
	database.Db.Delete(&student, id)
	c.JSON(200, gin.H{
		"message": "Student deleted",
		"data": student,
	})

}

func UpdateStudent(c *gin.Context){
	id := c.Params.ByName("id")
	var student models.Aluno
	database.Db.Find(&student, id)
	if student.ID <= 0 {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}

	
	if err := c.ShouldBindJSON(&student); err != nil{
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}

	database.Db.Model(&student).UpdateColumns(student)
	c.JSON(200, student)

}


func ShowStudentByCpf(c *gin.Context){
	cpf := c.Params.ByName("cpf")
	if len(cpf) <= 0 {
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}

	var student models.Aluno
	//database.Db.Model(&student).Where("cpf = ?", cpf).First(&student)
	database.Db.Model(&student).Where(&models.Aluno{Cpf: cpf}).First(&student)
	if student.ID <= 0 {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}

	c.JSON(200, student)
}