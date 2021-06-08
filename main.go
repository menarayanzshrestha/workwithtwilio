package main

import (
	"fmt"
	"net/http"

	"github.com/menarayanzshrestha/workwithtwilio/utils"
)

func main() {

	utils.LoadEnv()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.SendSms("+9779800000000", "Test")
	})

	PORT := "5000"

	fmt.Println("Serverr running on port:", PORT)

	http.ListenAndServe(":"+PORT, nil)

}
