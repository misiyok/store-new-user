package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

// User defines the structure for an API user
type User struct {
	ID          int    `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" validate:"required" bson:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber" validate:"required" bson:"phone_number"`
	CreatedOn   string `json:"-" bson:"created_on,omitempty"`
	UpdatedOn   string `json:"-" bson:"updated_on,omitempty"`
	DeletedOn   string `json:"-" bson:"deleted_on,omitempty"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)

	return d.Decode(u)
}

func AddUser(u *User, mc *mongo.Client) {
	coll := mc.Database("bookSwapDB").Collection("users")

	result, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New user has been created, id=%s\n", result.InsertedID)
	// oid, ok := result.InsertedID.(primitive.ObjectID)
	// if ok {
	// 	fmt.Printf("New user has been created, id=%s\n", oid.Hex())
	// }
}
