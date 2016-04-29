package client

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
	"slot/config"
)

func userRegister(c *gin.Context, db *gorp.DbMap) (result []interface{}, err error) {
	// register
	_, err = Register(c, db)
	if err != nil {
		return
	}

	// login for the auth-token
	result, err = Login(c, db)
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
	Id int `form:"id" json:"id" db:"id"`
	Account string `form:"account" json:"account" db:"account" binding:"required"`
	LoginIp string `db:"login_ip"`
	LoginTime string `db:"login_time"`
	LoginSession string `db:"login_session"`
}

type login_result struct {
	Session string `json:"session"`
}

func Login(c *gin.Context, db *gorp.DbMap) (result []interface{}, err error) {
	l := slot_login_session{
		Account:c.PostForm("account"),
		LoginTime:config.SelfTime(),
		LoginIp:c.ClientIP(),
		LoginSession:"aaaaaaaaaaaaaaa",
	}

	db.AddTableWithName(slot_login_session{}, "slot_login_session").SetKeys(true, "Id")
	err = db.Insert(&l)
	if err != nil {
		return
	}

	login_result := login_result{Session:l.LoginSession}
	result = append(result, login_result)

	return
}

/**
 * table slot_user_basic -- for register
 *
 *
 */
type slot_user_basic struct {
	Id int `form:"id" json:"id" db:"id"`
	Account string `form:"account" json:"account" binding:"required" db:"account"`
	Passwd string `form:"passwd" json:"passwd" binding:"required" db:"passwd"`
	RegDate string `db:"reg_date"`
	RegIp string `db:"reg_ip"`
}

func Register(c *gin.Context, db *gorp.DbMap) (result []interface{}, err error) {
	u := slot_user_basic{
		Id:0,
		Account:"",
		Passwd:"",
		RegDate:config.SelfTime(),
		RegIp:c.ClientIP(),
	}
	if err = c.Bind(&u); err != nil {
		fmt.Print(err)
		return
	}

	// judgement account that posted is exist or not.
	// TODO

	// register
	db.AddTableWithName(slot_user_basic{}, "slot_user_basic").SetKeys(true, "Id")
	err = db.Insert(&u)
	if err != nil {
		return
	}

	return
}
