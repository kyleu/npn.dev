module github.com/kyleu/npn/npndatabase

go 1.13

require (
	emperror.dev/errors v0.7.0
	github.com/jackc/pgx/v4 v4.7.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/kyleu/npn/npncore v1.0.0
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npncore => ../npncore
