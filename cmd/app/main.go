package main

import (
	"2025_2_PochtiVPraime/internal/handlers"
	"net/http"
)

func main() {

	h := handlers.NewHandler()

	http.HandleFunc("/register", h.Register)
	http.HandleFunc("/login", h.Login)
	http.HandleFunc("/get-boards", h.GetBoards)

	// TODO: Настроить CORS только для /register, /login, /get-boards — а не для всех
	// TODO: НУ и что нибудь ещё точно, я не прям супер знаю за cors
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	})

	println(" Сервер запущен на :8080")
	http.ListenAndServe(":8080", nil)
}
