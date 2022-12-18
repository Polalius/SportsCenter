package controller

import (
	"net/http"

	"github.com/Polalius/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

//------------------------------- Worm Up -----------------------------------

// POST /wormup
func CreateWormUp(c *gin.Context) {
	var WormUp entity.WormUp
	if err := c.ShouldBindJSON(&WormUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&WormUp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": WormUp})
}

// GET /wormup/:id
func GetWormUp(c *gin.Context) {
	var WormUp entity.WormUp
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM worm_up WHERE id = ?", id).Scan(&WormUp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": WormUp})
}

// List /wormup
func ListWormUp(c *gin.Context) {
	var WormUp []entity.WormUp
	if err := entity.DB().Raw("SELECT * FROM worm_up").Scan(&WormUp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": WormUp})
}

// DELETE /wormup/:id
func DeleteWormUp(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM worm_up WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WormUp not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /wormup
func UpdateWormUp(c *gin.Context) {
	var WormUp entity.WormUp
	if err := c.ShouldBindJSON(&WormUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", WormUp.ID).First(&WormUp); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}

	if err := entity.DB().Save(&WormUp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": WormUp})
}

//------------------------------- Exercise -----------------------------------

// POST /exercise
func CreateExercise(c *gin.Context) {
	var Exercise entity.Exercise
	if err := c.ShouldBindJSON(&Exercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Exercise).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Exercise})
}

// GET /exercise/:id
func GetExercise(c *gin.Context) {
	var Exercise entity.Exercise
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM exercise WHERE id = ?", id).Scan(&Exercise).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Exercise})
}

// List /exercise
func ListExercise(c *gin.Context) {
	var Exercise []entity.Exercise
	if err := entity.DB().Raw("SELECT * FROM exercise").Scan(&Exercise).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Exercise})
}

// DELETE /exercise/:id
func DeleteExercise(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM exercise WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exercise not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /exercise
func UpdateExercise(c *gin.Context) {
	var Exercise entity.Exercise
	if err := c.ShouldBindJSON(&Exercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Exercise.ID).First(&Exercise); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}

	if err := entity.DB().Save(&Exercise).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Exercise})
}

//------------------------------- Stretch -----------------------------------

// POST /stretch
func CreateStretch(c *gin.Context) {
	var Stretch entity.Stretch
	if err := c.ShouldBindJSON(&Stretch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Stretch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Stretch})
}

// GET /stretch/:id
func GetStretch(c *gin.Context) {
	var Stretch entity.Stretch
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM stretch WHERE id = ?", id).Scan(&Stretch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Stretch})
}

// List /stretch
func ListStretch(c *gin.Context) {
	var Stretch []entity.Stretch
	if err := entity.DB().Raw("SELECT * FROM stretch").Scan(&Stretch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Stretch})
}

// DELETE /stretch/:id
func DeleteStretch(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM stretch WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stretch not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /stretch
func UpdateStretch(c *gin.Context) {
	var Stretch entity.Stretch
	if err := c.ShouldBindJSON(&Stretch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Stretch.ID).First(&Stretch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}

	if err := entity.DB().Save(&Stretch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Stretch})
}

//----------------------------------- Program -------------------------------

// POST /program
func CreateExerciseProgramList(c *gin.Context) {

	var employee entity.Employee
	var wormup entity.WormUp
	var exercise entity.Exercise
	var stretch entity.Stretch
	var explist entity.ExerciseProgramList

	// ผลลัพธ์ที่ได้จากขั้นตอนที่  จะถูก bind เข้าตัวแปร exprlist
	if err := c.ShouldBindJSON(&explist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", explist.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// ค้นหา wormup ด้วย id
	if tx := entity.DB().Where("id = ?", explist.WormUpID).First(&wormup); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "worm up not found"})
		return
	}

	// ค้นหา exercise ด้วย id
	if tx := entity.DB().Where("id = ?", explist.ExerciseID).First(&exercise); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "exercise not found"})
		return
	}
	// ค้นหา stretch ด้วย id
	if tx := entity.DB().Where("id = ?", explist.StretchID).First(&stretch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stretch not found"})
		return
	}
	// 12: สร้าง ex prog list
	list := entity.ExerciseProgramList{
		Employee: employee,       // โยงความสัมพันธ์กับ Entity Employee
		WormUp:   wormup,         // โยงความสัมพันธ์กับ Entity WormUp
		Exercise: exercise,       // โยงความสัมพันธ์กับ Entity Exercise
		Stretch:  stretch,        // โยงความสัมพันธ์กับ Entity Stretch
		Minute:   explist.Minute, // ตั้งค่าฟิลด์ Minute
	}

	// 13: บันทึก
	if err := entity.DB().Create(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

// GET /explist/:id
func GetExPList(c *gin.Context) {
	var explist entity.ExerciseProgramList
	id := c.Param("id")
	if err := entity.DB().Preload("Employee").Preload("WormUp").Preload("Exercise").Preload("Stretch").Raw("SELECT * FROM exercise_program_list WHERE id = ?", id).Find(&explist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": explist})
}

// LIST /explist
func ListExPList(c *gin.Context) {
	var explist []entity.ExerciseProgramList
	if err := entity.DB().Preload("Employee").Preload("WormUp").Preload("Exercise").Preload("Stretch").Raw("SELECT * FROM exercise_program_list").Find(&explist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": explist})
}

// DELETE /explist/:id
func DeleteExPList(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM exercise_program_list WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "list not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /explist
func UpdateExPList(c *gin.Context) {
	var explist entity.ExerciseProgramList
	if err := c.ShouldBindJSON(&explist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", explist.ID).First(&explist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "list not found"})
		return
	}

	if err := entity.DB().Save(&explist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": explist})
}
