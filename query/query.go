package query

import (
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	"mp1/query/db"
)

//go:embed schema.sql
var ddl string
var query *db.Queries

func GetAll() ([]db.Project, error) {
	return query.ProjectsGetAll(context.TODO())
}

func GetNoUI() ([]db.Project, error) {
	return query.ProjectsGetNoUI(context.TODO())
}

func GetUI(low int64, high int64) ([]db.Project, error) {
	return query.ProjectsGetUI(context.TODO(), db.ProjectsGetUIParams{
		Points:   low,
		Points_2: high,
	})
}

func Insert(name string, url string, hasUI bool, points int64) error {
	h := int64(0)
	if hasUI {
		h = 1
	}
	return query.ProjectInsert(context.TODO(), db.ProjectInsertParams{
		Name:   name,
		HasUi:  h,
		Points: points,
		Url:    url,
	})
}

func Init() error {
	dbx, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}
	_, err = dbx.ExecContext(context.TODO(), ddl)
	if err != nil {
		return err
	}
	query = db.New(dbx)
	return nil
}
