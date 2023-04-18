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

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

var sessions = map[string]session{}

// each session contains the username of the user and the time at which it expires
type session struct {
	username string
	expiry   time.Time
}

// we'll use this method later to determine if the session has expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

var DBSessionsCleaned = time.Now()

// SessionLength is the length of the max life of a cookie in seconds
// 3600s = 1hr
const SessionLength int = 43200 // 12 hours

func CreateSessionBD(c *gin.Context, sessionToken string, LogLogin modelsStruc.LogLogin) string {

	clientip := c.ClientIP()
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

	_, err = stmt.Exec(sessionToken, clientip, agent, session_data)
	if err != nil {
		log.Fatalln(err)
	}

	expiresAt := time.Now().Add(43200 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	sessions[sessionToken] = session{
		username: LogLogin.Staff_fullname,
		expiry:   expiresAt,
	}

	cookie := http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		MaxAge:  SessionLength,
		Expires: expiresAt,
	}
	http.SetCookie(c.Writer, &cookie)
	return string(session_data)
}

// CreateSession creates a new session entry
func CreatelogUserLogin(c *gin.Context, LogLogin modelsStruc.LogLogin) (string, string) {
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
	session_data := CreateSessionBD(c, sessionToken, LogLogin)

	return sessionToken, session_data
}

func RefreshSessionDB(c *gin.Context, sid string) {
	sqlStmt := `UPDATE oauth_sessions
				SET session_last_activity = CURRENT_TIMESTAMP()
				WHERE sessions_id = ?;`

	stmt, err := initializers.DB.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sid)
	if err != nil {
		log.Fatalln(err)
	}
}

func DeleteSessionDB(sid string) {
	sqlStmt := `DELETE FROM oauth_sessions
				WHERE sessions_id = ?`

	stmt, err := initializers.DB.Prepare(sqlStmt)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sid)
	if err != nil {
		log.Fatalln(err)
	}
}

func AuthSession(c *gin.Context) error {
	Authorization := c.Request.Header.Get("Authorization")
	bearerToken := strings.TrimPrefix(Authorization, "Bearer ")
	fmt.Println(bearerToken)
	//fmt.Println("xxxx")
	reqCookie, err := c.Request.Cookie("session_token")
	if err != nil {
		//return err
		myError := errors.New("Yud")
		return myError
	}
	sesstionToken := reqCookie.Value

	//fmt.Println(sesstionToken)
	if bearerToken != sesstionToken {
		myError := errors.New("token unuthorized")
		return myError
	}

	userSession, exists := sessions[bearerToken]
	fmt.Println(userSession.username)
	fmt.Println(userSession.expiry)
	if !exists {
		myError := errors.New("sessions not found")
		return myError
	}

	if userSession.isExpired() {
		delete(sessions, bearerToken)
		myError := errors.New("sessions Expired")
		return myError
	}

	return nil
}
