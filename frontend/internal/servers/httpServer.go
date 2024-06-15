package servers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/xhermitx/gitpulse-tracker/frontend/internal/handlers"
	msql "github.com/xhermitx/gitpulse-tracker/frontend/internal/store/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DEFINE THE HOME PAGE
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home Page")
	fmt.Println("Endpoint hit: homepage")
}

// HANDLE THE ROUTES
func handleRequests(handler *handlers.TaskHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/job/create", handler.CreateJob).Methods("POST")
	router.HandleFunc("/job/update", handler.DeleteJob).Methods("POST")
	router.HandleFunc("/job/run", handler.UpdateJob).Methods("GET")

	log.Fatal(http.ListenAndServe(os.Getenv("ADDRESS"), router))
}

func HttpServer() {
	err := godotenv.Load("./../.ENV")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(os.Getenv("DB_SERVER"))

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_SERVER")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to DB")
	}
	mysqlDB := msql.NewMySQLStore(db)
	taskHandler := handlers.NewTaskHandler(mysqlDB)

	handleRequests(taskHandler)
}