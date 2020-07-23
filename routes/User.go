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

const userCreated = "USER_CREATED"
const userDeleted = "USER_DELETED"
const userUpdated = "USER_UPDATED"

// User CRUD functions

//GetAllUsers
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	query := `Select id, first_name, last_name, other_names, contact, email, password, 
				user_type, created_on, coalesce(modified_on, "") AS modified_on from users;`

	rows, err := config.Database.Query(query)
	if err != nil {
		err = logs.LogError(err)
		err = logs.LogToFile(err.Error(), "api-errors.txt")
		res.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(res).Encode(
			util.Fail{
				Status: 404,
				Reason: "Could not get list of users",
			},
		)
		return
	}
	defer rows.Close()

	users := models.Users{}
	user := models.User{}
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.OtherNames,
			&user.Contact,
			&user.Email,
			&user.Password,
			&user.UserType,
			&user.CreatedOn,
			&user.ModifiedOn,
		)
		if err != nil {
			err = logs.LogError(err)
			res.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(res).Encode(
				util.Fail{
					Status: 404,
					Reason: "Could not Get list of Users",
				},
			)
			return
		}
		user.CreatedOn, err = util.ConvertDateTimeStringToUTCString(util.CreatedOn)
		if err != nil {
			err = logs.LogError(err) // this might be troublesome... let's see
		}
		user.ModifiedOn, err = util.ConvertDateTimeStringToUTCString(util.ModifiedOn)
		if err != nil {
			logs.LogError(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		logs.LogError(err)
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(
			util.Fail{
				Status: 404,
				Reason: "I wonder why we have this one??",
			},
		)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(
		models.GetAllUsers{
			Status: 200,
			Data:   &users,
		},
	)
}

func GetOneUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	query := `Select id, first_name, last_name, other_names, contact, email, password, 
				user_type, created_on, coalesce(modified_on, "") AS modified_on, coalesce(user_document_id, "") As documentID from users
				where id = ?;`
	// id will be the query string param for an id
	vars := mux.Vars(req)

	//check for integer value
	id, err := util.ConvertStringToInt(vars["id"])
	if err != nil {
		logs.LogError(err)
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(
			util.Fail{
				Status: 400,
				Reason: "Invalid ID",
			})
	} else {
		user := models.User{}
		user.CreatedOn, _ = util.ConvertDateTimeStringToUTCString(util.CreatedOn)
		user.ModifiedOn, _ = util.ConvertDateTimeStringToUTCString(util.ModifiedOn)

		err = config.Database.QueryRow(fmt.Sprintf(query), id).Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.OtherNames,
			&user.Contact,
			&user.Email,
			&user.Password,
			&user.UserType,
			&user.CreatedOn,
			&user.ModifiedOn,
			&user.UserDocumentID,
		)

		if err != nil {
			logs.LogError(err)
			res.WriteHeader(http.StatusNotFound)
			json.NewEncoder(res).Encode(
				util.Fail{
					Status: 404,
					Reason: "No such user",
				})
		} else {
			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode(
				models.GetOneUser{
					Status: 200,
					Data:   &user,
				})
		}
	}
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Empty Request Body",
			})
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "No Data in Request Body",
			})
		return
	}

	// field check
	if user.FirstName == "" || user.LastName == "" || user.Contact == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "No Data Supplied",
			})
		return
	}

	if VerifyEmail(user.Email) == false {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Email exists",
			})
		return
	} else {
		//hash password
		password, err := util.HashPassword(user.Password)
		if err != nil {
			logs.LogError(err)
		}
		// check to see what's happening
		fmt.Println(password)

		query := `Insert into users (first_name, last_name, other_names, contact, email, password, created_on) 
					Values(?, ?, ?, ?, ?, ?, NOW());`

		queryResult, err := config.Database.Exec(fmt.Sprintf(query),
			user.FirstName,
			user.LastName,
			user.OtherNames,
			user.Contact,
			user.Email,
			password,
		)
		if err != nil {
			_ = logs.LogError(err)
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 400,
					Reason: "Unable to create Resource",
				})
			return
		}
		lastInsertID, _ := queryResult.LastInsertId()
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(
			util.PDSuccess{
				Status: 201,
				Data: &util.Data{
					ID:         int(lastInsertID),
					ActionType: userCreated,
				},
			})
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars:= mux.Vars(r)

	userID, err:= util.ConvertStringToInt(vars["id"])
	if err != nil {
		logs.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Invalid ID",
			})
		return
	} else {
		//check for nil body
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				util.Fail{
					Status: 400,
					Reason: "Empty Request Body",
				})
			return
		}
		var user models.User
		err:= json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			logs.LogError(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				util.Fail{
					Status: 404,
					Reason: "No Data",
				})
			return
		}

		//check for empty fields
		if user.FirstName == "" || user.LastName == "" || user.OtherNames == "" || user.Contact == "" || user.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				util.Fail{
					Status: 400,
					Reason: "No Data Supplied",
				})
			return
		}

		query:= `Update users Set first_name = ?, last_name = ?, other_names = ?, contact = ?, 
					email = ?, user_type = ?, modified_on = NOW() where id = ?;`
		res, err:= config.Database.Exec(fmt.Sprintf(query),
			user.FirstName,
			user.LastName,
			user.OtherNames,
			user.Contact,
			user.Email,
			user.UserType,
			userID,
		)

		if err != nil {
			_ = logs.LogError(err)
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(
				util.Fail{
					Status: 400,
					Reason: "Could not Update User",
				})
			return
		}

		//get updated user info
		userUpdate, err:= GetUser(userID)
		lastInsertID, _:= res.RowsAffected()
		w.WriteHeader(http.StatusOK)
		err =json.NewEncoder(w).Encode(
			models.UserUpdate{
				Status:       200,
				Message:      userUpdated,
				LastInsertID: int(lastInsertID),
				Data:         &userUpdate,
			})
	}
}

//helper func:: VerifyEmail will check to see if an email exist in the users table
func VerifyEmail(email string) bool {
	fmt.Println(email)
	query := `Select count(*) from users where email = ?;`
	err := config.Database.QueryRow(fmt.Sprintf(query), email).Scan(&util.Count)
	if err != nil {
		err = logs.LogError(err)
		return false // some error occurred
	}

	if util.Count == 0 {
		fmt.Println("User Email Does Not Exist")
		return true // user email does not exist, programme can continue to execute
	}

	return false // email exists
}

//GetUser will obtain a user
func GetUser(id int) (user models.User, err error) {
	query:= `Select id, first_name, last_name, other_names, contact, 
				email, user_type, user_document_id from users where id = ?;`
	err = config.Database.QueryRow(fmt.Sprintf(query), id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.OtherNames,
		&user.Contact,
		&user.Email,
		&user.UserType,
		&user.UserDocumentID,
	)

	if err != nil {
		logs.LogError(err)
		return user, err
	}
	return user, nil
}