package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/ryhnfhrza/med-cache-api/app"
	"github.com/ryhnfhrza/med-cache-api/controller"
	"github.com/ryhnfhrza/med-cache-api/exception"
	"github.com/ryhnfhrza/med-cache-api/helper"
	"github.com/ryhnfhrza/med-cache-api/repository"
	"github.com/ryhnfhrza/med-cache-api/service"
)

func main() {
	envPath := filepath.Join("..", ".env")

	if p := os.Getenv("CONFIG_PATH"); p != "" {
		envPath = p
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: gagal memuat %s: %v", envPath, err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	db := app.NewDB()

	validate := validator.New()

	patientMedicalRecordRepository := repository.NewpatientMedicalRecordRepository()
	patientMedicalRecordService := service.NewpatientMedicalRecordService(patientMedicalRecordRepository, db, validate)
	patientMedicalRecordController := controller.NewPatientMedicalRecordController(patientMedicalRecordService)

	router := app.NewRouter(patientMedicalRecordController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    ":" + port,
		Handler: app.CORS(router),
	}

	log.Printf("Server running on port %s", port)
	err := server.ListenAndServe()
	helper.PanicIfErr(err)

}
