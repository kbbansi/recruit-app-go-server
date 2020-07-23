package models

type JobPost struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	Organization      string `json:"organization"`
	CreatedBy		  int	 `json:"created_by"`
	JobPostDocumentID string `json:"job_post_document_id, omitempty"`
}

type JobPosts []JobPost

type GetOneJobPost struct {
	Status int      `json:"status"`
	Data   *JobPost `json:"data"`
}

type GetAllJobPost struct {
	Status int      `json:"status"`
	Data   *JobPosts `json:"data"`
}

type JobPostUpdate struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    *JobPost `json:"data"`
}
