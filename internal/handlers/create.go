package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andremuro/test-api-go/internal/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Printf("Erro ao fazer decode do json. %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(post)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	}

	if id >= 1 {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Registro inserido com sucesso!"),
		}
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(resp)

}
