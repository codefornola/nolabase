package policesubzones

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type PoliceSubzones struct {
	Features *geojson.FeatureCollection
}

func ParsePoliceSubzones(encoded []byte) (*PoliceSubzones, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &PoliceSubzones{
		Features: fc,
	}, nil
}
