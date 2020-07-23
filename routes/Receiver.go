package routes

import (
	"../models"
	"../pkg/config"
	"../pkg/logs"
	"../pkg/util"
	"encoding/json"
	"fmt"
	"net/http"
)

const receivedCreated = "Action_Created"


// Receiver
func CreateAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Empty Request Body",
			})
		return
	}

	var receiver models.Receiver

	err:= json.NewDecoder(r.Body).Decode(&receiver)
	if err != nil {
		_ = logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "No Data in Request Body",
			})
		return
	}

	// field check
	if receiver.Action == "" || receiver.DocumentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Payload is Empty",
			})
		return
	}

	// switch case on action
	switch receiver.Action {
	case "JobPost":
		// jobPost query and json response
		query:= `Insert into job_posts (job_post_document_id) Value(?);`

		result, err:= config.Database.Exec(fmt.Sprintf(query),
			receiver.DocumentID,
			)
		if err != nil {
			_ = logs.LogError(err)
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 400,
					Reason: "Unable to Create Resource: Additional Info In logs",
				})
			return
		}

		// get last insertID
		lastInsertID, _ := result.LastInsertId()
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(
			util.PDSuccess{
				Status: 201,
				Data:   &util.Data{
					ID:         int(lastInsertID),
					ActionType: receivedCreated,
				},
			})
		break
	case "UserProfile":
		// UserProfile query and json response
		query:= `Insert into users (user_document_id) Value(?);`

		result, err:= config.Database.Exec(fmt.Sprintf(query),
			receiver.DocumentID,
		)
		if err != nil {
			_ = logs.LogError(err)
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 400,
					Reason: "Unable to Create Resource: Additional Info In logs",
				})
			return
		}

		// get last insertID
		lastInsertID, _ := result.LastInsertId()
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(
			util.PDSuccess{
				Status: 201,
				Data:   &util.Data{
					ID:         int(lastInsertID),
					ActionType: receivedCreated,
				},
			})
		break

	default:
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Unable to Create Resource: Additional Info In logs",
			})
		break

	}
}