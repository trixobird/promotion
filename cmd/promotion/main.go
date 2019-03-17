package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"promotion/models"
	"promotion/restapi"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/register", restapi.RegisterPhone).Methods("POST")
	router.HandleFunc("/api/sms-promotion", restapi.SendPromoCode).Methods("POST")
	router.HandleFunc("/api/redeem-promotion", restapi.RedeemPromotion).Methods("POST")

	port := os.Getenv("server_port")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	defer models.ReleaseDb()
}
