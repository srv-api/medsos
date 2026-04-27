package medsos

import (
	dto "srv-api/medsos/dto"
)

func (r *medsosRepository) Get(req dto.MatchFeedRequest) ([]dto.MedsosResponse, error) {
	var medsosResponses []dto.MedsosResponse

	query := `
        SELECT m.id, m.caption, m.detail_id, m.image_url, m.user_id, m.created_by
        FROM medsos m
        WHERE m.user_id IN (
            SELECT 
                CASE 
                    WHEN user1_id = ? THEN user2_id
                    WHEN user2_id = ? THEN user1_id
                END
            FROM matches
            WHERE (user1_id = ? OR user2_id = ?)
        )
        AND m.user_id != ?
    `

	err := r.DB.Raw(query,
		req.UserID, req.UserID, // untuk CASE
		req.UserID, req.UserID, // untuk WHERE clause
		req.UserID, // untuk AND m.user_id != ?
	).Scan(&medsosResponses).Error

	if err != nil {
		return nil, err
	}

	return medsosResponses, nil
}
