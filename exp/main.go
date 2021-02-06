package main

import (
	"fmt"
	"webdev_go/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "db"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	// Update a user
	user.Name = "Updated Name"
	if err := us.Update(&user); err != nil {
		panic(err)
	}

	foundUser, err := us.ByEmail("michael@dundermifflin.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	// Delete a user
	if err := us.Delete(foundUser.ID); err != nil {
		panic(err)
	}

	// Verify user is deleted
	_, err = us.ByID(foundUser.ID)
	if err != models.ErrNotFound {
		panic("user was not deleted!")
	}
}
