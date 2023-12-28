package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func RepoExists() map[string]interface{} {
	resp, err := http.Get("https://api.github.com/repos/" + os.Getenv("CHECK_ORG") + "/" + os.Getenv("CHECK_REPO"))
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{
			"result": false,
		}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return map[string]interface{}{
			"result": false,
		}
	}
	return map[string]interface{}{
		"result": true,
	}
}
