package controller

import (
	"net/http"

	"github.com/19church/se_dicipline/entity"
	"github.com/gin-gonic/gin"
)

//Admin

// POST /admins

func CreateAdmin(c *gin.Context) {

	var admin entity.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&admin).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": admin})

}

// GET /admin/:id

func GetAdmin(c *gin.Context) {

	var admin entity.Admin

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&admin).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": admin})

}

// GET /admins

func ListAdmin(c *gin.Context) {

	var admins []entity.Admin

	if err := entity.DB().Raw("SELECT * FROM users").Scan(&admins).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": admins})

}

// PATCH /admins

func UpdateAdmin(c *gin.Context) {

	var admin entity.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", admin.ID).First(&admin); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "admin not found"})

		return

	}

	if err := entity.DB().Save(&admin).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": admin})

}

//Student

// POST /students

func CreateStudent(c *gin.Context) {

	var student entity.Student

	if err := c.ShouldBindJSON(&student); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&student).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": student})

}

// GET /student/:id

func GetStudent(c *gin.Context) {

	var student entity.Student

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&student).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": student})

}

// GET /students

func ListStudent(c *gin.Context) {

	var students []entity.Student

	if err := entity.DB().Raw("SELECT * FROM users").Scan(&students).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": students})

}

// PATCH /students

func UpdateStudent(c *gin.Context) {

	var student entity.Student

	if err := c.ShouldBindJSON(&student); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", student.ID).First(&student); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})

		return

	}

	if err := entity.DB().Save(&student).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": student})

}

//DisciplineType

// POST /disciplineType

func CreateDisciplineType(c *gin.Context) {

	var disciplineType entity.DisciplineType

	if err := c.ShouldBindJSON(&disciplineType); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&disciplineType).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplineType})

}

// GET /disciplineType/:id

func GetDisciplineType(c *gin.Context) {

	var disciplineType entity.DisciplineType

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&disciplineType).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplineType})

}

// GET /disciplineTypes

func ListDisciplineType(c *gin.Context) {

	var disciplineTypes []entity.DisciplineType

	if err := entity.DB().Raw("SELECT * FROM users").Scan(&disciplineTypes).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplineTypes})

}

// PATCH /disciplineTypes

func UpdateDisciplineType(c *gin.Context) {

	var disciplineType entity.DisciplineType

	if err := c.ShouldBindJSON(&disciplineType); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", disciplineType.ID).First(&disciplineType); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "disciplineType not found"})

		return

	}

	if err := entity.DB().Save(&disciplineType).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplineType})

}

//Discipline

// POST /discipline

func CreateDiscipline(c *gin.Context) {

	var discipline entity.Discipline

	if err := c.ShouldBindJSON(&discipline); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&discipline).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": discipline})

}

// GET /discipline/:id

func GetDiscipline(c *gin.Context) {

	var discipline entity.Discipline

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&discipline).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": discipline})

}

// GET /disciplines

func ListDiscipline(c *gin.Context) {

	var disciplines []entity.Discipline

	if err := entity.DB().Raw("SELECT * FROM users").Scan(&disciplines).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplines})

}

// PATCH /discipline

func UpdateDiscipline(c *gin.Context) {

	var discipline entity.Discipline

	if err := c.ShouldBindJSON(&discipline); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", discipline.ID).First(&discipline); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "discipline not found"})

		return

	}

	if err := entity.DB().Save(&discipline).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": discipline})

}
