package services

import (
	"context"
	"guviTask/api/controllers/models"
	"guviTask/api/dao"
)

var LoginSvc LoginService= login{}

type LoginService interface {
	CheckCredentials(userprofile models.Userprofile)(error)
	AddUser(userprofile models.Userprofile)(error)

}

type login struct {

}

func (l login) AddUser(userprofile models.Userprofile) error {
	err:=dao.Loginsdao.LoginUser(context.TODO(),userprofile)
	return err}

func (l login) CheckCredentials(userprofile models.Userprofile) error {
	err:=dao.Loginsdao.CheckIfUserRegistered(context.TODO(),userprofile)
	return err
}

