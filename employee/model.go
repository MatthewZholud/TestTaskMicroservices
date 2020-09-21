package employee

type Employee struct {
	ID         int64  `json:"id"`
	Name       string `json:"name";validate:"name,required"`
	SecondName string `json:"second_name"`
	Surname    string `json:"surname"`
	PhotoUrl   string `json:"photo_url";validate:"photo_url,required"`
	HireDate   string `json:"hire_date"`
	Position   string `json:"position"`
	CompanyID  int64  `json:"company_id"`
}

//func NewEmployee(id, companyId int64, name, secondName, surname, photoUrl, hireDate, position string) Employee {
//	return Employee{
//		ID:         id,
//		Name:       name,
//		SecondName: secondName,
//		Surname:    surname,
//		PhotoUrl:   photoUrl,
//		HireDate:   hireDate,
//		Position:   position,
//		CompanyID:  companyId,
//	}
//}
