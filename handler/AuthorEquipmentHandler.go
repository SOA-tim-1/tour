package handler

import (
	"database-example/service"
	"encoding/json"
	"net/http"
)

type AuthorEquipmentHandler struct {
	AuthorEquipmentService service.IAuthorEquipmentService
}

func (handler *AuthorEquipmentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	equipments, err := handler.AuthorEquipmentService.FindAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if len(equipments) == 0 {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("[]")) // Write empty array
		return
	}

	// Encode tours into JSON
	jsonData, err := json.Marshal(equipments)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Write the response
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}
