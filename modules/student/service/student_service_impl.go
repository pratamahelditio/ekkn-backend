package service

import (
	"backend-ekkn/modules/student/domain"
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type StudentServiceImpl struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &StudentServiceImpl{repo}
}

func (service *StudentServiceImpl) CreateStudent(request shareddomain.CreateStudentRequest) (domain.Student, error) {

	createdAt := time.Now().Unix()
	updateAt := createdAt
	student := domain.Student{
		Nim:       request.Nim,
		Name:      request.Name,
		Prodi:     request.Prodi,
		Fakultas:  request.Fakultas,
		CreatedAt: createdAt,
		UpdateAt:  updateAt,
		Position:  "none",
	}

	// hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Nim), bcrypt.MinCost)
	if err != nil {
		return student, err
	}

	student.Password = string(passwordHash)
	student, err = service.repo.Create(student)
	if err != nil {
		return student, err
	}

	return student, nil
}

func (service *StudentServiceImpl) FindStudentByNim(nim string) (domain.Student, error) {
	student, err := service.repo.FindByNim(nim)
	if err != nil {
		return student, err
	}

	if student.Nim == "" {
		return student, errors.New("No student found on that nim")
	}

	return student, nil
}

func (service *StudentServiceImpl) LoginStudent(request shareddomain.LoginStudentRequest) (domain.Student, error) {
	// check nim request is exist
	student, err := service.repo.FindByNim(request.Nim)

	if err != nil {
		return student, err
	}

	if student.Nim == "" {
		return student, errors.New("No student found on that nim")
	}

	// check password is match
	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(request.Password))
	if err != nil {
		return student, err
	}

	return student, nil
}

// get all student use repository
func (service *StudentServiceImpl) FindAllStudent() ([]domain.Student, error) {
	students, err := service.repo.FindAll()
	if err != nil {
		return students, err
	}

	return students, nil
}

func (service *StudentServiceImpl) UpdateStudent(request shareddomain.UpdateStudentRequest) (domain.Student, error) {
	// cek nim isExist
	student, err := service.repo.FindByNim(request.Nim)
	if err != nil {
		return student, err
	}

	if student.Nim == "" {
		return student, errors.New("No student found on that nim")
	}

	student = domain.Student{
		Nim:        request.Nim,
		Name:       request.Name,
		Prodi:      request.Prodi,
		Fakultas:   request.Fakultas,
		Gender:     request.Gender,
		MaduraLang: request.MaduraLang,
		UpdateAt:   time.Now().Unix(),
	}
	student, err = service.repo.Update(student)
	if err != nil {
		return student, err
	}

	return student, nil
}