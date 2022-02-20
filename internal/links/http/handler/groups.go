package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GroupsList(w http.ResponseWriter, r *http.Request) {
	var response GroupsResponse

	groups, err := h.s.GetGroups(r.Context())
	if err != nil {
		w.Write([]byte(err.Error()))

		return
	}

	if len(groups) > 0 {
		for _, group := range groups {
			tgGroupResponse := TelegramGroupResponse{}

			response.TelegramGroups = append(response.TelegramGroups, tgGroupResponse.FromModel(group))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Printf("failed to encode response: %s", err)
	}
}
