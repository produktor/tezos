package handler

import "link_api/domain/model"

type GroupsResponse struct {
	TelegramGroups []TelegramGroupResponse `json:"telegramGroups"`
}

type TelegramGroupResponse struct {
	ID               int64   `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	CriteriaType     *string `json:"criteriaType"`
	CriteriaToken    *string `json:"criteriaToken"`
	CriteriaCurrency *string `json:"criteriaCurrency"`
	CriteriaPrice    float64 `json:"criteriaPrice"`
}

func (TelegramGroupResponse) FromModel(m model.TelegramGroup) TelegramGroupResponse {
	price, _ := m.CriteriaPrice.Float64()

	return TelegramGroupResponse{
		ID:               m.ID,
		Title:            *m.Title,
		Description:      *m.Description,
		CriteriaType:     &m.CriteriaType,
		CriteriaToken:    m.CriteriaToken,
		CriteriaCurrency: m.CriteriaCurrency,
		CriteriaPrice:    price,
	}
}
