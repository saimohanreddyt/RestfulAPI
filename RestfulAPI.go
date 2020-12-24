package main
import (
	"encoding/json"
	"net/http"
)
type Coaster struct{
	Name 			string 	`json:"name"`
	Manufacturer 	string 	`json:"manufacturer"`
	ID 				string 	`json:"id"`
	InPark 			string 	`json:"inPark"`
	Height 			int 	`json:"heifht"`
}
type coasterHandlers struct {
	store map[string]Coaster
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request){
	coasters := make([]Coaster, len(h.store))
		i := 0
		for _, coaster := range h.store {
			coasters[i] = coaster
			i++
		}
		jsonBytes, err := json.Marshal(coasters)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
}
func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1":Coaster{
				Name: 	"SAI",
				Height:	175,
				ID:		"240",
				InPark:	"RRK",
				Manufacturer: "INDIA",
			},
		},
	}
}
func main() {
	coasterHandler := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandler.get)
	err := http.ListenAndServe(":8084",nil)
	if err != nil {
		panic(err)
	}
}