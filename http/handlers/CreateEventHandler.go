package handlers

import (
	"L2/http/cache"
	"L2/http/domain"
	"L2/http/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		details := models.Details{ErrCode: http.StatusMethodNotAllowed, ErrMessage: "Wrong method"}
		response, errMarshaling := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		if errMarshaling != nil {
			log.Println(errMarshaling)
			w.Write([]byte("Error erMarshalingResponse"))
			return
		}
		w.Write(response)
		return
	}

	dateQuery := r.URL.Query().Get("date")
	timeQuery := r.URL.Query().Get("time")
	userQuery := r.URL.Query().Get("user_id")

	_, errParse := time.Parse("2006-01-02", dateQuery)

	if errParse != nil {
		log.Println(errParse)
		w.Write([]byte("Error errParse"))
		return
	}

	_, errParseTime := time.Parse("15:00", timeQuery)

	if errParseTime != nil {
		log.Println(errParseTime)
		w.Write([]byte("Error errParseTime"))
		return
	}

	userInt, errParseUser := strconv.Atoi(userQuery)

	if errParseUser != nil {
		log.Println(errParseUser)
		w.Write([]byte("Error errParseUser"))
		return
	}

	event := domain.NewEvent(dateQuery, timeQuery, userInt)

	c.Create(*event)
	w.Write([]byte("Successful creation"))
}
