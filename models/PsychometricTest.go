package models

type PsychometricTest struct {
	ID int `json:"id"`
	UserID string `json:"user_id"`
	Result int `json:"result, omitempty"`
	PsychometricTestDocumentID string `json:"psychometric_test_document_id, omitempty"`
}

type PsychometricTests []PsychometricTest

type GetOnePsychometricTest struct {
	Status int `json:"status"`
	Data *PsychometricTest `json:"data"`
}