package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	danboDayIII := time.Date(2020, 4, 30, 0, 0, 0, 0, time.UTC)
	danboYear, danboMonth, danboDay := danboDayIII.Date()

	currentTime := time.Now()
	currentYear, currentMonth, currentDay := currentTime.Date()

	itIsDanboDay := danboYear == currentYear && danboMonth == currentMonth && danboDay == currentDay

	if itIsDanboDay { 
		fmt.Fprintf(w, "<div style=\"display:flex;align-items:center;justify-content:center;min-height:100vh;font-size:50px;text-align:center;\">YES, <code>%v</code> is Danbo Day III</div>", currentTime.Format(time.RFC850))		
	} else {
		fmt.Fprintf(w, "<div style=\"display:flex;align-items:center;justify-content:center;min-height:100vh;font-size:50px;text-align:center;\">nooooo</div>")
	}
} 