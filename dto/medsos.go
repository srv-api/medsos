package dto

import "mime/multipart"

type MedsosRequest struct {
	ID        string                `json:"id"`
	UserID    string                `json:"user_id"`
	Caption   string                `json:"caption"`
	Status    int                   `json:"status"`
	DetailID  string                `json:"detail_id"`
	CreatedBy string                `json:"created_by"`
	ImageURL  string                `form:"image_url"`
	Image     *multipart.FileHeader `form:"image"`
}

type MedsosResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Caption   string `json:"caption"`
	Status    string `json:"status"`
	DetailID  string `json:"detail_id"`
	CreatedBy string `json:"created_by"`
	ImageURL  string `json:"image_url"`
}
