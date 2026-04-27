package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ryhnfhrza/med-cache-api/exception"
	"github.com/ryhnfhrza/med-cache-api/helper"
	"github.com/ryhnfhrza/med-cache-api/model/domain"
	"github.com/ryhnfhrza/med-cache-api/model/web"
	"github.com/ryhnfhrza/med-cache-api/repository"
)

type PatientMedicalRecordServiceImpl struct {
	Repository repository.PatientMedicalRecordRepository
	Cache      repository.PatientMedicalRecordCache
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewpatientMedicalRecordService(repository repository.PatientMedicalRecordRepository, cache repository.PatientMedicalRecordCache, db *sql.DB, validate *validator.Validate) PatientMedicalRecordService {
	return &PatientMedicalRecordServiceImpl{
		Repository: repository,
		Cache:      cache,
		DB:         db,
		Validate:   validate,
	}
}

func (service *PatientMedicalRecordServiceImpl) CreateMedicalRecord(ctx context.Context, request web.PatientMedicalRecordCreateRequest) web.PatientMedicalRecordResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	helper.PanicIfErr(err)

	patientMedicalRecord := domain.PatientMedicalRecord{
		Name:        request.Name,
		Diagnosis:   request.Diagnosis,
		DrugTherapy: request.DrugTherapy,
	}

	patientMedicalRecord, err = service.Repository.Save(ctx, tx, patientMedicalRecord)
	helper.PanicIfErr(err)

	_ = service.Cache.Delete(ctx, 0)

	return helper.ToPatientMedicalRecordResponse(patientMedicalRecord)
}

func (service *PatientMedicalRecordServiceImpl) UpdateMedicalRecord(ctx context.Context, request web.PatientMedicalRecordUpdateRequest) web.PatientMedicalRecordResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	helper.PanicIfErr(err)

	patientMedicalRecord, err := service.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(fmt.Sprintf("PatientMedicalRecordService with id:%d not found", request.Id)))
	}

	if request.Name != "" {
		patientMedicalRecord.Name = request.Name
	}

	if request.Diagnosis != "" {
		patientMedicalRecord.Diagnosis = request.Diagnosis
	}

	if request.DrugTherapy != "" {
		patientMedicalRecord.DrugTherapy = request.DrugTherapy
	}

	patientMedicalRecord, err = service.Repository.Update(ctx, tx, patientMedicalRecord)
	helper.PanicIfErr(err)

	_ = service.Cache.Delete(ctx, patientMedicalRecord.Id)

	return helper.ToPatientMedicalRecordResponse(patientMedicalRecord)
}

func (service *PatientMedicalRecordServiceImpl) DeleteMedicalRecord(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.Repository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(fmt.Sprintf("PatientMedicalRecordService with id:%d not found", id)))
	}

	err = service.Repository.Delete(ctx, tx, id)
	helper.PanicIfErr(err)

	_ = service.Cache.Delete(ctx, id)
}

func (service *PatientMedicalRecordServiceImpl) FindMedicalRecordById(ctx context.Context, id int) web.PatientMedicalRecordResponse {
	cachedPatient, err := service.Cache.Get(ctx, id)
	if err == nil && cachedPatient != nil {
		return helper.ToPatientMedicalRecordResponse(*cachedPatient)
	}

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	patientMedicalRecord, err := service.Repository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(fmt.Sprintf("Patient Medical Record with id:%d not found", id)))
	}

	_ = service.Cache.Set(ctx, patientMedicalRecord, 15*time.Minute)

	return helper.ToPatientMedicalRecordResponse(patientMedicalRecord)
}

func (service *PatientMedicalRecordServiceImpl) FindAllMedicalRecord(ctx context.Context) []web.PatientMedicalRecordResponse {
	cachedList, err := service.Cache.GetAll(ctx)
	if err == nil && cachedList != nil {
		return helper.ToPatientMedicalRecordResponses(cachedList)
	}

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	patientMedicalRecords, err := service.Repository.FindAll(ctx, tx)
	helper.PanicIfErr(err)

	_ = service.Cache.SetAll(ctx, patientMedicalRecords, 15*time.Minute)

	return helper.ToPatientMedicalRecordResponses(patientMedicalRecords)
}
