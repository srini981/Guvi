package dao

import (
	"context"
	"database/sql"
	"fmt"
	"guviTask/api/controllers/models"
)

var Loginsdao logindao= &login{}

type logindao interface {
	LoginUser(ctx context.Context,userprofile models.Userprofile) error
	CheckIfUserRegistered(ctx context.Context,userprofile models.Userprofile) error
}

type login struct {

}

func (l login) LoginUser(ctx context.Context, userprofile models.Userprofile) error {
	const (
		host     = "localhost"
		port     = 6007
		user     = "root"
		password = "123"
		dbname   = "db_1"
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil{
		customErr:=fmt.Errorf("failed to create connection")
		return customErr
	}
	insertDynStmt := `insert into "users"("name", "email","password") values($1, $2,$3)`
	_, err = db.Exec(insertDynStmt,userprofile.Name, userprofile.Email,userprofile.Password)
	if err != nil{
		customErr:=fmt.Errorf("failed to run  query")
		return customErr
	}
	return nil
}

func (l login) CheckIfUserRegistered(ctx context.Context, userprofile models.Userprofile) error {
	const (
		host     = "localhost"
		port     = 6007
		user     = "root"
		password = "123"
		dbname   = "db_1"
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil{
		customErr:=fmt.Errorf("failed to create connection")
		return customErr
	}
	insertDynStmt := `select * from  users where email=$1 and password=$2;`
	_, err = db.Exec(insertDynStmt,userprofile.Email,userprofile.Password)
	if err != nil{
		customErr:=fmt.Errorf("failed to run  query")
		return customErr
	}
return nil
}
