package handler

import (
	"encoding/json"
	"encounters-service/dto"
	"encounters-service/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EncounterHandler struct {
	Service *service.EncounterService
}

func (h EncounterHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.Service.Get(int64(idInt))
	if err != nil {
		fmt.Println("Error while accepting the request:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error while encoding response:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func (h EncounterHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	result, err := h.Service.GetAll()
	if err != nil {
		fmt.Println("Error while accepting the request:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error while encoding response:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func (h EncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var requestDto dto.EncounterDto
	err := json.NewDecoder(req.Body).Decode(&requestDto)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.Service.Create(requestDto)
	if err != nil {
		fmt.Println("Error while accepting the request:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error while encoding response:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func (h EncounterHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var requestDto dto.EncounterDto
	err := json.NewDecoder(req.Body).Decode(&requestDto)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.Service.Update(requestDto)
	if err != nil {
		fmt.Println("Error while accepting the request:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error while encoding response:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}

func (h EncounterHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	h.Service.Delete(int64(idInt))
	if err != nil {
		fmt.Println("Error while accepting the request:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}
