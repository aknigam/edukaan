package main

import (
	"edukaan/common"
	"edukaan/routers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	//router := mux.NewRouter()
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity

	routers.SetVendorRoutes(router)

	http.Handle("/", router)
	server := http.Server{
		Addr: ":8080",
	}
	/*
		server := &http.Server{
			Addr:           config.Address,
			ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
			WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
			MaxHeaderBytes: 1 << 20,
		}
	*/
	err := server.ListenAndServe()

	if err != nil {
		common.Info.Println("Server not started", err)
	} else {
		common.Info.Println("Server started")
	}

}
