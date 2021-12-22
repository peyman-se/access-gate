package main

import (
	"fmt"
	gate "github.com/peyman-se/access-gate"
)

type User struct {
	ID int
}

type Profile struct {
	ID     int
	UserID int
}

func main() {
	g := gate.NewGate()
	user := &User{ID: 1}
	profile := &Profile{ID: 1, UserID: 1}

	g.Define("access-profile", func(args ...interface{}) bool {
		user := *args[0].(*User)
		profile := *args[1].(*Profile)
		return user.ID == profile.UserID
	})

	v := g.Can("access-profile", user, profile)
	if v {
		fmt.Println("can access profile")
	} else {
		fmt.Println("Access denied")
	}

}
