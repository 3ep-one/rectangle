package httphandler

import (
	// "bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/3ep-one/rectangle/rectanglesolver"
	"github.com/3ep-one/rectangle/rediswraper"
	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	redisClient := rediswraper.Makeredisclient()
	rectList := rediswraper.Getkeyvalue(redisClient)
	var rectClean []rectanglesolver.Rectangle
	for _, rect := range rectList {
		var ans rectanglesolver.Rectangle
		err := json.Unmarshal([]byte(rect), &ans)
		if err != nil {
			panic(err)
		}
		rectClean = append(rectClean, ans)
	}
	jsn, _ := json.Marshal(rectClean)
	w.Write(jsn)
}

func post(w http.ResponseWriter, r *http.Request) {
	p := rectanglesolver.Jsoninput{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal(err)
	}
	rectanglesolver.Haveoverlap(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Input recieved."}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

//Handler handle http request
func Handler() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
