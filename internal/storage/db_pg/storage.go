package db_pg

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"otus_highload/internal/domain"
	"otus_highload/internal/storage"
	"time"
)

type PostgresDB struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *PostgresDB {
	return &PostgresDB{
		conn: conn,
	}
}

func (d *PostgresDB) AddUser(ctx context.Context, user domain.User) (domain.User, error) {
	var id string
	err := d.conn.QueryRow(
		ctx,
		"insert into users(id, first_name, surname, birthdate, biography, password_hash, city_id) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		user.ID,
		user.FirstName,
		user.Surname,
		user.Birthdate,
		user.Biography,
		user.PasswordHash,
		user.City.ID,
	).Scan(&id)
	if err != nil {
		return domain.User{}, err
	}

	user.ID = uuid.MustParse(id)

	return user, nil
}

func (d *PostgresDB) GetUserByID(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	var firstName, surname, biography, cityName string
	var birthdate time.Time
	var cityID int

	err := d.conn.QueryRow(
		ctx,
		`
select u.id, first_name, surname, birthdate, biography, city_id, c.name
from users u
         join cities c on c.id = u.city_id where u.id=$1;`,
		userID,
	).Scan(
		&userID,
		&firstName,
		&surname,
		&birthdate,
		&biography,
		&cityID,
		&cityName,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.User{}, storage.ErrNotFound
	} else if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		ID:        userID,
		FirstName: firstName,
		Surname:   surname,
		Birthdate: birthdate,
		Biography: biography,
		City: domain.City{
			ID:   cityID,
			Name: cityName,
		},
	}

	return user, nil
}

func (d *PostgresDB) GetPasswordHashByUserID(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	var passwordHash string
	err := d.conn.QueryRow(
		ctx,
		"select password_hash from users where id=$1",
		userID,
	).Scan(&passwordHash)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.User{}, storage.ErrNotFound
	} else if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		PasswordHash: passwordHash,
	}

	return user, nil
}

func (d *PostgresDB) GetCityByName(ctx context.Context, name string) (domain.City, error) {
	var city domain.City
	err := d.conn.QueryRow(
		ctx,
		"select id, name from cities where name=$1",
		name,
	).Scan(&city.ID, &city.Name)
	if err != nil {
		return domain.City{}, err
	}

	return city, nil
}
