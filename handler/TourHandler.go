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

type TourHandler struct {
	TourService service.ITourService
}

func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle parsing error
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	log.Printf("Tura sa id-em %d", id)
	tour, err := handler.TourService.FindTour(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tour)

}

func (handler *TourHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	idStr := mux.Vars(req)["authorId"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Handle parsing error
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	log.Printf("Tura sa authorId-em %d", id)
	tours, err := handler.TourService.FindByAuthorId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if len(tours) == 0 {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("[]")) // Write empty array
		return
	}

	// Encode tours into JSON
	jsonData, err := json.Marshal(tours)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Write the response
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourDto dtos.TourDto
	err := json.NewDecoder(req.Body).Decode(&tourDto)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdTourDto, err := handler.TourService.Create(&tourDto)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	// Marshal the createdTourDto object into JSON
	tourJSON, err := json.Marshal(createdTourDto)
	if err != nil {
		println("Error while encoding tourDto to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set response headers and write the JSON response
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write(tourJSON)
	if err != nil {
		println("Error while writing JSON response")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourHandler) PublishTour(writer http.ResponseWriter, req *http.Request) {
	var tourId int64
	err := json.NewDecoder(req.Body).Decode(&tourId)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.PublishTour(tourId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		// Set response headers and write the JSON response
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *TourHandler) ArchiveTour(writer http.ResponseWriter, req *http.Request) {
	var tourId int64
	err := json.NewDecoder(req.Body).Decode(&tourId)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourService.ArchiveTour(tourId)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		// Set response headers and write the JSON response
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.WriteHeader(http.StatusOK)

}
