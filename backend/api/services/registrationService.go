package services

import (
	"context"
	"guviTask/api/controllers/models"
	"guviTask/api/dao"
)


var RegistrationsSvc RegistrationService= registerProfile{}

type RegistrationService interface {
	RegisterUser (userprofile models.Userprofile)(error)
	CheckIfUserRegistered(email string)(models.Userprofile,error)
	UserDetails(userprofile string)(models.Userprofile,error)

}

type registerProfile struct {

}

func (r registerProfile) UserDetails(userprofile string) (models.Userprofile, error) {
	panic("implement me")
}

func (r registerProfile)RegisterUser (userprofile models.Userprofile)(error) {

	err:=dao.Registrationsdao.Register(context.TODO(),userprofile)

	return err
}
func (r registerProfile)CheckIfUserRegistered(email string)(models.Userprofile,error) {
	user,err:=dao.Registrationsdao.CheckIfUserRegistered(context.TODO(),email)
	return user,err
}
