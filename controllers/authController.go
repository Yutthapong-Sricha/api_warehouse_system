package controllers

import (
	"log"
	"net/http"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"

	WSession "server_go/api_warehouse_system/models/sessions"
)

func Login(c *gin.Context) {
	var json modelsStruc.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//var SessionStruc modelsStruc.SessionStruc
	var LogLogin modelsStruc.LogLogin
	var err error
	var pwd_verification string
	// var staff_fullname string
	// var tmp_position_name string
	statement, err := initializers.DB.Prepare("SELECT id_staff_core AS id_user, id_branch,  tmp_branch_name, staff_email AS email_login, pwd_verification, staff_fullname, IFNULL(tmp_position_name,'') AS tmp_position_name FROM staff_core WHERE id_branch=? and staff_email=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	row_auth := statement.QueryRow(json.Id_branch, json.Username)
	err = row_auth.Scan(
		&LogLogin.Id_user,
		&LogLogin.Id_branch,
		&LogLogin.Tmp_branch_name,
		&LogLogin.Email_login,
		&pwd_verification,
		&LogLogin.Staff_fullname,
		&LogLogin.Tmp_position_name,
	)

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "ไม่พบข้อมูลในระบบ"})
		return
	}
	err_valid := bcrypt.CompareHashAndPassword([]byte(pwd_verification), []byte(json.Password))
	if err_valid != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "รหัสผ่านไม่ถูกต้อง"})
		return
	}

	sessionToken, session_data := WSession.SessionLogin(c, LogLogin)

	c.IndentedJSON(200, gin.H{
		"status":  "success",
		"message": "Login Success",
		"token":   sessionToken,
		"data":    session_data,
	})

}

func Logout(c *gin.Context) {
	//fmt.Println(c.Request.Header.Get("Authorization"))
	logout_status, message := WSession.SessionLogout(c)
	status := "error"
	if logout_status {
		status = "success"
	}

	c.IndentedJSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}
