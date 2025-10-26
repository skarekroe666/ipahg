package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchApi() {

	url := "https://api.github.com/users/skarekroe666"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("could not get the response:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("could not read the response:", err)
	}

	var u User
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Fatal("could not unmarshal data:", err)
	}
	fmt.Println(u)
}
