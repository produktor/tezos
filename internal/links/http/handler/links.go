package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) Links(w http.ResponseWriter, r *http.Request) {
	var response LinksResponse
	var groupsIDsReq LinksRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &groupsIDsReq)

	tgGroupsLinks, err := h.s.GetTgLinksByGroupsIDs(r.Context(), groupsIDsReq.GroupsIDs)
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
