package dto

import "mime/multipart"

type MedsosRequest struct {
	ID        string                `json:"id"`
	UserID    string                `json:"user_id"`
	Caption   string                `json:"caption"`
	DetailID  string                `json:"detail_id"`
	CreatedBy string                `json:"created_by"`
	ImageURL  string                `form:"image_url"`
	Image     *multipart.FileHeader `form:"image"`
}

type MedsosResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Caption   string `json:"caption"`
	DetailID  string `json:"detail_id"`
	CreatedBy string `json:"created_by"`
	ImageURL  string `json:"image_url"`
}
