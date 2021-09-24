package service

//import (
//	"context"
//	"database/sql"
//	kitlog "github.com/go-kit/kit/log"
//)
//
//type repoSQL struct {
//	db *sql.DB
//	logger kitlog.Logger
//}
//
//func NewRepoSQL(db *sql.DB, logger kitlog.Logger) (RepositorySQL, error) {
//	return &repoSQL{
//		db: db,
//		logger: kitlog.With(logger, "repoSQL", "sql"),
//	}, nil
//}
//
//func (repoSQL *repoSQL) UserCreate(ctx context.Context, email string, password string) error {
//	sql := `
//		INSERT INTO users (email, password) VALUES ($1, $2)`
//	_, err := repoSQL.db.ExecContext(ctx, sql, email, password)
//	if err != nil {return err}
//	return nil
//}
//
//func (repoSQL *repoSQL) UserGet(ctx context.Context, id int) (string, error) {
//	var email string
//	sql := `
//		SELECT email FROM users WHERE id=$1`
//	err := repoSQL.db.QueryRow(sql, id).Scan(&email)
//	if err != nil {return "", err}
//	return email, nil
//}
