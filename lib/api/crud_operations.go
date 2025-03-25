package api

import (
	"encoding/json"
	"net/http"

	"github.com/joikoi1212/Go_CRUD/core"
)

type Message struct {
	Message   string `json:"message"`
	CreatedBy string `json:"created_by"`
	ID        int    `json:"id"`
}

func crud_operations(w http.ResponseWriter, r *http.Request) {
	DB := core.InitDB()
	var message Message
	switch r.Method {

	case "GET":
		getMessages(w, r)

	case "POST":
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		DB.Exec("Insert into messages (message, created_by) values (?, ?)", message.Message, message.CreatedBy)
		w.WriteHeader(http.StatusCreated)

	case "PUT":
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		DB.Exec("Update messages set message = ? where id = ?", message.Message, message.ID)
		w.WriteHeader(http.StatusOK)

	case "DELETE":
		id := r.URL.Query().Get("id")
		DB.Exec("Delete from messages where id = ?", id)
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	DB := core.InitDB()
	rows, err := DB.Query("SELECT id, message FROM messages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.Message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
