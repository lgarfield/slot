package db

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"slot/config"
)

type (
	LoginRpc struct {
		username string
	}

	LoginReply struct {

	}

	RegisterRpc struct {
		username string
		passwd string
		regIP string
		regTime string
	}

	RegisterReply struct {

	}
)

func (d *Db)Login(params LoginRpc, re *LoginReply) (err error) {
	session, err := mgo.Dial(config.MgoServer)
	defer session.Close()
	if err != nil {
		return err
	}

	c := session.DB("User").C("login")
	err = c.Insert(&LoginRpc{username:"good"})
	if err != nil {
		return err
	}

	return
}

func (d *Db)UserRegister(params RegisterRpc, re *RegisterReply) (err error) {
	session, err := mgo.Dial(config.MgoServer)
	defer session.Close()
	if err != nil {
		return err
	}

	c := session.DB("slot").C("userregister")
	// usernmae exist or not
	var result RegisterRpc
	err = c.Find(bson.M{"username": params.username}).One(&result)
	if err != nil {
		return err
	}
	if result.username != "" {
		return errors.New("Invalid username: exist already.")
	}

	// creat<e this record
	err = c.Insert(params)
	if err != nil {
		return err
	}

	return
}
