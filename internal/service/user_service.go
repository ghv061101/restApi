package service

import (
	"github.com/ghv061101/RestApiAge/internal/models"
	"github.com/ghv061101/RestApiAge/internal/repository"
)

type UserService struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(u *models.Users) error           { return s.repo.CreateUser(u) }
func (s *UserService) List() ([]models.Users, error)          { return s.repo.ListUsers() }
func (s *UserService) GetByID(id uint) (*models.Users, error) { return s.repo.GetUserByID(id) }
func (s *UserService) Update(u *models.Users) error           { return s.repo.UpdateUser(u) }
func (s *UserService) Delete(id uint) (int64, error)          { return s.repo.DeleteUser(id) }
