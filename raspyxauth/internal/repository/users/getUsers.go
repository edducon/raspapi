package users

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/repository/constRepository"
	"context"
	"fmt"
)

func (r *UsersRepository) GetUsers() (*[]models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s.%s",
		constRepository.AUTH_SCHEMA,
		constRepository.USERS_TABLE,
	)

	ctx := context.Background()
	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var Users []models.User

	for rows.Next() {
		var User models.User

		errScan := rows.Scan(
			&User.UUID,
			&User.Username,
			&User.PasswordHash,
			&User.AccessLevel,
		)

		if errScan != nil {
			return nil, errScan
		}

		Users = append(Users, User)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &Users, nil
}

func (r *UsersRepository) GetUserByUUID(userUUID string) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE uuid=$1",
		constRepository.AUTH_SCHEMA,
		constRepository.USERS_TABLE,
	)

	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, query, userUUID)

	var User models.User

	errScan := row.Scan(
		&User.UUID,
		&User.Username,
		&User.PasswordHash,
		&User.AccessLevel,
	)

	if errScan != nil {
		return nil, errScan
	}

	return &User, nil
}

func (r *UsersRepository) GetUsersByUsername(username string) (*[]models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE LOWER(username) LIKE LOWER($1)",
		constRepository.AUTH_SCHEMA,
		constRepository.USERS_TABLE,
	)

	ctx := context.Background()
	rows, err := r.Pool.Query(ctx, query, "%"+username+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var Users []models.User

	for rows.Next() {
		var User models.User

		errScan := rows.Scan(
			&User.UUID,
			&User.Username,
			&User.PasswordHash,
			&User.AccessLevel,
		)

		if errScan != nil {
			return nil, errScan
		}

		Users = append(Users, User)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &Users, nil
}

func (r *UsersRepository) GetUsersByAccessLevel(userAccessLevel int) (*[]models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE access_level<=$1",
		constRepository.AUTH_SCHEMA,
		constRepository.USERS_TABLE,
	)

	ctx := context.Background()
	rows, err := r.Pool.Query(ctx, query, userAccessLevel)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var Users []models.User

	for rows.Next() {
		var User models.User

		errScan := rows.Scan(
			&User.UUID,
			&User.Username,
			&User.PasswordHash,
			&User.AccessLevel,
		)

		if errScan != nil {
			return nil, errScan
		}

		Users = append(Users, User)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &Users, nil
}
