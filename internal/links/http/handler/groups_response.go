package handler

import "link_api/domain/model"

type GroupsResponse struct {
	TelegramGroups []TelegramGroupResponse `json:"telegramGroups"`
}

type TelegramGroupResponse struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (TelegramGroupResponse) FromModel(m model.TelegramGroup) TelegramGroupResponse {
	price, _ := m.Price.Float64()

	return TelegramGroupResponse{
		ID:          m.ID,
		Title:       *m.Title,
		Description: *m.Description,
		Price:       price,
	}
}