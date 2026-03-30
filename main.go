package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chiM "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rafaeldepontes/god-bot/internal/bot"
)

func init() {
	_ = godotenv.Load(".env")
	bot.BotToken = os.Getenv("DISCORD_KEY")
}

// httpServer just a fake server for Render to allow to use their free plan
// while I'm hosting the real thing... MY BOT!!!
func httpServer() {
	r := chi.NewRouter()
	r.Use(chiM.StripSlashes)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: false,
		MaxAge:           300, //5 min
	}))

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	log.Printf("HTTP server running on %s\n", os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

func main() {
	go httpServer()
	bot.Run()
}
