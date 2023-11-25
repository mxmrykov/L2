package handlers

import (
	"L2/http/cache"
	"L2/http/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func GetWeekEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodGet {
		log.Println("Method error")
		w.WriteHeader(http.StatusInternalServerError)
		details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: "Wrong method"}
		response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		w.Write(response)
		return
	}

	dateQuery := r.URL.Query().Get("date")

	_, errParse := time.Parse("2006-01-02", dateQuery)

	if errParse != nil {
		log.Println(errParse)
		w.WriteHeader(http.StatusInternalServerError)
		details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: "Error at date parsing"}
		response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		w.Write(response)
		return
	}

	if val, ok := c.ReadWeek(dateQuery); ok {
		response, erMarshalingResponse := json.MarshalIndent(val, "", "\t")

		if erMarshalingResponse != nil {
			log.Println(erMarshalingResponse)
			w.WriteHeader(http.StatusInternalServerError)
			details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: "Error at marshaling response"}
			response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
			w.Write(response)
			return
		}

		w.Write(response)
	} else {
		log.Println(errParse)
		w.WriteHeader(http.StatusServiceUnavailable)
		details := models.Details{ErrCode: http.StatusServiceUnavailable, ErrMessage: "Error at date parsing"}
		response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
		w.Write(response)
		return
	}
}
