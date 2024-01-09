package handlers

import (
	"fmt"
	"net/http"
	"os"
)
//Args:{CHECK_ORG, CHECK_REPO}
func FileExists() map[string]interface{} {
	resp, err := http.Get("https://api.github.com/repos/" + os.Getenv("CHECK_ORG") + "/" + os.Getenv("CHECK_REPO")+"/contents")
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
		"file_exists" : resp.StatusCode,
	}
}
