package neighborhoods

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type Neighborhoods struct {
	Features *geojson.FeatureCollection
}

func ParseNeighborhoods(encoded []byte) (*Neighborhoods, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &Neighborhoods{
		Features: fc,
	}, nil
}
