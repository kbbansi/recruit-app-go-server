package models

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	OtherNames string `json:"other_names, omitempty"`
	Contact    string `json:"contact"` // telephone
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserType   string `json:"user_type"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	UserDocumentID string `json:"user_document_id, omitempty"`
}

type Users []User

type GetOneUser struct {
	Status int `json:"status"`
	Data   *User  `json:"data"`
}

type GetAllUsers struct {
	Status int `json:"status"`
	Data   *Users `json:"data"`
}

type UserUpdate struct {
	Status int `json:"status"`
	Message string `json:"message"`
	LastInsertID int `json:"last_insert_id"`
	Data *User `json:"data"`
} 
