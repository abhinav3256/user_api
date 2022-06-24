package main

import "fmt"

func UniqueEmail(email string) bool {
	message := false

	user := User{}
	for _, value := range Data {

		//check if present value is equals to userValue
		if value == user {

			//if same return true
			return true
		}
	}

	fmt.Println(user)

	return message
}
