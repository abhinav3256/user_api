package main

import "fmt"

func isUnique(userID string) bool {
	message := false

	user := User{}
	if user, ok := Data[userID]; ok {

		message = true
		fmt.Println(user)

		return message
	}
	fmt.Println(user)

	return message
}
