package main

import (
	"fmt"
	"net/http"
	"sample-api/internal/repository"
	"sample-api/internal/router"
	"sample-api/internal/sms"
)

func main() {
	fmt.Println("start on port :9000")
	//init connection
	conn := "tests"

	//init repository
	userDetailRepo := repository.NewUserDetail(conn)

	//init service
	service := sms.NewService(userDetailRepo)

	//init router
	r := router.InitRouter(service)

	//start server
	err := http.ListenAndServe(":9000", r)
	if err != nil {
		panic(err)
	}
}
