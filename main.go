package main

import (
	"fmt"
	"net/http"
	"os"

	"personal_note/app"
	"personal_note/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// TODO add user
	router.HandleFunc("/api/v1/note", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/api/v1/note/{id}", controllers.DeleteNote).Methods("DELETE")
	router.HandleFunc("/api/v1/note/{id}", controllers.UpdateNote).Methods("PUT")
	router.HandleFunc("/api/v1/note/{id}/archive", controllers.ArchiveNote).Methods("PUT")
	router.HandleFunc("/api/v1/note/{id}/unarchive", controllers.UnArchiveNote).Methods("PUT")
	router.HandleFunc("/api/v1/note/archived", controllers.GetArtchivedList).Methods("GET")
	router.HandleFunc("/api/v1/note/unarchived", controllers.GetUnArtchivedList).Methods("GET")

	router.Use(app.BasicAuthentication) //attach Basic auth middleware

	// router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("port : " + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/
	if err != nil {
		fmt.Print(err)
	}
}
