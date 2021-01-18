package votingprecincts

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
)

type Repo struct {
	conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (r *Repo) StoreVotingPrecincts(precincts *VotingPrecincts) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, feature := range precincts.Features.Features {
		sql := `
		INSERT INTO
			geometries.voting_precincts (
				name,
				registered_voters,
				registered_voters_white,
				registered_voters_black,
				registered_voters_other,
				registered_democrats,
				registered_democrats_white,
				registered_democrats_black,
				registered_republicans,
				registered_republicans_white,
				registered_republicans_black,
				registered_republicans_other,
				registered_other_parties,
				registered_other_parties_white,
				registered_other_parties_black,
				registered_other_parties_other,
				geom
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8,
				$9, $10, $11, $12, $13, $14,
				$15, $16, $17);
		`
		ewkb, err := ewkbhex.Encode(feature.Geometry, ewkbhex.NDR)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx,
			sql,
			feature.Properties["VotingPrecinctsPRECINCTID"],
			feature.Properties["RegisteredVoterInformationRegisteredVotersTotal"],
			feature.Properties["RegisteredVoterInformationRegisteredVotersWhite"],
			feature.Properties["RegisteredVoterInformationRegisteredVotersBlack"],
			feature.Properties["RegisteredVoterInformationRegisteredVotersOther"],
			feature.Properties["RegisteredVoterInformationRegisteredDemocratsTotal"],
			feature.Properties["RegisteredVoterInformationRegisteredDemocratsWhite"],
			feature.Properties["RegisteredVoterInformationRegisteredDemocratsBlack"],
			feature.Properties["RegisteredVoterInformationRegisteredRepublicansTotal"],
			feature.Properties["RegisteredVoterInformationRegisteredRepublicansWhite"],
			feature.Properties["RegisteredVoterInformationRegisteredRepublicansBlack"],
			feature.Properties["RegisteredVoterInformationRegisteredRepublicansOther"],
			feature.Properties["RegisteredVoterInformationOtherPartiesTotal"],
			feature.Properties["RegisteredVoterInformationOtherPartiesWhite"],
			feature.Properties["RegisteredVoterInformationOtherPartiesBlack"],
			feature.Properties["RegisteredVoterInformationOtherPartiesOther"],
			ewkb,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
}
