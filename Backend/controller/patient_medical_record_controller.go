package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PatientMedicalRecordController interface {
	CreateMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindMedicalRecordById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
