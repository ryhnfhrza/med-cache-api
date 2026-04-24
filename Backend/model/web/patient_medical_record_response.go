package web

type PatientMedicalRecordResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Diagnosis   string `json:"diagnosis"`
	DrugTherapy string `json:"drug_therapy"`
}
