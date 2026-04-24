package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ryhnfhrza/med-cache-api/model/domain"
)

type patientMedicalRecordImpl struct{}

func NewpatientMedicalRecordRepository() PatientMedicalRecordRepository {
	return &patientMedicalRecordImpl{}
}

func (repository *patientMedicalRecordImpl) Save(ctx context.Context, tx *sql.Tx, patientMedicalRecord domain.PatientMedicalRecord) (domain.PatientMedicalRecord, error) {
	query := "insert into patient_medical_records(name,diagnosis,drug_therapy) values (?,?,?)"
	result, err := tx.ExecContext(ctx, query, patientMedicalRecord.Name, patientMedicalRecord.Diagnosis, patientMedicalRecord.DrugTherapy)
	if err != nil {
		return patientMedicalRecord, fmt.Errorf("failed to insert patient_medical_records: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return patientMedicalRecord, fmt.Errorf("failed to get last insert id : %w", err)
	}

	patientMedicalRecord.Id = int(id)

	return patientMedicalRecord, nil
}

func (repository *patientMedicalRecordImpl) Update(ctx context.Context, tx *sql.Tx, patientMedicalRecord domain.PatientMedicalRecord) (domain.PatientMedicalRecord, error) {
	query := "update patient_medical_records set name = ? , diagnosis = ?, drug_therapy = ? where id = ?"
	result, err := tx.ExecContext(ctx, query, patientMedicalRecord.Name, patientMedicalRecord.Diagnosis, patientMedicalRecord.DrugTherapy, patientMedicalRecord.Id)
	if err != nil {
		return patientMedicalRecord, fmt.Errorf("failed to update patient_medical_records (id=%d): %w", patientMedicalRecord.Id, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return patientMedicalRecord, fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rows == 0 {
		return patientMedicalRecord, fmt.Errorf("no patient_medical_records updated (id=%d)", patientMedicalRecord.Id)
	}

	return patientMedicalRecord, nil
}

func (repository *patientMedicalRecordImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "delete from patient_medical_records where id = ?"
	_, err := tx.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("failed to delete patient_medical_records (id=%d): %w", id, err)
	}

	return nil
}

func (repository *patientMedicalRecordImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.PatientMedicalRecord, error) {
	query := "select id,name,diagnosis,drug_therapy from patient_medical_records where id = ?"
	row := tx.QueryRowContext(ctx, query, id)

	patientMedicalRecord := domain.PatientMedicalRecord{}
	err := row.Scan(
		&patientMedicalRecord.Id,
		&patientMedicalRecord.Name,
		&patientMedicalRecord.Diagnosis,
		&patientMedicalRecord.DrugTherapy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return patientMedicalRecord, fmt.Errorf("patient_medical_records not found (id=%d): %w", id, err)
		}
		return patientMedicalRecord, fmt.Errorf("failed to scan patient_medical_records (id=%d): %w", id, err)
	}
	return patientMedicalRecord, nil
}

func (repository *patientMedicalRecordImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PatientMedicalRecord, error) {
	query := "select id,name,diagnosis,drug_therapy from patient_medical_records"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query patient_medical_records: %w", err)
	}

	defer rows.Close()

	var patientMedicalRecords []domain.PatientMedicalRecord

	for rows.Next() {
		var patientMedicalRecord domain.PatientMedicalRecord
		err := rows.Scan(
			&patientMedicalRecord.Id,
			&patientMedicalRecord.Name,
			&patientMedicalRecord.Diagnosis,
			&patientMedicalRecord.DrugTherapy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan patient_medical_records row : %w", err)
		}

		patientMedicalRecords = append(patientMedicalRecords, patientMedicalRecord)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error : %w", err)
	}

	return patientMedicalRecords, nil
}
