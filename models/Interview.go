package models

type Interview struct {
	ID int `json:"id"`
	Title string `json:"title"`
	CreatedOn string `json:"created_on"`
	InterviewDocumentID string `json:"interview_document_id, omitempty"`
}

type Interviews []Interview

type GetOneInterview struct {
	Status int `json:"status"`
	Data *Interview `json:"data"`
}