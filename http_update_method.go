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

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUid(uid string){
	for index, item := range inventory {
		if item.UID == uid{
			//Delete item from slice
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)
	//Delete the item at UID
	_deleteItemAtUid(params["uid"])
	//Create it with new data
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)

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

	json.NewEncoder(w).Encode(inventory)

}


func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

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