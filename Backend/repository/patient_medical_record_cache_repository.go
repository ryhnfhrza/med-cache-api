package repository

import (
	"context"
	"time"

	"github.com/ryhnfhrza/med-cache-api/model/domain"
)

type PatientMedicalRecordCache interface {
	Set(ctx context.Context, patient domain.PatientMedicalRecord, ttl time.Duration) error
	Get(ctx context.Context, id int) (*domain.PatientMedicalRecord, error)
	Delete(ctx context.Context, id int) error
	SetAll(ctx context.Context, patients []domain.PatientMedicalRecord, ttl time.Duration) error
	GetAll(ctx context.Context) ([]domain.PatientMedicalRecord, error)
}
