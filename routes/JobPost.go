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

const jobPostCreated = "JOB_POST_CREATED"
const jobPostDeleted = "JOB_POST_DELETED"
const jobPostUpdated = "JOB_POST_UPDATED"

// Job Post CRUD functions

// Create Job Post
func CreateJobPost(w http.ResponseWriter, r *http.Request) {
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

	var jobPost models.JobPost

	err := json.NewDecoder(r.Body).Decode(&jobPost)
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
	if jobPost.Title == "" || jobPost.Organization == "" || jobPost.CreatedBy == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Payload is Empty",
			})
		return
	}

	query := `Insert into job_posts (title, organization, created_by) Values(?, ?, ?);`

	result, err:= config.Database.Exec(fmt.Sprintf(query),
		jobPost.Title,
		jobPost.Organization,
		jobPost.CreatedBy,
	)
	if err != nil {
		_ = logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Unable to Create Resource: Additional info in logs",
			})
		return
	}
	lastInsertID, _ := result.LastInsertId()
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(
		util.PDSuccess{
			Status: 201,
			Data:   &util.Data{
				ID:         int(lastInsertID),
				ActionType: jobPostCreated,
			},
		})
	//todo:::
	return
}

// Get all Job Posts
func GetAllJobPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `Select id, title, organization, created_by, coalesce(job_post_document_id, "") as JobDocumentID from job_posts;`

	rows, err := config.Database.Query(query)
	if err != nil {
		err = logs.LogError(err)
		err = logs.LogToFile(err.Error(), "api-errors.txt")
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 404,
				Reason: "Could Not Get List of Job Posts",
			})
		return
	}

	defer rows.Close()

	jobPosts := models.JobPosts{}
	jobPost := models.JobPost{}
	for rows.Next() {
		err := rows.Scan(
			&jobPost.ID,
			&jobPost.Title,
			&jobPost.Organization,
			&jobPost.CreatedBy,
			&jobPost.JobPostDocumentID,
		)

		if err != nil {
			err = logs.LogError(err)
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "Could Not Get List of Job Posts",
				})
			return
		}
		jobPosts = append(jobPosts, jobPost)
	}
	err = rows.Err()
	if err != nil {
		err = logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 404,
				Reason: "Bad Request",
			})
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(
		models.GetAllJobPost{
			Status: 200,
			Data:   &jobPosts,
		})
	return
}

// Get all Job Posts
func GetAllJobPostsByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars:= mux.Vars(r)

	id, err:= util.ConvertStringToInt(vars["id"])
	if err != nil {
		err = logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Invalid ID",
			})
	} else {
		query := `Select id, title, organization, created_by, coalesce(job_post_document_id, "")
    			as JobDocumentID from job_posts where created_by = ?;`

		rows, err := config.Database.Query(query, id)
		if err != nil {
			err = logs.LogError(err)
			err = logs.LogToFile(err.Error(), "api-errors.txt")
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "Could Not Get List of Job Posts",
				})
			return
		}

		defer rows.Close()

		jobPosts := models.JobPosts{}
		jobPost := models.JobPost{}
		for rows.Next() {
			err := rows.Scan(
				&jobPost.ID,
				&jobPost.Title,
				&jobPost.Organization,
				&jobPost.CreatedBy,
				&jobPost.JobPostDocumentID,
			)

			if err != nil {
				err = logs.LogError(err)
				w.WriteHeader(http.StatusNotFound)
				err = json.NewEncoder(w).Encode(
					util.Fail{
						Status: 404,
						Reason: "Could Not Get List of Job Posts",
					})
				return
			}
			jobPosts = append(jobPosts, jobPost)
		}
		err = rows.Err()
		if err != nil {
			err = logs.LogError(err)
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "Bad Request",
				})
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(
			models.GetAllJobPost{
				Status: 200,
				Data:   &jobPosts,
			})
		return
	}
}

// Get one Job Post
func GetOneJobPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `Select id, title, organization, created_by,
				coalesce(job_document_id, "") as 
				JobDocumentID from job_posts where job_document_id = ?`

	vars:= mux.Vars(r)

	jobPostDocumentID:= vars["document_id"]

	if jobPostDocumentID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Cannot Have Empty document ID",
			})
	} else {
		fmt.Println(jobPostDocumentID)

		jobPost:= models.JobPost{}

		err:= config.Database.QueryRow(fmt.Sprintf(query), jobPostDocumentID).Scan(
			&jobPost.ID,
			&jobPost.Title,
			&jobPost.Organization,
			&jobPost.CreatedBy,
			&jobPost.JobPostDocumentID,
		)

		if err != nil {
			err = logs.LogError(err)
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "Could Not Get Job Post",
				})
		} else {
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(
				models.GetOneJobPost{
					Status: 200,
					Data:   &jobPost,
				})
		}
	}
}