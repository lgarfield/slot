package login

import (

)

type (
	LoginRpc struct {
		Username string
		Passwd string
	}

	LoginReply struct {

	}
)

func (l *Login)Login(params LoginRpc, re *LoginReply) (err error) {
	return err
}
