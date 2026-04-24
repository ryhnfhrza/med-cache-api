package service

import (
	"context"

	"github.com/ryhnfhrza/med-cache-api/model/web"
)

type PatientMedicalRecordService interface {
	CreateMedicalRecord(ctx context.Context, request web.PatientMedicalRecordCreateRequest) web.PatientMedicalRecordResponse
	UpdateMedicalRecord(ctx context.Context, request web.PatientMedicalRecordUpdateRequest) web.PatientMedicalRecordResponse
	DeleteMedicalRecord(ctx context.Context, id int)
	FindMedicalRecordById(ctx context.Context, id int) web.PatientMedicalRecordResponse
	FindAllMedicalRecord(ctx context.Context) []web.PatientMedicalRecordResponse
}
