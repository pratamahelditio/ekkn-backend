package shareddomain

import "backend-ekkn/modules/student/domain"

type CreateStudentRequest struct {
	Nim      string `json:"nim" binding:"required,max=13"`
	Name     string `json:"name" binding:"required,max=100"`
	Prodi    string `json:"prodi"`
	Fakultas string `json:"fakultas"`
}

type LoginStudentRequest struct {
	Nim      string `json:"nim" binding:"required,max=13"`
	Password string `json:"password" binding:"required"`
}

type StudentResponse struct {
	Nim      string `json:"nim"`
	Name     string `json:"name"`
	Prodi    string `json:"prodi"`
	Fakultas string `json:"fakultas"`
	Token    string `json:"jwt_manager"`
}

type FindStudentByNimResponse struct {
	Nim          string `json:"nim"`
	Name         string `json:"name"`
	Prodi        string `json:"prodi"`
	Fakultas     string `json:"fakultas"`
	Gender       string `json:"gender"`
	Position     string `json:"position"`
	BahasaMadura bool   `json:"bahasa_madura"`
	GroupKkn     string `json:"groupKkn"`
	Grade        string `json:"grade"`
	CreatedAt    int64  `json:"created_at"`
	UpdateAt     int64  `json:"update_at"`
}

func ToResponseStudent(student domain.Student, token string) StudentResponse {
	studentResponse := StudentResponse{
		Nim:      student.Nim,
		Name:     student.Name,
		Prodi:    student.Prodi,
		Fakultas: student.Fakultas,
		Token:    token,
	}

	return studentResponse

}

func ToResponseFindStudentByNim(student domain.Student) FindStudentByNimResponse {
	studentResponse := FindStudentByNimResponse{
		Nim:          student.Nim,
		Name:         student.Name,
		Prodi:        student.Prodi,
		Fakultas:     student.Fakultas,
		Gender:       student.Gender,
		Position:     student.Position,
		BahasaMadura: student.BahasaMadura,
		GroupKkn:     student.GroupKkn,
		Grade:        student.Grade,
		CreatedAt:    student.CreatedAt,
		UpdateAt:     student.UpdateAt,
	}

	return studentResponse
}
