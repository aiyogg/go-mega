package controller

import (
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/dota2mm/go-mega/config"
	"github.com/dota2mm/go-mega/vm"
	"gopkg.in/gomail.v2"
)

// PopulateTemplates func
// Create map template name to template.Template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}
			defer f.Close()
			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}
			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}
			result[fi.Name()] = tmpl
		}()
	}

	return result
}

//region Session 操作

func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	fmt.Println("val: ", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can't get session user")
	}
	fmt.Println("username: ", username)
	return username, nil
}

func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

//endregion

//region 注册登录相关

func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}

// 验证表单字段合法性
func checkUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}
func checkPassword(password string) string {
	return checkLen("Password", password, 6, 50)
}
func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w._]{1,20})@(\w+).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}
func checkUserPassword(username, password string) string {
	if !vm.CheckLogin(username, password) {
		return fmt.Sprintf("Username and password is not correct")
	}
	return ""
}
func checkUserExist(username string) string {
	if !vm.CheckUserExist(username) {
		return fmt.Sprintf("Username already exist, please choose another username")
	}
	return ""
}

// checkLogin - 登录参数校验
func checkLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserPassword(username, password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// checkRegister - 注册参数校验
func checkRegister(username, email, pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkUserExist(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// addUser - 调用 vm.AddUser
func addUser(username, password, email string) error {
	return vm.AddUser(username, password, email)
}

//endregion

//region Flash Message

func setFlash(w http.ResponseWriter, r *http.Request, message string) {
	session, _ := store.Get(r, sessionName)
	session.AddFlash(message, flashName)
	session.Save(r, w)
}

func getFlash(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, sessionName)
	fm := session.Flashes(flashName)
	if fm == nil {
		return ""
	}
	session.Save(r, w)
	return fmt.Sprintf("%v", fm[0])
}

//endregion

//region pagination 分页

func getPage(r *http.Request) int {
	url := r.URL
	query := url.Query() // Values (map[string][]string)

	q := query.Get("page")
	if q == "" {
		return 1
	}
	page, err := strconv.Atoi(q)
	if err != nil {
		return 1
	}
	return page
}

//endregion

// region Mail
func checkEmailExistRegister(email string) string {
	if vm.CheckEmailExist(email) {
		return fmt.Sprintf("Email has registered by others, please use another email.")
	}
	return ""
}

// checkEmailExist 确认是否是已经注册过的邮箱
func checkEmailExist(email string) string {
	if !vm.CheckEmailExist(email) {
		return fmt.Sprintf("Email does not register yet. Please check email")
	}
	return ""
}
func sendMail(target, subject, content string) {
	server, port, user, pwd := config.GetSMTPConfig()
	d := gomail.NewDialer(server, port, user, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", target)
	m.SetAddressHeader("Cc", user, "admin")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Email Error:", err)
		return
	}
}

// endregion

// region 重置密码校验

func checkResetPasswordRequest(email string) []string {
	var errs []string
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmailExist(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}
func checkResetPassword(pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "Two password does not match")
	}
	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	return errs
}

// #endregion

// region Time format
const (
	minute  = 1
	hour    = minute * 60
	day     = hour * 24
	month   = day * 30
	year    = month * 365
	quarter = year / 4
)

func round(f float64) int {
	return int(math.Floor(f + .50))
}

// FromDuration returns a friendly string representing an approximation of the given duration
func FromDuration(d time.Duration) string {
	seconds := round(d.Seconds())
}

// endregion
