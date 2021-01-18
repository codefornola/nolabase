package schooldistricts

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type SchoolDistricts struct {
	Features *geojson.FeatureCollection
}

func ParseSchoolDistricts(encoded []byte) (*SchoolDistricts, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &SchoolDistricts{
		Features: fc,
	}, nil
}
