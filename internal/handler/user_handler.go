package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Alice Smith"},
	{ID: 3, Name: "Bob Johnson"},
}

// GetUserByID ฟังก์ชันที่ใช้ดึงข้อมูลผู้ใช้ตาม ID
// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags Users
// @Produce  json
// @Param   id   path      int     true  "User ID"
// @Success 200  {object}  User
// @Failure 404  {object}  ErrorResponse
// @Router  /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "name": "ณัฐโชติ พรหมฤทธิ์"})
}

// GetAllUsers ฟังก์ชันที่ใช้ดึงข้อมูลผู้ใช้ทั้งหมด
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags Users
// @Produce  json
// @Success 200  {array} User
// @Router  /users [get]
func GetAllUsers(c *gin.Context) {
	c.JSON(200, users)
}

// Healthcheck ฟังก์ชันที่ใช้ตอบคำขอสำหรับการตรวจสอบสถานะ
// @Summary Healthcheck API
// @Description Return a simple health check message
// @Tags Health
// @Produce  json
// @Success 200  {string}  string  "i am fine"
// @Router  /health [get]
func Healthcheck(c *gin.Context) {
	c.JSON(200, "i am fine")
}

// CreateUser ฟังก์ชันที่ใช้สร้างผู้ใช้ใหม่
// @Summary Create a new user
// @Description Create a new user with provided name
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   user  body      User  true  "User data"
// @Success 201  {object}  User
// @Failure 400  {object}  ErrorResponse
// @Router  /users [post]
func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, ErrorResponse{Message: "Invalid input"})
		return
	}

	// ตัวอย่างการสร้างผู้ใช้ใหม่
	newUser.ID = 4
	c.JSON(201, newUser)
}

// UpdateUser ฟังก์ชันที่ใช้อัปเดตข้อมูลผู้ใช้
// @Summary Update a user by ID
// @Description Update details of a user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id   path      int     true  "User ID"
// @Param   user  body      User  true  "User data"
// @Success 200  {object}  User
// @Failure 404  {object}  ErrorResponse
// @Router  /users/{id} [put]
func UpdateUser(c *gin.Context) {
	// ดึงค่า id จาก URL
	id := c.Param("id")

	// แปลง id จาก string เป็น int

	parsedID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, ErrorResponse{Message: "Invalid ID format"})
		return
	}

	// สร้างข้อมูลผู้ใช้ที่ต้องการอัปเดต
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, ErrorResponse{Message: "Invalid input"})
		return
	}

	// กำหนด ID ที่แปลงแล้วให้กับผู้ใช้
	updatedUser.ID = parsedID

	// ส่งข้อมูลผู้ใช้ที่อัปเดตแล้วกลับไป
	c.JSON(200, updatedUser)
}

// DeleteUser ฟังก์ชันที่ใช้ลบผู้ใช้
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags Users
// @Param   id   path      int     true  "User ID"
// @Success 200  {string}  string  "User deleted successfully"
// @Failure 404  {object}  ErrorResponse
// @Router  /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// ตัวอย่างการลบผู้ใช้
	c.JSON(200, gin.H{"message": "User deleted successfully", "id": id})
}
