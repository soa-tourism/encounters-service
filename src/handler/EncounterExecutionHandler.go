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

type EncounterExecutionHandler struct {
	ExecutionService *service.EncounterExecutionService
	EncounterService *service.EncounterService
}

func (h EncounterExecutionHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.ExecutionService.Get(int64(idInt))
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

func (h EncounterExecutionHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var requestDto dto.EncounterExecutionDto
	err := json.NewDecoder(req.Body).Decode(&requestDto)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := h.ExecutionService.Update(requestDto)
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

func (h EncounterExecutionHandler) Activate(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	touristId := mux.Vars(req)["touristId"]
	touristLatitude := mux.Vars(req)["touristLatitude"]
	touristLongitude := mux.Vars(req)["touristLongitude"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristIdInt, err3 := strconv.Atoi(touristId)
	if err3 != nil {
		fmt.Println("Error converting id to int:", err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLatitudeFloat, err1 := strconv.ParseFloat(touristLatitude, 64)
	if err1 != nil {
		fmt.Println("Error converting id to int:", err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLongitudeFloat, err2 := strconv.ParseFloat(touristLongitude, 64)
	if err2 != nil {
		fmt.Println("Error converting id to int:", err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err5 := h.ExecutionService.Activate(int64(touristIdInt), touristLatitudeFloat, touristLongitudeFloat, int64(idInt))
	if err5 != nil {
		fmt.Println("Error while accepting the request:", err5)
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

func (h EncounterExecutionHandler) CompleteExecution(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	touristId := mux.Vars(req)["touristId"]
	touristLatitude := mux.Vars(req)["touristLatitude"]
	touristLongitude := mux.Vars(req)["touristLongitude"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristIdInt, err3 := strconv.Atoi(touristId)
	if err3 != nil {
		fmt.Println("Error converting id to int:", err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLatitudeFloat, err1 := strconv.ParseFloat(touristLatitude, 64)
	if err1 != nil {
		fmt.Println("Error converting id to int:", err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLongitudeFloat, err2 := strconv.ParseFloat(touristLongitude, 64)
	if err2 != nil {
		fmt.Println("Error converting id to int:", err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err5 := h.ExecutionService.CompleteExecution(int64(idInt), int64(touristIdInt), touristLatitudeFloat, touristLongitudeFloat)
	if err5 != nil {
		fmt.Println("Error while accepting the request:", err5)
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

func (h EncounterExecutionHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	touristId := mux.Vars(req)["touristId"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristIdInt, err3 := strconv.Atoi(touristId)
	if err3 != nil {
		fmt.Println("Error converting id to int:", err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	h.ExecutionService.Delete(int64(idInt), int64(touristIdInt))
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (h EncounterExecutionHandler) GetAllByTourist(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result := h.ExecutionService.GetAllByTourist(int64(idInt))
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

func (h EncounterExecutionHandler) GetAllCompletedByTourist(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result := h.ExecutionService.GetAllCompletedByTourist(int64(idInt))
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

func (h EncounterExecutionHandler) GetByTour(writer http.ResponseWriter, req *http.Request) {
	var encounterIds []int64
	err := json.NewDecoder(req.Body).Decode(&encounterIds)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLatitude := mux.Vars(req)["touristLatitude"]
	touristLongitude := mux.Vars(req)["touristLongitude"]
	touristId := mux.Vars(req)["touristId"]

	touristIdInt, err3 := strconv.Atoi(touristId)
	if err3 != nil {
		fmt.Println("Error converting id to int:", err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	touristLatitudeFloat, err1 := strconv.ParseFloat(touristLatitude, 64)
	if err1 != nil {
		fmt.Println("Error converting id to int:", err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLongitudeFloat, err2 := strconv.ParseFloat(touristLongitude, 64)
	if err2 != nil {
		fmt.Println("Error converting id to int:", err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result := h.ExecutionService.GetVisibleByTour(encounterIds, touristLongitudeFloat, touristLatitudeFloat, int64(touristIdInt))
	result = h.EncounterService.AddEncounterToExecution(result)
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

func (h EncounterExecutionHandler) CheckPosition(writer http.ResponseWriter, req *http.Request) {
	var encounterIds []int64
	err := json.NewDecoder(req.Body).Decode(&encounterIds)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id := mux.Vars(req)["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLatitude := mux.Vars(req)["touristLatitude"]
	touristLongitude := mux.Vars(req)["touristLongitude"]
	touristId := mux.Vars(req)["touristId"]

	touristIdInt, err3 := strconv.Atoi(touristId)
	if err3 != nil {
		fmt.Println("Error converting id to int:", err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	touristLatitudeFloat, err1 := strconv.ParseFloat(touristLatitude, 64)
	if err1 != nil {
		fmt.Println("Error converting id to int:", err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLongitudeFloat, err2 := strconv.ParseFloat(touristLongitude, 64)
	if err2 != nil {
		fmt.Println("Error converting id to int:", err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result := h.ExecutionService.GetWithUpdatedLocation(encounterIds, int64(idInt), touristLongitudeFloat, touristLatitudeFloat, int64(touristIdInt))
	result = h.EncounterService.AddEncounterToExecution(result)
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

func (h EncounterExecutionHandler) CheckPositionLocationEncounter(writer http.ResponseWriter, req *http.Request) {
	var encounterIds []int64
	err := json.NewDecoder(req.Body).Decode(&encounterIds)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id := mux.Vars(req)["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLatitude := mux.Vars(req)["touristLatitude"]
	touristLongitude := mux.Vars(req)["touristLongitude"]
	touristId := mux.Vars(req)["touristId"]

	touristIdInt, err3 := strconv.Atoi(touristId)
	if err3 != nil {
		fmt.Println("Error converting id to int:", err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	touristLatitudeFloat, err1 := strconv.ParseFloat(touristLatitude, 64)
	if err1 != nil {
		fmt.Println("Error converting id to int:", err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristLongitudeFloat, err2 := strconv.ParseFloat(touristLongitude, 64)
	if err2 != nil {
		fmt.Println("Error converting id to int:", err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result := h.ExecutionService.GetHiddenLocationEncounterWithUpdatedLocation(encounterIds, int64(idInt), touristLongitudeFloat, touristLatitudeFloat, int64(touristIdInt))
	result = h.EncounterService.AddEncounterToExecution(result)
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

func (h EncounterExecutionHandler) GetActiveByTour(writer http.ResponseWriter, req *http.Request) {
	var encounterIds []int64
	err := json.NewDecoder(req.Body).Decode(&encounterIds)
	if err != nil {
		fmt.Println("Error while parsing json:", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	touristId := mux.Vars(req)["touristId"]

	touristIdInt, err3 := strconv.Atoi(touristId)
	if err3 != nil {
		fmt.Println("Error converting id to int:", err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result := h.ExecutionService.GetActiveByTour(int64(touristIdInt), encounterIds)
	result = h.EncounterService.AddEncountersToExecution(result)
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
