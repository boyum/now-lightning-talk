package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	year, month, day := currentTime.Date()

	if year == 2020 && month == "April" && day == 30 {
		fmt.Fprintf(w, "<div style=\"display:flex;align-items:center;justify-content:center;min-height:100vh;font-size:50px;text-align:center;\">YES</div>")
		
	}
	else {
		fmt.Fprintf(w, "<div style=\"display:flex;align-items:center;justify-content:center;min-height:100vh;font-size:50px;text-align:center;\">nooooo</div>")
	}
} 