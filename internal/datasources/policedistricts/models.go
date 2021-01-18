package policedistricts

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type PoliceDistricts struct {
	Features *geojson.FeatureCollection
}

func ParsePoliceDistricts(encoded []byte) (*PoliceDistricts, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &PoliceDistricts{
		Features: fc,
	}, nil
}
