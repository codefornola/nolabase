package bikelanes

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type BikeLanes struct {
	Features *geojson.FeatureCollection
}

func ParseBikeLanes(encoded []byte) (*BikeLanes, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &BikeLanes{
		Features: fc,
	}, nil
}
