package user

import "boilerplate/models"

type Service interface {
	Find(id int) (models.User, error)
}

//service struct
type service struct {
	repository Repository
}

// NewService is the single instance repo that is being created.
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

//ReturnFinds - Return Finds
func (s *service) Find(id int) (models.User, error) {
	return s.repository.Find(id)
}
