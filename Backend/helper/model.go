package helper

import (
	"github.com/ryhnfhrza/med-cache-api/model/domain"
	"github.com/ryhnfhrza/med-cache-api/model/web"
)

func ToPatientMedicalRecordResponse(patientMedicalRecord domain.PatientMedicalRecord) web.PatientMedicalRecordResponse {
	return web.PatientMedicalRecordResponse{
		Id:          patientMedicalRecord.Id,
		Name:        patientMedicalRecord.Name,
		Diagnosis:   patientMedicalRecord.Diagnosis,
		DrugTherapy: patientMedicalRecord.DrugTherapy,
	}
}
func ToPatientMedicalRecordResponses(patientMedicalRecord []domain.PatientMedicalRecord) []web.PatientMedicalRecordResponse {
	var patientMedicalRecordResponse []web.PatientMedicalRecordResponse
	for _, medicalRecord := range patientMedicalRecord {
		patientMedicalRecordResponse = append(patientMedicalRecordResponse, ToPatientMedicalRecordResponse(medicalRecord))
	}
	return patientMedicalRecordResponse
}
