package handler

import (
	"database-example/dtos"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CheckpointHandler struct {
	CheckpointService service.ICheckpointService
}

func (handler *CheckpointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Checkpoint sa id-em %s", id)
	checkpoint, err := handler.CheckpointService.FindCheckpoint(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(checkpoint)

}

func (handler *CheckpointHandler) GetByTourId(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["tourId"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle parsing error
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	log.Printf("Checkpoints with tourId %d", id)
	checkpoints, err := handler.CheckpointService.FindByTourId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if len(checkpoints) == 0 {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("[]")) // Write empty array
		return
	}

	// Encode tours into JSON
	jsonData, err := json.Marshal(checkpoints)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Write the response
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}

func (handler *CheckpointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var checkpointDto dtos.CheckpointDto
	err := json.NewDecoder(req.Body).Decode(&checkpointDto)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdCheckpointDto, err := handler.CheckpointService.Create(&checkpointDto)
	if err != nil {
		println("Error while creating a new checkpoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	// Marshal the createdCheckpointDto object into JSON
	checkpointJSON, err := json.Marshal(createdCheckpointDto)
	if err != nil {
		println("Error while encoding checkpointDto to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set response headers and write the JSON response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write(checkpointJSON)
	if err != nil {
		println("Error while writing JSON response")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
