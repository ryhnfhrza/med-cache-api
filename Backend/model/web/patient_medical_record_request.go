package web

type PatientMedicalRecordCreateRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Diagnosis   string `json:"diagnosis" validate:"required,max=200"`
	DrugTherapy string `json:"drug_therapy" validate:"required,max=250"`
}
type PatientMedicalRecordUpdateRequest struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"max=100"`
	Diagnosis   string `json:"diagnosis" validate:"max=200"`
	DrugTherapy string `json:"drug_therapy" validate:"max=250"`
}
