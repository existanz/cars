package database

import (
	model "cars/internal/models"
	"cars/internal/rest/query"
	"database/sql"

	"log/slog"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type CardyB interface {
	AddNewCar(car model.Car) error
	GetCars(paginator query.Paginator, filters []query.Filter) ([]model.Car, error)
	GetCarById(id string) (model.Car, error)
	GetCarByRegNum(regNum string) (model.Car, error)
	GetCarsByOwner(id int) ([]model.Car, error)
	DeleteCarById(id string) error
	UpdateCarById(id int, car model.Car) error
	AddNewPeople(people model.People) (int, error)
	UpdatePeopleById(id int, people model.People) error
	DeletePeopleById(id int) error
}

type PostgresDB struct {
	*sql.DB
}

func (db PostgresDB) AddNewCar(car model.Car) error {
	id, err := db.AddNewPeople(car.Owner)
	if err != nil {
		return err
	}
	car.Owner.Id = id
	_, err = db.Exec("INSERT INTO cars (reg_num, mark, model, year, owner_id) VALUES ($1, $2, $3, $4, $5)", car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Id)
	if err != nil {
		return err
	}
	return nil
}

func (db PostgresDB) GetCars(paginator query.Paginator, filters []query.Filter) ([]model.Car, error) {
	var cars []model.Car
	query := "SELECT * FROM cars left join peoples on cars.owner_id = peoples.id" + getFilersString(filters) + getPaginatorString(paginator)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Id, &car.Owner.Id, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func getFilersString(filters []query.Filter) string {
	filtersString := " WHERE "
	if len(filters) == 0 {
		return ""
	}
	for _, filter := range filters {
		if filter.Field == "year" && strings.Contains(filter.Value, ":") {
			from_to := strings.Split(filter.Value, ":")
			filtersString += filter.Field + " BETWEEN " + from_to[0] + " AND " + from_to[1] + " AND "
		} else {
			filtersString += filter.Field + " = '" + filter.Value + "' AND "
		}
	}
	filtersString = strings.TrimSuffix(filtersString, "AND ")
	return filtersString
}

func getPaginatorString(paginator query.Paginator) string {
	page, _ := strconv.Atoi(paginator.Page)
	limit, _ := strconv.Atoi(paginator.Limit)
	offset := (page - 1) * limit
	return " OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(limit)
}

func (db PostgresDB) GetCarById(id string) (model.Car, error) {
	var car model.Car
	row := db.QueryRow("SELECT * FROM cars left join peoples on cars.owner_id = peoples.id WHERE cars.id = $1", id)
	if err := row.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Id, &car.Owner.Id, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic); err != nil {
		return model.Car{}, err
	}
	return car, nil
}

func (db PostgresDB) GetCarByRegNum(regNum string) (model.Car, error) {
	var car model.Car
	row := db.QueryRow("SELECT * FROM cars WHERE reg_num = $1", regNum)
	if err := row.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Id); err != nil {
		return model.Car{}, err
	}
	return car, nil
}

func (db PostgresDB) GetCarsByOwner(id int) ([]model.Car, error) {
	var cars []model.Car
	rows, err := db.Query("SELECT * FROM cars WHERE owner_id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Id); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (db PostgresDB) DeleteCarById(id string) error {
	_, err := db.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (db PostgresDB) UpdateCarById(id int, car model.Car) error {
	_, err := db.Exec("UPDATE cars SET reg_num = $1, mark = $2, model = $3, year = $4, owner_id = $5 WHERE id = $6", car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Id, id)
	if err != nil {
		return err
	}
	return nil
}
func (db PostgresDB) AddNewPeople(people model.People) (int, error) {
	var id int
	rows, err := db.Query("SELECT TOP 1 id FROM peoples WHERE name = $1 AND surname = $2 AND patronymic = $3", people.Name, people.Surname, people.Patronymic)
	if err == nil && rows.Next() {
		rows.Scan(&id)
		return id, nil
	}
	err = db.QueryRow("INSERT INTO peoples (name, surname, patronymic) VALUES ($1, $2, $3) RETURNING id", people.Name, people.Surname, people.Patronymic).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (db PostgresDB) UpdatePeopleById(id int, people model.People) error {
	_, err := db.Exec("UPDATE peoples SET name = $1, surname = $2, patronymic = $3 WHERE id = $4", people.Name, people.Surname, people.Patronymic, id)
	if err != nil {
		return err
	}
	return nil
}

func (db PostgresDB) DeletePeopleById(id int) error {
	_, err := db.Exec("DELETE FROM people WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewPostgresDB(psqlInfo string) (PostgresDB, error) {
	slog.Info("Connecting to database: ", "connection string", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		slog.Error("Cold not connect to database: ", err)
		return PostgresDB{nil}, err
	}
	err = db.Ping()
	if err != nil {
		return PostgresDB{nil}, err
	}
	slog.Info("Connected! \n")
	MigrateUp(db)
	return PostgresDB{db}, nil
}

func MigrateUp(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Debug("Can't create driver for migration : %v\n", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver)
	if err != nil {
		slog.Debug("Can't create migration : %v\n", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Debug("Migration error: %v", err)
	}
}
