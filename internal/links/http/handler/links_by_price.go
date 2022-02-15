package handler

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
)

func(h *Handler) LinksByPrice(w http.ResponseWriter, r *http.Request) {
	var response LinksResponse

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	price, err := decimal.NewFromString(r.Form.Get("price"))
	if err != nil {
		fmt.Println(err)
	}

	tgGroupsLinks, err := h.s.GetTgLinksByPrice(r.Context(), price)
	if err != nil {
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response.FromModel(tgGroupsLinks)); err != nil {
		log.Printf("failed to encode response: %s", err)
	}
}