package dto

import "mime/multipart"

type MedsosRequest struct {
	ID        string                `form:"id" json:"id"`
	UserID    string                `form:"user_id" json:"user_id"`
	Caption   string                `form:"caption" json:"caption"`
	DetailID  string                `form:"detail_id" json:"detail_id"`
	CreatedBy string                `form:"created_by" json:"created_by"`
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

type MatchFeedRequest struct {
	UserID string `json:"user_id"`
}
