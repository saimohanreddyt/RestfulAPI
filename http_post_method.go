package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Item struct{
	UID 	string `json:"UID"`
	Name 	string	`json:"Name"`
	Desc 	string	`json:"Desc"`
	Price	float64	`json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage()")

}

func getInventory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Println("Function Called:getInventory()")

	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var item Item
	 _ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080",router))
}

func main() {
	inventory = append(inventory, Item{
		UID:	"240",
		Name:	"Tom Cruise",
		Desc:	"He is an American Actor",
		Price:	40.5,
	})
	inventory = append(inventory, Item{
		UID:	"241",
		Name:	"Json Statham",
		Desc:	"He is an American Actor",
		Price:	40.3,
	})
	handleRequests()
}