package votingprecincts

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type VotingPrecincts struct {
	Features *geojson.FeatureCollection
}

func ParseVotingPrecincts(encoded []byte) (*VotingPrecincts, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &VotingPrecincts{
		Features: fc,
	}, nil
}
