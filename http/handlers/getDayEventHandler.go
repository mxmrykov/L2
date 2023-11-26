package handlers

import (
	"L2/http/cache"
	"L2/http/domain"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func GetDayEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodGet {
		domain.ErrorLogger(w, errors.New("method error"))
		return
	}

	dateQuery := r.URL.Query().Get("date")

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		domain.ErrorLogger(w, errParse)
		return
	}

	if val, ok := c.ReadDay(dateQuery); ok {
		if response, erMarshalingResponse := json.MarshalIndent(val, "", "\t"); erMarshalingResponse != nil {
			domain.ErrorLogger(w, erMarshalingResponse)
			return
		} else {
			w.Write(response)
			return
		}
	} else {
		domain.ErrorLogger(w, errors.New("error at date parsing"))
		return
	}
}
