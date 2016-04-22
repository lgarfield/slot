package client

import(
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
	//"net/http"
	"reflect"
	"slot/config"
)

func userRegister(c *gin.Context, db *gorp.DbMap) (result []reflect.Value, err error) {
	// register
	Register := slot_user_basic{}
	_, err = Register.Register(c, db)
	if err != nil {
		return
	}

	// login for the auth-token
	Login := slot_login_session{}
	result, err = Login.Login(c, db)
	if err != nil {
		return
	}

	return
}

/**
 * table slot_login_session -- for login
 *
 *
 */
type slot_login_session struct {
	Id int `form:"id" json:"id" db:"id,size:11,primarykey,autoincrement"`
	Account string `form:"account" json:"account" db:"account,size:20" binding:"required"`
	LoginIp string `db:"login_ip"`
	LoginTime string `db:"login_time"`
	LoginSession string `db:"login_session"`
}

func (l *slot_login_session) Login(c *gin.Context, db *gorp.DbMap) (result []reflect.Value, err error) {
	l.Account = c.PostForm("account")
	l.LoginTime = config.SelfTime()
	l.LoginIp = config.SelfIp(c.Request)
	l.LoginSession = "ssfesdfnnanniensndfjensndnf"

	err = db.Insert(&l)
	if err != nil {
		return
	}

	result = config.Result{}
	return
}

/**
 * table slot_user_basic -- for register
 *
 *
 */
type slot_user_basic struct {
	Id int `form:"id" json:"id" db:"id,size:11,primarykey,autoincrement"`
	Account string `form:"account" json:"account" binding:"required" db:"account,size:20"`
	Passwd string `form:"passwd" json:"passwd" binding:"required" db:"passwd,size:20"`
	RegDate string `db:"reg_date"`
	RegIp string `db:"reg_ip"`
}

func (u *slot_user_basic)Register(c *gin.Context, db *gorp.DbMap) (result []reflect.Value, err error) {
	if c.BindJSON(&u) == nil {
		// judgement account that posted is exist or not.
		// TODO

		// set reg_date, reg_ip
		u.RegDate = config.SelfTime()
		u.RegIp = config.SelfIp(c.Request)

		// register
		err = db.Insert(&u)
		if err != nil {
			return
		}

		return
	}

	err = errors.New("Nothing is provided.")
	return
}
