# GitHub Checks Microservice

The GitHub Checks Microservice is a versatile and extensible Go-based application designed to interact with the GitHub API. Its primary purpose is to provide comprehensive checks for GitHub repositories, individual files, and file contents. This microservice is configured through environment variables, offering users flexibility to adapt it to their specific requirements.


## Features

- Repository Exists:
This feature enables users to perform a check that the GitHub repository exists.

- File Exists:
Users can check the status of specific files within a repository. This functionality is useful for monitoring the added individual files over time.

- File Contains:
The microservice allows users to inspect the content of files within a GitHub repository. This feature is particularly valuable for analyzing file contents and ensuring compliance with specific content criteria.


## Configuration
The GitHub Checks Microservice is configured through environment variables, providing a straightforward and adaptable setup process. Users can customize the microservice behavior by setting the following environment variables:

<b>CHECK_CALLBACK_URL:</b> URL provided will recieve the check resp.<br>
<b>CHECK_ACTION:</b> The check in which you want to want to preform.<br>
<b>CHECK_ORG:</b> The ORG to check.<br>
<b>CHECK_REPO:</b> The Repository to check.<br>
<b>CHECK_FILE:</b> The File to check <br>
<b>CHECK_STRING:</b> String to check for in file.<br>




Examples:<br>
```
export CHECK_CALLBACK_URL=localhost:3000
export CHECK_ACTION=repo_exists
export CHECK_ORG=my_org
export CHECK_REPO=my_repo
export CHECK_FILE=my_file
export CHECK_STRING=my_string
```

