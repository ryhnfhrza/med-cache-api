package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ryhnfhrza/med-cache-api/controller"
)

func NewRouter(patientMedicalRecordController controller.PatientMedicalRecordController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/medical-record", patientMedicalRecordController.CreateMedicalRecord)
	router.PUT("/api/medical-record/:id", patientMedicalRecordController.UpdateMedicalRecord)
	router.DELETE("/api/medical-record/:id", patientMedicalRecordController.DeleteMedicalRecord)
	router.GET("/api/medical-record/:id", patientMedicalRecordController.FindMedicalRecordById)
	router.GET("/api/medical-record", patientMedicalRecordController.FindAllMedicalRecord)

	return router
}
