package scraper

import (
	"context"

	"github.com/bhelx/nolabase/infra"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ScraperRepo struct {
	pgConn *pgxpool.Pool
}

func NewScraperRepo(pgConn *pgxpool.Pool) *ScraperRepo {
	return &ScraperRepo{
		pgConn: pgConn,
	}
}

func (r *ScraperRepo) CountUnfinishedJobs() (int, error) {
	sql := `
		SELECT count(id) FROM infra.jobs WHERE state < $1;
	`
	var count int
	err := r.pgConn.QueryRow(context.Background(), sql, infra.JobSucceeded).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ScraperRepo) ListReadyJobs(action func(j *infra.Job) error) error {
	sql := `
		SELECT
			id,
			url,
			scraper_name,
			metadata,
			inserted_at
		FROM infra.jobs WHERE state = $1;
	`
	rows, err := r.pgConn.Query(context.Background(), sql, infra.JobReady)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var j infra.Job
		err := rows.Scan(
			&j.Id,
			&j.Url,
			&j.ScraperName,
			&j.MetaData,
			&j.InsertedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		err = action(&j)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}

func (r *ScraperRepo) PersistJobs(jobs []infra.Job) (newJobs []infra.Job, err error) {
	tx, err := r.pgConn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return newJobs, err
	}

	for _, j := range jobs {
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
		RETURNING id;
		`
		err := tx.QueryRow(
			context.Background(),
			sql,
			j.Url,
			j.ScraperName,
			j.MetaData,
			j.State,
		).Scan(&j.Id)
		log.Debug("Stored job ", j.Id)
		if err != nil {
			return newJobs, err
		}

		job := j
		newJobs = append(newJobs, job)
	}

	err = tx.Commit(context.Background())
	return newJobs, err
}

func (r *ScraperRepo) MarkJobFailed(id int, erro string) error {
	sql := `
		UPDATE infra.jobs SET state = $1, error = $2 WHERE id = $3;
	`
	_, err := r.pgConn.Exec(context.Background(), sql, infra.JobFailed, erro, id)
	return err

}
func (r *ScraperRepo) MarkJob(id int, state infra.JobState) error {
	sql := `
		UPDATE infra.jobs SET state = $1 WHERE id = $2;
	`
	_, err := r.pgConn.Exec(context.Background(), sql, state, id)
	return err
}

func (r *ScraperRepo) PurgeCompletedJobs() error {
	sql := `
		DELETE FROM infra.jobs WHERE state = $1;
	`
	_, err := r.pgConn.Exec(context.Background(), sql, infra.JobSucceeded)
	return err
}
