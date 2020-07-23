package models

type Application struct {
	ID                    int    `json:"id"`
	UserID                int 	 `json:"userID"`
	PostID				  int 	 `json:"postID`
	Title				  string  `json:"title"`	
	JobPostDocumentID     string `json:"job_post_document_id"`
	ApplicationDocumentID string `json:"application_document_id, omitempty"`
}

type Applications []Application

type GetOneApplication struct {
	Status int          `json:"status"`
	Data   *Application `json:"data"`
}

type GetAllApplications struct {
	Status int `json:"status"`
	Data   *Applications `json:"data"`
}

type ApplicationUpdate struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *Application `json:"data"`
}
