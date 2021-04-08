package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

type Phone struct {
	ID     int
	Number string
}

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) Seed() error {
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, number := range data {
		if _, err := insertPhone(d.db, number); err != nil {
			return err
		}
	}
	return nil
}

func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}

	err = createPhonesTable(db)
	if err != nil {
		return err
	}

	return db.Close()
}

func Reset(driverName, dataSource, name string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}

	err = resetDB(db, name)

	if err != nil {
		return err
	}

	return db.Close()
}

func (d *DB) AllPhones() ([]Phone, error) {
	rows, err := d.db.Query("SELECT id, value from phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []Phone
	for rows.Next() {
		var p Phone
		if err = rows.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}

		ret = append(ret, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (d *DB) FindPhone(number string) (*Phone, error) {
	var p Phone
	err := d.db.QueryRow("SELECT id, value FROM phone_numbers WHERE value=$1", number).Scan(&p.ID, &p.Number)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &p, nil
}

func (d *DB) DeletePhone(id int) error {
	statement := "DELETE FROM phone_numbers WHERE id=$1"
	_, err := d.db.Exec(statement, id)
	return err
}

func (d *DB) UpdatePhone(p *Phone) error {
	statement := "UPDATE phone_numbers SET value=$2 WHERE id=$1"
	_, err := d.db.Exec(statement, p.ID, p.Number)
	return err
}

func createPhonesTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)`

	_, err := db.Exec(statement)

	return err
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}

	return nil
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		panic(err)
	}

	return createDB(db, name)
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	var id int

	statement := "INSERT INTO phone_numbers(value) VALUES($1) RETURNING id"
	err := db.QueryRow(statement, phone).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

// Not using this for now
func getPhone(db *sql.DB, id int) (string, error) {
	var number string

	err := db.QueryRow("SELECT value FROM phone_numbers WHERE id=$1", id).Scan(&number)

	if err != nil {
		return "", err
	}

	return number, nil
}
