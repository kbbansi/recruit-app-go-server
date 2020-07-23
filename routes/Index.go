package routes

import (
	"../models"
	"encoding/json"
	"net/http"
	"time"
)
const service = "RecruitApp-Go"
func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(models.Alive{
		Alive:   true,
		Service: service,
		Date:    time.Now().Weekday(),
	})
}
