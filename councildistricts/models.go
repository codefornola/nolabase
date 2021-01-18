package councildistricts

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type CouncilDistricts struct {
	Features *geojson.FeatureCollection
}

func ParseCouncilDistricts(encoded []byte) (*CouncilDistricts, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &CouncilDistricts{
		Features: fc,
	}, nil
}
