package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"error/dao"
	"error/service"
)

func main() {
	err := dao.OpenDB()
	if err != nil {
		fmt.Printf("original error: %T, %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		os.Exit(1)
	}
	defer dao.Close()

	http.HandleFunc("/getuserbyip", getUserByIP)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getUserByIP(w http.ResponseWriter, r *http.Request) {
	service.FindUserByIP(w, r)
}
