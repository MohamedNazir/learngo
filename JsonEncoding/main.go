package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name       string     `json:"userName"`
	Password   string     `json:"-"`
	Permission permission `json:",omitempty"` // evenif you don't need to change the name, you need to have the comma :(
}

type permission map[string]bool

func main() {

	users := []user{
		{"riz", "123", nil},
		{"angel", "hahaha", permission{"admin": true}},
		{"devil", "xxxx", permission{"write": true}},
	}

	//	out, err := json.Marshal(users)
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("Error ::", err)
		return
	}
	fmt.Println(string(out))
}
