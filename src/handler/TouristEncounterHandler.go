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

type TouristEncounterHandler struct {
	Service        *service.EncounterService
	RequestService *service.EncounterRequestService
}

func (h TouristEncounterHandler) GetById(writer http.ResponseWriter, req *http.Request) {
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

func (h TouristEncounterHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
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

// func (h TouristEncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
// 	var requestDto dto.EncounterDto
// 	err := json.NewDecoder(req.Body).Decode(&requestDto)
// 	if err != nil {
// 		fmt.Println("Error while parsing json:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	vars := mux.Vars(req)
// 	checkpointID, err := strconv.ParseInt(vars["checkpointId"], 10, 64)
// 	if err != nil {
// 		fmt.Println("Error while parsing checkpointId:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	isSecretPrerequisite, err := strconv.ParseBool(vars["isSecretPrerequisite"])
// 	if err != nil {
// 		fmt.Println("Error while parsing checkpointId:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	result, err := h.Service.Create(requestDto)
// 	if err != nil {
// 		fmt.Println("Error while accepting the request:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	encountersClient := servers.NewToursServer("http://host.docker.internal:8081")
// 	err = encountersClient.UpdateCheckpointEncounter(strconv.FormatInt(checkpointID, 10), strconv.FormatInt(requestDto.Id, 10), isSecretPrerequisite)
// 	if err != nil {
// 		fmt.Println("Error while sending request to tours-microservice:", err)
// 		// Handle error appropriately
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	h.RequestService.Create(dto.CreateEncounterRequestDto(model.EncounterRequest{
// 		TouristId:   result.AuthorId,
// 		EncounterId: result.Id,
// 		Status:      0,
// 	}))

// 	response, err := json.Marshal(result)
// 	if err != nil {
// 		fmt.Println("Error while encoding response:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	writer.WriteHeader(http.StatusOK)
// 	writer.Header().Set("Content-Type", "application/json")
// 	writer.Write(response)
// }

func (h TouristEncounterHandler) Update(writer http.ResponseWriter, req *http.Request) {
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

func (h TouristEncounterHandler) Delete(writer http.ResponseWriter, req *http.Request) {
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
