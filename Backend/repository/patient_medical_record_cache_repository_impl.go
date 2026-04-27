package repository

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryhnfhrza/med-cache-api/helper"
	"github.com/ryhnfhrza/med-cache-api/model/domain"
)

type patientMedicalRecordCacheImpl struct {
	Client *redis.Client
}

func NewPatientMedicalRecordCache(client *redis.Client) PatientMedicalRecordCache {
	return &patientMedicalRecordCacheImpl{
		Client: client,
	}
}

func (repository *patientMedicalRecordCacheImpl) getRedisKey(id string) string {
	return "patient:" + id
}

func (repository *patientMedicalRecordCacheImpl) Set(ctx context.Context, patient domain.PatientMedicalRecord, ttl time.Duration) error {
	data, err := json.Marshal(patient)
	helper.PanicIfErr(err)

	key := repository.getRedisKey(strconv.Itoa(patient.Id))
	return repository.Client.Set(ctx, key, data, ttl).Err()
}

func (repository *patientMedicalRecordCacheImpl) Get(ctx context.Context, id int) (*domain.PatientMedicalRecord, error) {
	key := repository.getRedisKey(strconv.Itoa(id))

	val, err := repository.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	patient := new(domain.PatientMedicalRecord)
	err = json.Unmarshal([]byte(val), patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (repository *patientMedicalRecordCacheImpl) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return repository.Client.Del(ctx, "patients:all").Err()
	}
	key := repository.getRedisKey(strconv.Itoa(id))
	err := repository.Client.Del(ctx, key).Err()

	_ = repository.Client.Del(ctx, "patients:all")

	return err
}

func (repository *patientMedicalRecordCacheImpl) SetAll(ctx context.Context, patients []domain.PatientMedicalRecord, ttl time.Duration) error {
	data, err := json.Marshal(patients)
	if err != nil {
		return err
	}

	return repository.Client.Set(ctx, "patients:all", data, ttl).Err()
}

func (repository *patientMedicalRecordCacheImpl) GetAll(ctx context.Context) ([]domain.PatientMedicalRecord, error) {
	val, err := repository.Client.Get(ctx, "patients:all").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var patients []domain.PatientMedicalRecord
	err = json.Unmarshal([]byte(val), &patients)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
