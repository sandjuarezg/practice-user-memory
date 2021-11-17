package user

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	User string
	Pass string
	Post []string
}

func AddUser() (u User, err error) {
	fmt.Println("Enter user name")
	_, err = fmt.Scan(&u.User)
	if err != nil {
		log.Println("Error to add name", err)
		return
	}

	fmt.Println("Enter password")
	_, err = fmt.Scan(&u.Pass)
	if err != nil {
		log.Println("Error to add password", err)
		return
	}

	return
}

func LogIn(users []User) (u User, err error) {
	fmt.Println("Enter user name")
	_, err = fmt.Scan(&u.User)
	if err != nil {
		log.Println("Error to find name", err)
		return
	}

	fmt.Println("Enter password")
	_, err = fmt.Scan(&u.Pass)
	if err != nil {
		log.Println("Error to find password", err)
		return
	}

	for _, v := range users {
		if v.User == u.User && v.Pass == u.Pass {
			return
		}
	}

	err = errors.New("User not found")

	return
}

func AddPost() {

}

func EditPost() {

}

func DeletePost() {

}

func ShowAllPosts() {

}

func ShowUserPost() {

}