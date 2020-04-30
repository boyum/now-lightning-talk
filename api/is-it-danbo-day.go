package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	danboDayIII := time.Date(2020, 4, 30, 0, 0, 0, 0, time.UTC)
	currentTime := time.Now()

	if currentTime == danboDayIII { 
		fmt.Fprintf(w, "<div style=\"display:flex;align-items:center;justify-content:center;min-height:100vh;font-size:50px;text-align:center;\">YES</div>")		
	} else {
		fmt.Fprintf(w, "<div style=\"display:flex;align-items:center;justify-content:center;min-height:100vh;font-size:50px;text-align:center;\">nooooo</div>")
	}
} 