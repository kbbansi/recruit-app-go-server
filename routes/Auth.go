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

func Auth(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	var auth models.Auth
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Empty Request Body",
			})
		return
	}

	err:=json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		logs.LogError(err)
		logs.LogToFile(err.Error(), util.API_ERRORS)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Empty Request Body",
			})
		return
	}

	if auth.Email == "" || auth.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 400,
				Reason: "Empty Request Object",
			})
		return
	}

	//todo:: Compare passwords
	user:= models.User{}

	checkPassword:= `Select password from users where email = ?;`
	err = config.Database.QueryRow(fmt.Sprintf(checkPassword), auth.Email).Scan(&user.Password)

	// let's see what's happening
	fmt.Println(string(auth.Password), " -> ", string(user.Password))

	/*Careful implementation of auth begins from here*/
	password:= util.CheckPasswordHash(auth.Password, user.Password)

	if password == true {
		//login query
		query:= `Select count(*) from users where email = ? and password = ?;`

		err = config.Database.QueryRow(fmt.Sprintf(query), auth.Email, user.Password).Scan(&util.Count)

		if util.Count != 0 {
			_, err:= config.Database.Exec(fmt.Sprintf(query), auth.Email, user.Password)
			if err != nil {
				logs.LogError(err)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(
					util.Fail{
						Status: 400,
						Reason: "Bad Request",
					})
				return
			}
			//
			user, _:= GetUserDetails(auth.Email)
			fmt.Println(user)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(
				models.AuthSuccess{
					Status:      200,
					Message:     "User Logged In",
					SessionData: &user,
				})
			return
		} else {
			//count greater 0
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				util.Fail{
					Status: 400,
					Reason: "Invalid Login Credentials",
				})
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			util.Fail{
				Status: 404,
				Reason: "User Not Found",
			})
		return
	}
	/*Careful implementation of auth ends here*/
}

func GetUserDetails(email string) (user models.User, err error) {
	query:= `Select id, first_name, last_name, other_names, contact, 
				email, user_type from users where email = ?;`

	err = config.Database.QueryRow(fmt.Sprintf(query), email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.OtherNames,
		&user.Contact,
		&user.Email,
		&user.UserType,
	)

	if err != nil {
		logs.LogError(err)
		return user, err
	}
	return user, nil
}