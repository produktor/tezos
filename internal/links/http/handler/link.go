package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) Link(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	groupID, err := strconv.ParseInt(r.Form.Get("groupID"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	tgGroupLink, err := h.s.GetTgLinkByGroupID(r.Context(), groupID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(tgGroupLink); err != nil {
		log.Printf("failed to encode response: %s", err)
	}
}
