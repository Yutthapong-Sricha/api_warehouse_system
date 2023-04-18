package sessions

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

// this map stores the users sessions. For larger scale applications, you can use a database or cache for this purpose
//var sessions = map[string]session{}

// each session contains the username of the user and the time at which it expires
type session struct {
	username string
	expiry   time.Time
}

// we'll use this method later to determine if the session has expired
func isExpired(expiry time.Time) bool {
	return expiry.Before(time.Now())
}

// SessionLength is the length of the max life of a cookie in seconds
// 3600s = 1hr
const SessionLength int = 21600 // 6 hours

func AddSessionDB(c *gin.Context, sessionToken string, LogLogin modelsStruc.LogLogin) string {
	clientIp := c.ClientIP()
	agent := c.Request.UserAgent()
	session_data, err := json.Marshal(LogLogin)
	if err != nil {
		log.Fatalln(err.Error())
	}
	sqlStmt := `INSERT INTO oauth_sessions(session_id, session_start, session_last_activity, session_ip_address, session_user_agent, session_data) 
		VALUES (?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), ?, ?, ? )`
	stmt, err := initializers.DB.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionToken, clientIp, agent, session_data)
	if err != nil {
		log.Fatalln(err)
	}

	return string(session_data)
}

func UpdateSessionDB() {

}

func SessionLogin(c *gin.Context, LogLogin modelsStruc.LogLogin) (string, string) {
	now := time.Now()
	var timestamp = strconv.FormatInt(now.Unix(), 10)
	login_time, _ := strconv.ParseInt(timestamp, 10, 0)
	sessionToken := uuid.NewString()

	sqlStmt := `INSERT INTO oauth_log_user_login(session_id, id_user, id_branch, email_login, is_success1_fail2_flag, login_time) 
				VALUES (?, ?, ?, ?, ?, UNIX_TIMESTAMP())`
	stmt, err := initializers.DB.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionToken, LogLogin.Id_user, LogLogin.Id_branch, LogLogin.Email_login, 1)
	if err != nil {
		log.Fatalln(err)
	}

	LogLogin.Is_success1_fail2_flag = 1
	LogLogin.Login_time = int(login_time)

	session_data := AddSessionDB(c, sessionToken, LogLogin)

	expiresAt := time.Now().Add(21600 * time.Second)

	mysession := sessions.Default(c)
	mysession.Set(sessionToken, LogLogin.Staff_fullname)
	err_mysess := mysession.Save()
	if err_mysess != nil {
		log.Fatalln(err_mysess.Error())
	}
	fmt.Println(sessionToken)
	value := mysession.Get(sessionToken)
	fmt.Println(value)

	//c.SetCookie("session_token", sessionToken, SessionLength, "/", "127.0.0.1", false, false)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		MaxAge:  SessionLength,
		Expires: expiresAt,
	})

	return sessionToken, session_data
}

func SessionLogout(c *gin.Context) (bool, string) {
	Authorization := c.Request.Header.Get("Authorization")
	bearerToken := strings.TrimPrefix(Authorization, "Bearer ")
	if len(bearerToken) == 0 {
		return false, "not found token"
	}

	mysession := sessions.Default(c)
	mysession.Delete(bearerToken)
	err := mysession.Save()
	if err != nil {
		log.Fatalln(err)
		return false, err.Error()
	}

	sqlStmt := `UPDATE oauth_log_user_login
				SET logout_time = CURRENT_TIMESTAMP()
				WHERE session_id = ?`
	stmt, err := initializers.DB.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
		return false, err.Error()
	}
	defer stmt.Close()
	_, err = stmt.Exec(bearerToken)
	if err != nil {
		log.Fatalln(err)
		return false, err.Error()
	}

	sqlStmt_del := `DELETE oauth_sessions
					WHERE session_id = ?`
	stmt_del, err_del := initializers.DB.Prepare(sqlStmt_del)
	if err_del != nil {
		log.Fatalln(err_del)
		return false, err_del.Error()
	}
	defer stmt_del.Close()

	_, err = stmt_del.Exec(bearerToken)
	if err != nil {
		log.Fatalln(err)
		return false, err_del.Error()
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		MaxAge:  0,
		Expires: time.Now(),
	})
	return true, "logout success"
}

func SessionAuth(c *gin.Context) error {
	Authorization := c.Request.Header.Get("Authorization")
	bearerToken := strings.TrimPrefix(Authorization, "Bearer ")

	fmt.Println(bearerToken)

	mysession := sessions.Default(c)
	sess_value := mysession.Get(bearerToken)
	if sess_value == nil {
		// If the session token is not present in session map, return an unauthorized error
		myError := errors.New("there is no user in the session")
		return myError
	}

	return nil
}
