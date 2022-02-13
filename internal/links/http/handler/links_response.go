package handler

import "link_api/domain/model"

type LinksResponse struct {
	TelegramLinks []TelegramLinkResponse `json:"telegramLinks"`
}

type TelegramLinkResponse struct {
	TelegramGroup TelegramGroupResponse `json:"telegramGroup"`
	Link          string                `json:"telegramLink"`
}

func (LinksResponse) FromModel(tgLinks []model.TelegramLinks) LinksResponse {
	var response LinksResponse

	for _, tgLink := range tgLinks {
		price, _ := tgLink.TgGroup.Price.Float64()

		response.TelegramLinks = append(response.TelegramLinks, TelegramLinkResponse{
			TelegramGroup: TelegramGroupResponse{
				ID:          tgLink.TgGroup.ID,
				Title:       *tgLink.TgGroup.Title,
				Description: *tgLink.TgGroup.Description,
				Price:       price,
			},
			Link:          tgLink.Link,
		})
	}


	return response
}