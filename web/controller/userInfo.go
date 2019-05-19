package controller

import "github.com/chenwbyx/Fabric-Traceability/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName	string
	Password	string
	IsAdmin	    string
}


var users []User

func init() {

	admin := User{LoginName:"admin", Password:"123456", IsAdmin:"T"}
	wenbo := User{LoginName:"wenbo", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, wenbo)

}