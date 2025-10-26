package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

func loadEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("could not load .env: %w", err)
	}

	url := os.Getenv("URL")

	return url, nil
}

func FetchApi(cmd string) {
	ghUrl, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	fullURL := ghUrl + cmd

	limiter := rate.NewLimiter(rate.Every(5*time.Second), 5)

	client := http.Client{}

	if err := limiter.Wait(context.Background()); err != nil {
		log.Fatal("Error waiting for rate limiter:", err)
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatal("could not execute the request for the url:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("could not get the response:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("API not available. Status: %d, Body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("could not read the response:", err)
	}

	var u User
	if err := json.Unmarshal(body, &u); err != nil {
		log.Fatal("could not unmarshal data:", err)
	}
	// fmt.Println(u)

	fmt.Printf("User URL: %v\n", u.URL)
	fmt.Printf("Username: %v\n", u.Username)
	fmt.Printf("Public Repos: %v\n", u.Repos)
	fmt.Printf("Created At: %v\n", u.CreatedAt)
	fmt.Printf("Social Handles: %v\n", u.Social)
}
