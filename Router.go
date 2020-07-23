package main

import (
	"./models"
	"./pkg/logs"
	"./routes"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter is the router for the api
func NewRouter() *mux.Router {
	router:= mux.NewRouter().StrictSlash(true)
	for _, route := range routing {
		var handler http.Handler
		handler = route.HandlerFunction
		handler = logs.Log(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var routing = models.Routes{
	//Index route
	models.Route{
		Name:            "Index",
		Method:          http.MethodGet,
		Pattern:         "/",
		HandlerFunction: routes.Index,
	},
	
	//User routes
	models.Route{Name: "CreateNewUser", Method:http.MethodPost, Pattern: "/user", HandlerFunction: routes.CreateNewUser},
	models.Route{Name: "GetOneUser", Method:http.MethodGet, Pattern: "/user/{id}", HandlerFunction: routes.GetOneUser},
	models.Route{Name: "GetAllUsers", Method:http.MethodGet, Pattern: "/users", HandlerFunction: routes.GetAllUsers},
	models.Route{Name: "UpdateUser", Method:http.MethodPut, Pattern: "/user/{id}", HandlerFunction: routes.UpdateUser},

	//Auth route
	models.Route{Name: "Auth", Method:http.MethodPost, Pattern: "/user/auth", HandlerFunction: routes.Auth},

	//RecruitApp Routes
	// Applications,
	models.Route{Name:"Apply to Job(CreateApplication)", Method:http.MethodPost, Pattern:"/application/create", HandlerFunction:routes.CreateApplication},
	models.Route{Name:"Get User Applications(GetUserApplications)", Method:http.MethodGet, Pattern:"/application/user/id", HandlerFunction:routes.GetAllUserApplications},
	models.Route{Name:"Get Submitted Applications(GetAllApplicationsSubmitted)", Method:http.MethodGet, Pattern:"/application/submitted/id", HandlerFunction:routes.GetAllApplicationsSubmitted},

	//JobPosts,
	models.Route{Name: "GetAllJobPosts", Method:http.MethodGet, Pattern:"/jobPosts", HandlerFunction: routes.GetAllJobPosts},
	models.Route{Name: "GetAllJobPostsByUser", Method:http.MethodGet, Pattern: "/jobPosts/{id}", HandlerFunction: routes.GetAllJobPostsByUser},
	models.Route{Name:"Create Job Post", Method:http.MethodPost, Pattern:"/jobPost/create", HandlerFunction:routes.CreateJobPost},

	//Receiver for updating document IDs
	models.Route{Name:"CreateDocumentID", Method:http.MethodPost, Pattern:"/receiver", HandlerFunction: routes.CreateAction},
}