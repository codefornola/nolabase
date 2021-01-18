package infra

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type JobState int

const (
	JobReady     = JobState(iota)
	JobQueued    = JobState(iota)
	JobStarted   = JobState(iota)
	JobSucceeded = JobState(iota)
	JobFailed    = JobState(iota)
)

type Job struct {
	Id          int
	Url         string
	ScraperName string
	InsertedAt  *time.Time
	MetaData    string
	State       JobState
	Error       *string
}

type JobRepo struct {
	pgConn *pgxpool.Pool
}

func NewJobRepo(pgConn *pgxpool.Pool) *JobRepo {
	return &JobRepo{
		pgConn: pgConn,
	}
}

func (r *JobRepo) ListJobs(state JobState, action func(job *Job)) error {
	sql := `
	SELECT
		id,
		url,
		scraper_name,
		metadata,
		state,
		error,
		inserted_at
	FROM infra.jobs WHERE state = $1;
	`
	rows, err := r.pgConn.Query(context.Background(), sql, state)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var j Job
		err := rows.Scan(
			&j.Id,
			&j.Url,
			&j.ScraperName,
			&j.MetaData,
			&j.State,
			&j.Error,
			&j.InsertedAt,
		)
		if err != nil {
			return err
		}
		action(&j)
	}

	return nil
}

func (r *JobRepo) GetJob(id int) (job Job, err error) {
	sql := `
		SELECT
			id,
			url,
			scraper_name,
			metadata,
			state,
			error,
			inserted_at
		FROM infra.jobs WHERE id = $1;
	`
	err = r.pgConn.QueryRow(
		context.Background(),
		sql,
		id,
	).Scan(
		&job.Id,
		&job.Url,
		&job.ScraperName,
		&job.MetaData,
		&job.State,
		&job.Error,
		&job.InsertedAt,
	)
	if err != nil {
		return job, err
	}

	return job, nil
}

func (r *JobRepo) NewJobCreator() (*JobCreator, error) {
	tx, err := r.pgConn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &JobCreator{
		tx: tx,
	}, nil
}

func (r *JobRepo) AddJob(j *Job) (err error) {
	jobCreator, err := r.NewJobCreator()
	if err != nil {
		return err
	}
	err = jobCreator.AddJob(j)
	if err != nil {
		return err
	}
	err = jobCreator.Commit()
	return err
}

type JobCreator struct {
	tx pgx.Tx
}

func (r *JobCreator) AddJob(j *Job) (err error) {
	ctx := context.Background()
	sql := `
		INSERT INTO
			infra.jobs (
				url,
				scraper_name,
				metadata,
				state,
				inserted_at
			)
		VALUES ($1, $2, $3, $4, now())
		ON CONFLICT (url, metadata) DO NOTHING;
	`
	_, err = r.tx.Exec(ctx, sql,
		j.Url,
		j.ScraperName,
		j.MetaData,
		JobReady,
	)
	if err != nil {
		r.tx.Rollback(ctx)
		return err
	}
	return nil
}

func (r *JobCreator) Commit() error {
	return r.tx.Commit(context.Background())
}
