package products

import (
	"database/sql"
	"log"

	"github.com/pradeep/golang-micro/model"
	"github.com/pradeep/golang-micro/query"
)

type BackendRepo struct {
	db *sql.DB
}

func NewBackendRepo(db *sql.DB) *BackendRepo {
	return &BackendRepo{db: db}
}
func (s *BackendRepo) GetAllEmployees() ([]model.Users, error) {
	rows, err := s.db.Query(query.Getalluser)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Uuid, &user.FirstName, &user.LastName, &user.Password, &user.Email,
			&user.Phone, &user.Role, &user.User_id, &user.Created_at,
			&user.Updated_at)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		users = append(users, user)
	}
	transformedUsers := make([]model.Users, len(users))
	for i, user := range users {
		transformedUsers[i] = model.Users{
			User_id:    user.User_id,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			Phone:      user.Phone,
			Role:       user.Role,
			Created_at: user.Created_at,
			Updated_at: user.Updated_at,
		}
	}
	return transformedUsers, nil
}
func (s *BackendRepo) GetUserbyID(id string) (*model.Users, error) {
	rows, err := s.db.Query(query.GetUser, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user model.User
	if rows.Next() {
		err := rows.Scan(&user.Uuid, &user.FirstName, &user.LastName, &user.Password, &user.Email,
			&user.Phone, &user.Role, &user.User_id,
			&user.Created_at, &user.Updated_at)
		if err != nil {
			return nil, err
		}
	}
	return &model.Users{
		User_id:   user.User_id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
	}, nil
}
