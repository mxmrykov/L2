package handlers

import (
	"L2/http/cache"
	"L2/http/domain"
	"L2/http/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodPost {
		log.Println("Method error")
		w.WriteHeader(http.StatusMethodNotAllowed)
		details := models.Details{ErrCode: http.StatusMethodNotAllowed, ErrMessage: "Wrong method"}
		response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		w.Write(response)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event
	decodingBodyErr := decoder.Decode(&decoded)

	if decodingBodyErr != nil {
		log.Println(decodingBodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: "Wrong method"}
		response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		w.Write(response)
		return
	}

	dateQuery := decoded.Date
	timeQuery := decoded.Time
	userQuery := decoded.UserId

	_, errParse := time.Parse("2006-01-02", dateQuery)

	if errParse != nil {
		log.Println(errParse)
		w.WriteHeader(http.StatusInternalServerError)
		details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: "Error at date parsing"}
		response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		w.Write(response)
		return
	}

	_, errParseTime := time.Parse("15:00", timeQuery)

	if errParseTime != nil {
		log.Println(errParseTime)
		w.WriteHeader(http.StatusInternalServerError)
		details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: "Error at time parsing"}
		response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		w.Write(response)
		return
		return
	}

	event := domain.NewEvent(dateQuery, timeQuery, userQuery)

	c.Create(*event)
	w.Write([]byte("Successful creation"))
}
