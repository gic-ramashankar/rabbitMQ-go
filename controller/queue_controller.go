package controller

import (
	"demo/helper"
	"demo/model"
	"demo/service"
	"encoding/json"
	"net/http"
)

func SendToQueue(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var reqBody model.Request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	res, err := service.InsertInQueue(reqBody)

	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.RespondWithJson(w, http.StatusOK, "Record inserted successfully", helper.StatusCodeOK, res)
}
