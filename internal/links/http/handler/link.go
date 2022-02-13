package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func(h *Handler) Link(w http.ResponseWriter, r *http.Request) {
	test := r.Form.Get("groupID")

	test1 := r.Form

	fmt.Println(test1)

	groupID, err := strconv.ParseInt(test, 10, 64)
	if err == nil {
		w.Write([]byte(err.Error()))

		return
	}

	tgGroupLink, err := h.s.GetTgLinkByGroupID(r.Context(), groupID)
	if err != nil {
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
