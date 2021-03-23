package repository

import (
	"database/sql"

	"github.com/marcelocpinheiro/communication-engine/entity"
)

type CompanyMySQL struct {
	db *sql.DB
}

func NewCompanyMySQL(db *sql.DB) *CompanyMySQL {
	return &CompanyMySQL{
		db: db,
	}
}

func (r *CompanyMySQL) Create(e *entity.Company) (int64, error) {
	stmt, err := r.db.Prepare(`insert into companies (name, email) values (?, ?)`)
	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(e.Name, e.Email)
	if err != nil {
		return -1, err
	}

	err = stmt.Close()
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
