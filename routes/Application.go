package routes

import (
	"../models"
	"../pkg/config"
	"../pkg/logs"
	"../pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const applicationCreated  = "Application_Created"
const applicationUpdated  = "Application_Updated"
const applicantDeleted    = "Application_Deleted"

// CRUD FUNCTIONS

//create an application i.e apply to a job
func CreateApplication(w http.ResponseWriter, r *http.Request) {
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

	var application models.Application

	err := json.NewDecoder(r.Body).Decode(&application)
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
	if application.UserID == 0 || application.PostID == 0 || application.Title == "" || application.JobPostDocumentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Empty Request Body",
			})
		return
	}

	query := `Insert into applications (userID, postID, title, job_post_document_id) Values(?, ?, ?, ?);`

	result, err:= config.Database.Exec(fmt.Sprintf(query),
		application.UserID,
		application.PostID,
		application.Title,
		application.JobPostDocumentID,
	)
	if err != nil {
		_ = logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Unable to Create Resource",
			})
		return
	}
	lastInsertID, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(util.PDSuccess{
		Status: 201,
		Data:   &util.Data{
			ID:         int(lastInsertID),
			ActionType: applicationCreated,
		},
	})
	//todo:: send user_id and job_post_document_id to python server
	return
}

//get all user applications made
func GetAllUserApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r) // access id param in url

	userID, err := util.ConvertStringToInt(vars["id"])
	if err != nil {
		logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Invalid ID",
			})
	} else {
		fmt.Println(userID)

		query := `Select * from applications where userID = ?`

		applications := models.Applications{}

		result, err:= config.Database.Query(query, userID)
		if err != nil {
			_ = logs.LogError(err)
			_ = logs.LogToFile(err.Error(), "api-errors.txt")
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "Could not get list of user applications",
				})
			return
		}
		defer result.Close()
		application := models.Application{}
		for result.Next() {
			err := result.Scan(
				&application.ID,
				&application.UserID,
				&application.PostID,
				&application.Title,
				&application.JobPostDocumentID,
				&application.ApplicationDocumentID,
			)
			if err != nil {
				_ = logs.LogError(err)
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(
					util.Fail{
						Status: 404,
						Reason: "Could not get list of user applications",
					})
				return
			}
			applications = append(applications, application)
		}
		err = result.Err()
		if err != nil {
			_ = logs.LogError(err)
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "I wonder why we have this one",
				})
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(
				models.GetAllApplications{
					Status: 200,
					Data:   &applications,
				})
	}
}


//get all applications made per job_post_document_id
func GetAllApplicationsSubmitted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r) // access id param in url

	jobPostDocumentID := vars["id"]

	// some random check?
	// don't know if this is necessary
	if jobPostDocumentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Cannot Have Empty ID params",
			})
	} else {
		fmt.Println(jobPostDocumentID)

		query := `Select * from applications where job_post_document_id = ?`

		applications := models.Applications{}

		result, err:= config.Database.Query(query, jobPostDocumentID)
		if err != nil {
			_ = logs.LogError(err)
			_ = logs.LogToFile(err.Error(), "api-errors.txt")
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "Could not get list of submitted applications",
				})
			return
		}
		defer result.Close()

		application := models.Application{}
		for result.Next() {
			err := result.Scan(
				&application.ID,
				&application.UserID,
				&application.PostID,
				&application.Title,
				&application.JobPostDocumentID,
				&application.ApplicationDocumentID,
			)
			if err != nil {
				_ = logs.LogError(err)
				w.WriteHeader(http.StatusNotFound)
				_ = json.NewEncoder(w).Encode(
					util.Fail{
						Status: 404,
						Reason: "Could not get list of submitted user applications",
					})
				return
			}
			applications = append(applications, application)
		}
		err = result.Err()
		if err != nil {
			_ = logs.LogError(err)
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "I wonder why we have this one",
				})
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(
			models.GetAllApplications{
				Status: 200,
				Data:   &applications,
			})
	}
}