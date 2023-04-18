package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"guviTask/api/controllers/models"

)



var Registrationsdao Registrationdao= &registerProfiledao{}

type Registrationdao interface {
	Register(ctx context.Context,userprofile models.Userprofile) error
	CheckIfUserRegistered(ctx context.Context,email string) (models.Userprofile,error)
	GetUser(ctx context.Context,userEmail string)(models.Userprofile,error)
}

type registerProfiledao struct {
}

func (r *registerProfiledao) GetUser(ctx context.Context, userEmail string) (models.Userprofile, error) {
	panic("implement me")
}

func (r *registerProfiledao) Register(ctx context.Context, userprofile models.Userprofile) error {
	dburl:="mongodb://mongoadmin:secret@127.0.0.1:6000/?authSource=admin&readPreference=primaryPreferred&directConnection=true&ssl=false"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburl))

	userCollection := client.Database("test").Collection("userprofiles")

	_, err = userCollection.InsertOne(ctx, userprofile)

	if err != nil {
		fmt.Errorf("error creating user profile %s", userprofile.Name)
		customErr := fmt.Errorf("error creating user profile %s", userprofile.Name)

		return  customErr
	}

	return  nil
}

func (r *registerProfiledao) CheckIfUserRegistered(ctx context.Context,email string)  (models.Userprofile,error) {
	var user models.Userprofile
	dburl:="mongodb://mongoadmin:secret@127.0.0.1:6000/?authSource=admin&readPreference=primaryPreferred&directConnection=true&ssl=false"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburl))

	userCollection := client.Database("test").Collection("userprofiles")
	filter:=bson.M{"email": email}
	err = userCollection.FindOne(ctx, filter).Decode(&user)
	if err!= nil {
		fmt.Errorf("error getting userprofile")
		customErr := fmt.Errorf("error getting user profile %v", err)

		return  models.Userprofile{},customErr
	}

 	return user,nil
}


