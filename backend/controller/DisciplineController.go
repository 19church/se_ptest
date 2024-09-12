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

	if err := entity.DB().Raw("SELECT * FROM admins WHERE id = ?", id).Scan(&admin).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": admin})

}

// GET /admins

func ListAdmin(c *gin.Context) {

	var admins []entity.Admin

	if err := entity.DB().Raw("SELECT * FROM admins").Scan(&admins).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": admins})

}

// DELETE /admins/:id
func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM admins WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
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

	if err := entity.DB().Raw("SELECT * FROM students WHERE id = ?", id).Scan(&student).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": student})

}

// GET /students

func ListStudent(c *gin.Context) {

	var students []entity.Student

	if err := entity.DB().Raw("SELECT * FROM students").Scan(&students).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": students})

}

// DELETE /students/:id
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM students WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
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

	if err := entity.DB().Raw("SELECT * FROM discipline_types WHERE id = ?", id).Scan(&disciplineType).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplineType})

}

// GET /disciplineTypes

func ListDisciplineType(c *gin.Context) {

	var disciplineTypes []entity.DisciplineType

	if err := entity.DB().Raw("SELECT * FROM discipline_types").Scan(&disciplineTypes).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplineTypes})

}

// DELETE /admins/:id
func DeleteDisciplineType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM discipline_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DisciplineType not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
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

	var Discipline entity.Discipline
	var Admin entity.Admin
	var Student entity.Student
	var DisciplineType entity.DisciplineType

	//1
	if err := c.ShouldBindJSON(&Discipline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2: ค้นหา DisciplineType ด้วย id
	if tx := entity.DB().Where("id = ?", Discipline.DisciplineTypeID).First(&DisciplineType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DisciplineType not found"})
		return
	}

	// 3: ค้นหา Student ด้วย id
	if tx := entity.DB().Where("id = ?", Discipline.StudentID).First(&Student); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}

	// 4: ค้นหา Admin ด้วย id
	if tx := entity.DB().Where("id = ?", Discipline.AdminID).First(&Admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "visitor_type not found"})
		return
	}

	// 5: สร้าง Discipline
	dp := entity.Discipline{
		AdminID:               Discipline.AdminID,          // โยงความสัมพันธ์กับ Entity Admin
		StudentID:             Discipline.StudentID,        // โยงความสัมพันธ์กับ Entity Student
		DisciplineTypeID:      Discipline.DisciplineTypeID, // โยงความสัมพันธ์กับ Entity DisciplineType
		Discipline_Reason:     Discipline.Discipline_Reason,
		Discipline_Punishment: Discipline.Discipline_Punishment,
		Discipline_Point:      Discipline.Discipline_Point,
		Added_Time:            Discipline.Added_Time, // ตั้งค่าฟิลด์ Added_Time
	}

	//6: บันทึก
	if err := entity.DB().Create(&dp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dp})
}

// GET /discipline/:id

func GetDiscipline(c *gin.Context) {

	var discipline entity.Discipline

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM disciplines WHERE id = ?", id).Scan(&discipline).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": discipline})

}

// GET /disciplines

func ListDiscipline(c *gin.Context) {

	var disciplines []entity.Discipline

	if err := entity.DB().Raw("SELECT * FROM disciplines").Scan(&disciplines).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": disciplines})

}

// DELETE /disciplines/:id
func DeleteDiscipline(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM disciplines WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Discipline not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
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
