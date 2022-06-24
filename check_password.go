package main

func checkPassword(user_id, password string) bool {
	result := false
	if _, ok := Data[user_id]; ok {

		if password == Data[user_id].Password {
			result = true
			return result
		}
	}
	return result
}
