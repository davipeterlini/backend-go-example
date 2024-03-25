# Backend Golang Example

This repo contains a example the backend in golang
This example contains the features below: 
* CRUD - Register Vehicle
    - API
    - Database Conection (Postgres)
* Env Files 
* 

# PROJECT ORGANIZATION 

/backend-go-example
    /cmd
        /api           # Application entrypoint
    /curl 
        vehicle.sh     # Shell script for testing vehicle APIs using curl
    /internal
        /vehicle       # Domain-specific logic for vehicles
        /generics
            /repository # Generic repository implementations
            /handler    # Generic HTTP handlers
            /model      # Generic data models
            /service    # Generic service logic
    /pkg    
        /config        # Configuration related to .env file
        /response      # Utilities for HTTP responses
        /httperror     # HTTP error handling utilities
    /postman
        backend_go_example.postman_collection.json # Postman collection for API testing
    /sql
        create_table.sql # SQL script for creating database tables
    /config            # Application configurations (e.g., database configurations)
    /.env              # Environment variables file
    /.gitignore        # Files and directories ignored by Git
    /README.md         # API documentation
    /go.mod            # Go module dependencies
    /go.sum            # Sum file for module dependencies



# Setup Go 
## MAC - Install with brew
* Befare setup go use this reference to setup dev environment - [scripts](https://github.com/davipeterlini/scripts)
```shell script
brew install go
go version
```

# Start with golang (only first)
```shell script
# its not necessary to use, only use in create the new project
go mod init 
```

# Adding Libs 
## API
```shell script
go get . # update all
go get github.com/gorilla/mux # router
```

## Env File
```shell script
go get . # update all
go get github.com/joho/godotenv # env files.
go get -u github.com/joho/godotenv
```

## Database
```shell script
go get github.com/lib/pq #postgres
```

# Update Libs
```shell script
go get . # update all
go get -u github.com/<dependency> # update specific dependency
```

# Dev - Local
## Start Database in manual
```shell script
brew install --cask rancher-desktop #brew install --cask docker
open -a "Rancher Desktop" # Open GUI for start
docker run hello-world # Test
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
docker ps 
```

# Run Project
```shell script
go run cmd/api/main.go
```



# Debug in Golang - DEV
## Install Delve
```shell script
go install github.com/go-delve/delve/cmd/dlv@latest
```

## Config Debug with VScode
* Open Project in Go
* Set de Breakpoint in Line
* Open Run and Debug in VSCode
* Click in "Create a launch.json --> GO --> Go Launch Package
* Install envFiles - DotENV
```shell script

```
* Open file launch.json in folder .vscode
# Put the code 
```shell script
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Main",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/cmd/api/main.go",
            "envFile": "${workspaceFolder}/.env",  // Especifica o caminho para o seu arquivo .env
            "args": []
        }
    ]
}
```