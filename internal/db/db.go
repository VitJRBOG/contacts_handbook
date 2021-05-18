package db

import (
	"database/sql"
	"fmt"
	"log"
	"runtime/debug"
	"strings"

	"github.com/VitJRBOG/contacts_handbook/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(dbConn config.DBConn) (*sql.DB, error) {
	c := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		dbConn.Login, dbConn.Password, dbConn.Address, dbConn.DBName)
	db, err := sql.Open("mysql", c)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Contact struct {
	ID          int
	Name        string
	PhoneNumber string
}

func (c *Contact) InsertInto(db *sql.DB) (int, int, error) {
	query := `INSERT INTO contact(name, phonenumber) values(?, ?)`

	result, err := db.Exec(query, c.Name, c.PhoneNumber)
	if err != nil {
		return 0, 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	return int(id), int(count), nil
}

func (c *Contact) SelectFrom(db *sql.DB) ([]Contact, error) {
	query := `SELECT * FROM contact`

	var f []interface{}

	if c.ID > 0 {
		query += ` WHERE id = ?`
		f = append(f, c.ID)
	}

	if c.Name != "" {
		if strings.Contains(query, "WHERE") {
			query += ` AND name = ?`
		} else {
			query += ` WHERE name = ?`
		}
		f = append(f, c.Name)
	}

	if c.PhoneNumber != "" {
		if strings.Contains(query, "WHERE") {
			query += ` AND phonenumber = ?`
		} else {
			query += ` WHERE phonenumber = ?`
		}
		f = append(f, c.PhoneNumber)
	}

	rows, err := db.Query(query, f...)
	if err != nil {
		return []Contact{}, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("%s\n%s\n", err, debug.Stack())
		}
	}()

	var contacts []Contact

	for rows.Next() {
		var contact Contact

		if err := rows.Scan(&contact.ID, &contact.Name, &contact.PhoneNumber); err != nil {
			return []Contact{}, err
		}

		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return []Contact{}, err
	}

	return contacts, nil
}

func (c *Contact) Update(db *sql.DB) (int, int, error) {
	query := `UPDATE contact SET name = ?, phonenumber = ? WHERE id = ?`

	result, err := db.Exec(query, c.Name, c.PhoneNumber, c.ID)
	if err != nil {
		return 0, 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	return int(id), int(count), nil
}

func (c *Contact) Delete(db *sql.DB) (int, int, error) {
	query := `DELETE FROM contact WHERE id = ?`

	result, err := db.Exec(query, c.ID)
	if err != nil {
		return 0, 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	return int(id), int(count), nil
}
