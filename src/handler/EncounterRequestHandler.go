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

type EncounterRequestHandler struct {
	EncounterRequestService *service.EncounterRequestService
}

func NewEncounterRequestHandler(encounterRequestService *service.EncounterRequestService) *EncounterRequestHandler {
	return &EncounterRequestHandler{
		EncounterRequestService: encounterRequestService,
	}
}

func (handler *EncounterRequestHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var requestDto dto.EncounterRequestDto
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&requestDto)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := handler.EncounterRequestService.Create(requestDto)
	if err != nil {
		fmt.Println("Error while creating a new request:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		writer.Header().Set("Content-Type", "application/json")
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error while encoding response:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write(response)
}

func (handler *EncounterRequestHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	encounterRequests, err := handler.EncounterRequestService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(encounterRequests)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterRequestHandler) AcceptRequest(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := handler.EncounterRequestService.AcceptRequest(int64(idInt))
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

func (handler *EncounterRequestHandler) RejectRequest(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := handler.EncounterRequestService.RejectRequest(int64(idInt))
	if err != nil {
		fmt.Println("Error while rejecting the request:", err)
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
