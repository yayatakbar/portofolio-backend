package main

import (
	"fmt"
	"log"
	"net/http"
	"portfolio-backend/admin"
	"portfolio-backend/database"
	"portfolio-backend/public"
)

func main() {
	client := database.ConnectDB()

	// Route untuk public (GET data diri)
	http.HandleFunc("/personal-info", enableCORS(public.PersonalInfoHandler(client)))
	http.HandleFunc("/sidebar", enableCORS(public.SidebarHandler(client)))

	// Route untuk admin (POST data diri)
	http.HandleFunc("/admin/add-personal-info", enableCORS(admin.AddPersonalInfoHandler(client)))

	http.HandleFunc("/admin/add-sidebar", enableCORS(admin.AddSidebarHandler(client)))

	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h(w, r)
	}
}
