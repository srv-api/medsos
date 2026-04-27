package medsos

import (
	dto "srv-api/medsos/dto"
)

func (r *medsosRepository) Get(req dto.MatchFeedRequest) ([]dto.MedsosResponse, error) {
	var medsosResponses []dto.MedsosResponse

	query := `
        SELECT id, caption, detail_id, image_url, user_id, created_by
        FROM medsos
        WHERE user_id = ?
        
        UNION ALL
        
        SELECT m.id, m.caption, m.detail_id, m.image_url, m.user_id, m.created_by
        FROM medsos m
        WHERE m.user_id IN (
            SELECT 
                CASE 
                    WHEN matches.user1_id = ? THEN matches.user2_id
                    WHEN matches.user2_id = ? THEN matches.user1_id
                END
            FROM matches
            WHERE (matches.user1_id = ? OR matches.user2_id = ?)
        )
        
        ORDER BY id DESC
    `

	err := r.DB.Raw(query,
		req.UserID,
		req.UserID, req.UserID,
		req.UserID, req.UserID,
	).Scan(&medsosResponses).Error

	if err != nil {
		return nil, err
	}

	return medsosResponses, nil
}
