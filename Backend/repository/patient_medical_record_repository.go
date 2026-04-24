package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/med-cache-api/model/domain"
)

type PatientMedicalRecordRepository interface {
	Save(ctx context.Context, tx *sql.Tx, patientMedicalRecord domain.PatientMedicalRecord) (domain.PatientMedicalRecord, error)
	Update(ctx context.Context, tx *sql.Tx, patientMedicalRecord domain.PatientMedicalRecord) (domain.PatientMedicalRecord, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PatientMedicalRecord, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PatientMedicalRecord, error)
}
