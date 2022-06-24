package main

import "fmt"

func delete_user(userID string) string {
	message := "User not Found"

	user := User{}
	if user, ok := Data[userID]; ok {
		delete(Data, userID)
		message = "user Deleted"
		fmt.Println(user)

		return message
	}
	fmt.Println(user)

	return message
}
