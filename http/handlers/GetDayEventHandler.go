package handlers

import (
	"L2/http/cache"
	"L2/http/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func GetDayEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodGet {
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

	_, errParse := time.Parse("2006-01-02", dateQuery)

	if errParse != nil {
		log.Println(errParse)
		w.Write([]byte("Error errParse"))
		return
	}

	if val, ok := c.ReadDay(dateQuery); ok {
		response, erMarshalingResponse := json.MarshalIndent(val, "", "\t")

		if erMarshalingResponse != nil {
			log.Println(erMarshalingResponse)
			w.Write([]byte("Error erMarshalingResponse"))
			return
		}

		w.Write(response)
	} else {
		log.Println(errParse)
		w.Write([]byte("Error getting date"))
		return
	}
}
