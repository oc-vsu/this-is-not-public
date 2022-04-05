package main

import "opitz-consulting.com/examples/oop/notifier"

func main() {
	user := notifier.User{
		Email: "user@example.com",
	}
	user.Notify()

	admin := notifier.Admin{
		Email:                   "user@example.com",
		SuperSpecificAdminField: "Lorem Ipsum",
	}
	admin.Notify()
}
