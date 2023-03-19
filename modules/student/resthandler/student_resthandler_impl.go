package resthandler

import (
	"backend-ekkn/modules/student/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentResthandlerImpl struct {
	service service.StudentService
}

func NewStudentResthandler(service service.StudentService) StudentResthandler {
	return &StudentResthandlerImpl{service}
}

func (handler *StudentResthandlerImpl) CreateStudent(c *gin.Context) {
	var studentRequest shareddomain.CreateStudentRequest

	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		// format better error
		errors := helper.FormatValidationError(err)
		errorData := gin.H{"errors": errors}

		// create response
		response := helper.APIResponse(http.StatusUnprocessableEntity, false, "Mahasiswa gagal ditambahkan", errorData)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send data to service and get return
	student, err := handler.service.CreateStudent(studentRequest)
	if err != nil {
		errorData := gin.H{"error": err.Error()}

		// create response
		response := helper.APIResponse(http.StatusBadRequest, false, "Mahasiswa gagal ditambahkan", errorData)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// jwt service generation

	//map domain to respon data
	responseStudent := shareddomain.ToResponseStudent(student, "sfsffdfds")

	// create response
	response := helper.APIResponse(http.StatusCreated, true, "Mahasiswa berhasil ditambahkan", responseStudent)

	c.JSON(http.StatusCreated, response)
}

func (handler *StudentResthandlerImpl) FindStudentByNim(c *gin.Context) {
	// get params from url path
	nim := c.Param("nim")

	// send data and get return from service
	student, err := handler.service.FindStudentByNim(nim)
	if err != nil {
		errorData := gin.H{"error": err.Error()}

		// create response
		response := helper.APIResponse(http.StatusNotFound, false, "Mahasiswa tidak ditemukan", errorData)

		c.JSON(http.StatusNotFound, response)
		return
	}

	// map data domain to response data
	responseStudent := shareddomain.ToResponseFindStudentByNim(student)

	// create response
	response := helper.APIResponse(http.StatusOK, true, "Mahasiswa berhasil didapatkan", responseStudent)

	c.JSON(http.StatusOK, response)
}

func (handler *StudentResthandlerImpl) LoginStudent(c *gin.Context) {
	var studentRequest shareddomain.LoginStudentRequest

	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		errors := helper.FormatValidationError(err)
		errorData := gin.H{"errors": errors}

		// create response
		response := helper.APIResponse(http.StatusUnprocessableEntity, false, "Mahasiswa gagal login", errorData)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send data and get return from service
	student, err := handler.service.LoginStudent(studentRequest)
	if err != nil {
		errorData := gin.H{"error": err.Error()}

		// create response
		response := helper.APIResponse(http.StatusNotFound, false, "Mahasiswa gagal login, nim tidak ditemukan", errorData)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// jwt service generation

	//map domain to respon data
	responseStudent := shareddomain.ToResponseStudent(student, "sfsffdfds")

	// create response
	response := helper.APIResponse(http.StatusOK, true, "Mahasiswa berhasil Login", responseStudent)

	c.JSON(http.StatusOK, response)
}
