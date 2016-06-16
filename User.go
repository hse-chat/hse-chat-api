package main

import "github.com/hse-chat/hse-chat-api/HseMsg"

// User struct represents user in system
type User struct {
	Username string
	Password string
	Online   bool
}

// ToProtoMessage convers struct to *HseMsg.User
func (usr User) ToProtoMessage() *HseMsg.User {
	return &HseMsg.User{
		Username: &usr.Username,
		Online:   &usr.Online,
	}
}
