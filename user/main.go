package main

import (
	userservice "github.com/gdan0324/ByteWeGo/user/kitex_gen/userservice/userservice"
	"log"
)

func main() {
	svr := userservice.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
