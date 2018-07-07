package main

import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

type DietDay struct {
	Date    time.Time `json:"date"`
	Choice	string `json:"choice"`
}

var dietDays []DietDay

func getDietDaysHandler(w http.ResponseWriter, r *http.Request) {
	dietDaysBytes, err := json.Marshal(dietDays)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(dietDaysBytes)
}

func createDietDayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you hit a thing")
	dietDay := DietDay{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dietDay.Choice = r.Form.Get("choice")
	dietDay.Date = time.Now() // r.Form.Get("date") // this probably has to be typed

	// Get the information about the bird from the form info
	//bird.Species = r.Form.Get("species")
	//bird.Description = r.Form.Get("description")

	// Append our existing list of birds with a new entry
	//birds = append(birds, bird)
	dietDays = append(dietDays, dietDay)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
