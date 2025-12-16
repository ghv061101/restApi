package repository

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"

	"github.com/ghv061101/RestApiAge/internal/models"
)

// Repository is the concrete repository implementation used by the service.
type Repository struct {
	DB  *gorm.DB
	SQL *sql.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func NewSQL(db *sql.DB) *Repository {
	return &Repository{SQL: db}
}

func (r *Repository) CreateUser(u *models.Users) error {
	if r.SQL != nil {
		// use SQL implementation
		var id int64
		err := r.SQL.QueryRowContext(context.Background(), "INSERT INTO users (name, dob) VALUES ($1, $2) RETURNING id", u.Name, u.Dob).Scan(&id)
		if err != nil {
			return err
		}
		u.ID = uint(id)
		return nil
	}
	return r.DB.Create(u).Error
}

func (r *Repository) ListUsers() ([]models.Users, error) {
	if r.SQL != nil {
		rows, err := r.SQL.QueryContext(context.Background(), "SELECT id, name, dob FROM users ORDER BY id")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var users []models.Users
		for rows.Next() {
			var u models.Users
			var dob time.Time
			if err := rows.Scan(&u.ID, &u.Name, &dob); err != nil {
				return nil, err
			}
			u.Dob = dob
			users = append(users, u)
		}
		return users, nil
	}
	var users []models.Users
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserByID(id uint) (*models.Users, error) {
	if r.SQL != nil {
		var u models.Users
		var dob time.Time
		err := r.SQL.QueryRowContext(context.Background(), "SELECT id, name, dob FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &dob)
		if err != nil {
			return nil, err
		}
		u.Dob = dob
		return &u, nil
	}
	var u models.Users
	if err := r.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) UpdateUser(u *models.Users) error {
	if r.SQL != nil {
		res, err := r.SQL.ExecContext(context.Background(), "UPDATE users SET name=$1, dob=$2 WHERE id=$3", u.Name, u.Dob, u.ID)
		if err != nil {
			return err
		}
		_, err = res.RowsAffected()
		return err
	}
	return r.DB.Save(u).Error
}

func (r *Repository) DeleteUser(id uint) (int64, error) {
	if r.SQL != nil {
		res, err := r.SQL.ExecContext(context.Background(), "DELETE FROM users WHERE id=$1", id)
		if err != nil {
			return 0, err
		}
		rows, err := res.RowsAffected()
		return rows, err
	}
	res := r.DB.Delete(&models.Users{}, id)
	return res.RowsAffected, res.Error
}
