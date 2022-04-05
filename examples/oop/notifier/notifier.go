package notifier

import (
	"fmt"
)

type Notifier interface {
	Notify()
}

type User struct {
	Email string
}

type Admin struct {
	Email                   string
	SuperSpecificAdminField string
}

func (u User) Notify() {
	fmt.Printf("notifying user with email: %s\n", u.Email)
}

func (a Admin) Notify() {
	fmt.Printf("notifying admin with email: %s (specific Field: %s)\n", a.Email, a.SuperSpecificAdminField)
}
