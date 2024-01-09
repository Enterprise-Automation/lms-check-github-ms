package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/Enterprise-Automation/check-github-ms/models"
)
//Args:{CHECK_ORG, CHECK_REPO, CHECK_FILE, CHECK_STRING}
func FileContains() map[string]interface{} {
	var FileContains models.FileContains

	resp, err := http.Get("https://api.github.com/repos/" + os.Getenv("CHECK_ORG") + "/" + os.Getenv("CHECK_REPO")+"/contents/" + os.Getenv("CHECK_FILE"))
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{
			"result": false,
		}
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
			
	
	if err := json.Unmarshal(body, &FileContains); err != nil {   
    	fmt.Println("Can not unmarshal JSON")
	}
	
	decoded, err := base64.StdEncoding.DecodeString(FileContains.Content)
    if err != nil {
        fmt.Println("Error decoding string:", err)
	} 

	ans := strings.Contains(string(decoded), os.Getenv("CHECK_STRING") )
	fmt.Println(ans)
	if (ans == true){
		fmt.Println("Found '" + os.Getenv("CHECK_STRING") + "' in " + os.Getenv("CHECK_FILE") )
	} else {
		fmt.Println("Could Not Find '" + os.Getenv("CHECK_STRING") + "' in " + os.Getenv("CHECK_FILE"))
	}
	if resp.StatusCode == 404 {
		return map[string]interface{}{
			"result": false,
		}
	}
	
	return map[string]interface{}{
		"result": true,
		"file_contains" : true,
	}
}
