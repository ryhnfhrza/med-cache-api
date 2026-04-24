package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ryhnfhrza/med-cache-api/helper"
	"github.com/ryhnfhrza/med-cache-api/model/web"
	"github.com/ryhnfhrza/med-cache-api/service"
)

type patientMedicalRecordControllerImpl struct {
	Service service.PatientMedicalRecordService
}

func NewPatientMedicalRecordController(patientMedicalRecordService service.PatientMedicalRecordService) PatientMedicalRecordController {
	return &patientMedicalRecordControllerImpl{Service: patientMedicalRecordService}
}

func (controller *patientMedicalRecordControllerImpl) CreateMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	patientMedicalRecordCreateRequest := web.PatientMedicalRecordCreateRequest{}
	helper.ReadFromRequestBody(request, &patientMedicalRecordCreateRequest)

	patientMedicalRecordResponse := controller.Service.CreateMedicalRecord(request.Context(), patientMedicalRecordCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   patientMedicalRecordResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *patientMedicalRecordControllerImpl) UpdateMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	patientMedicalRecordUpdateRequest := web.PatientMedicalRecordUpdateRequest{}
	helper.ReadFromRequestBody(request, &patientMedicalRecordUpdateRequest)

	idString := params.ByName("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfErr(err)

	patientMedicalRecordUpdateRequest.Id = id

	patientMedicalRecordResponse := controller.Service.UpdateMedicalRecord(request.Context(), patientMedicalRecordUpdateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   patientMedicalRecordResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *patientMedicalRecordControllerImpl) DeleteMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idString := params.ByName("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfErr(err)

	controller.Service.DeleteMedicalRecord(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusNoContent,
		Status: "NO CONTENT",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *patientMedicalRecordControllerImpl) FindMedicalRecordById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idString := params.ByName("id")
	id, err := strconv.Atoi(idString)
	helper.PanicIfErr(err)

	patientMedicalRecordResponse := controller.Service.FindMedicalRecordById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   patientMedicalRecordResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *patientMedicalRecordControllerImpl) FindAllMedicalRecord(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	patientMedicalRecordResponse := controller.Service.FindAllMedicalRecord(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   patientMedicalRecordResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
