package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sales/vehicle/internal/generics/repository"
	"sales/vehicle/internal/vehicle"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	env := os.Getenv("ENV")
	log.Println("Environment:", env)

	// Check if the PostgreSQL
	checkContainerCmd := exec.Command("docker", "inspect", "--format={{.State.Running}}", "postgres")
	var out bytes.Buffer
	checkContainerCmd.Stdout = &out
	err := checkContainerCmd.Run()
	isRunning := out.String() == "true\n"

	if err == nil && isRunning {
		// If the container is running, clear the database
		log.Println("PostgreSQL container is running, clearing the database...")
		clearDBCmd := exec.Command("docker", "exec", "postgres", "psql", "-U", "postgres", "-c", "DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
		if err := clearDBCmd.Run(); err != nil {
			log.Fatalf("Failed to clear PostgreSQL database: %v", err)
		}
	} else {
		// If the container is not running, stop and remove the container (if it exists)
		log.Println("PostgreSQL container is not running, stopping and removing existing container if it exists...")
		stopCmd := exec.Command("docker", "stop", "postgres")
		if err := stopCmd.Run(); err != nil {
			log.Printf("Não foi possível parar o contêiner 'postgres': %v", err)
		}
		rmCmd := exec.Command("docker", "rm", "postgres")
		if err := rmCmd.Run(); err != nil {
			log.Printf("Não foi possível remover o contêiner 'postgres': %v", err)
		}
		rmiCmd := exec.Command("docker", "rmi", "postgres")
		if err := rmiCmd.Run(); err != nil {
			log.Printf("Não foi possível remover a imagem 'postgres': %v", err)
		}

		if env == "local" && runtime.GOOS == "darwin" && runtime.GOARCH == "amd64" {
			log.Println("Attempting to start PostgreSQL container...")
			cmd := exec.Command("docker", "run", "--name", "postgres", "-e", "POSTGRES_PASSWORD="+os.Getenv("DB_PASSWORD"), "-d", "-p", "5432:5432", "postgres")
			cmd.Stdout = os.Stdout // Direct standard output of the command to the Go program's stdout
			cmd.Stderr = os.Stderr // Direct standard error of the command to the Go program's stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("Failed to start database container: %v", err)
			}
			// Optionally, add logic here to wait and ensure the database is ready before proceeding
		} else {
			log.Fatalf("Database not available: %v", err)
		}
	}

	// Initialize the repository and service layers here as per your application's requirements
	// Example (you should replace this with your actual implementation):
	baseRepo, err := repository.NewBaseRepository()
	if err != nil {
		log.Fatalf("Failed to initialize the base repository: %v", err)
	}

	// Lê a pasta 'sql' e cria uma lista de arquivos SQL
	sqlFiles, err := ioutil.ReadDir("sql")
	if err != nil {
		log.Fatalf("Failed to read sql directory: %v", err)
	}

	// Cria as tabelas se não existir de acordo com os arquivos SQL que representam as entidades
	for _, f := range sqlFiles {
		if filepath.Ext(f.Name()) == ".sql" {
			content, err := ioutil.ReadFile(filepath.Join("sql", f.Name()))
			if err != nil {
				log.Fatalf("Failed to read SQL file (%s): %v", f.Name(), err)
			}

			tableName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
			var exists bool
			query := fmt.Sprintf("SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = '%s');", tableName)
			err = baseRepo.DB().QueryRow(query).Scan(&exists)
			if err != nil {
				log.Fatalf("Error checking if table %s exists: %v", tableName, err)
			}

			if !exists {
				log.Printf("Table %s does not exist, creating...", tableName)
				_, err = baseRepo.DB().Exec(string(content))
				if err != nil {
					log.Fatalf("Failed to create table %s: %v", tableName, err)
				}
			}
		}
	}

	// Aqui você usa a base do repositório para criar um repositório específico do veículo.
	vehicleRepo := vehicle.NewVehicleRepository(baseRepo)

	// Então, você passa o repositório de veículo para a função NewVehicleService.
	service := vehicle.NewVehicleService(vehicleRepo)

	// Cria um novo handler com o serviço.
	handler := vehicle.NewHandler(service)

	// Cria um novo roteador Gorilla Mux.
	router := mux.NewRouter()

	// Registra as rotas do Handler no roteador.
	handler.RegisterVehicleRoutes(router)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Alive and kicking!"))
	})

	// Start the server with the Gorilla Mux router
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
