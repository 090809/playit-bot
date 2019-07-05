package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var emptyUser = errors.New("empty user")

// var checkTagUrl = "https://geek.bonch.dev/api/user/tag/%s?_key=2281337"
// var confirmTagURL = "https://geek.bonch.dev/api/user/tag/%s?confirm=true&_key=2281337"
type UserApi struct {
	UserData *UserData `json:"user,omitempty"`
}

type UserData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CheckTag(tag string) (string, error) {
	checkTagUrl := os.Getenv("CHECK_ID_URL")
	url := fmt.Sprintf(checkTagUrl, tag)

	log.Printf("url: %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	defer resp.Body.Close()

	var ua UserApi

	htmlData, err := ioutil.ReadAll(resp.Body)

	log.Print(string(htmlData))

	if err := json.Unmarshal(htmlData, &ua); err != nil {
		log.Printf("[ERROR] %v", err)
	}

	if ua.UserData == nil {
		return "", emptyUser
	}

	return fmt.Sprintf("%s %s", ua.UserData.FirstName, ua.UserData.LastName), nil
}

func Confirm(what string, tag string) error {
	var URL string

	switch what {
	case "tag":
		URL = fmt.Sprintf(os.Getenv("CONFIRM_TAG_URL"), tag)
	case "idea":
		URL = fmt.Sprintf(os.Getenv("CONFIRM_IDEA_URL"), tag)
	case "test":
		URL = fmt.Sprintf(os.Getenv("CONFIRM_TEST_URL"), tag)
	}

	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
	defer resp.Body.Close()

	return nil
}
