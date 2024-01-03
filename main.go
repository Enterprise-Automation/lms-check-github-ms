package main

import (
	handlers "github.com/Enterprise-Automation/check-github-ms/handlers"
	checks "github.com/Enterprise-Automation/check-ms-helper"
)

func main() {

	checks.NewCheck("repo_exists", handlers.RepoExists)
	checks.NewCheck("file_contains", handlers.FileContains)
	checks.NewCheck("file_exists", handlers.FileExists)

	checks.RegisterChecks()
}
