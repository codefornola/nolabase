package restaurants

import (
	"encoding/json"

	"github.com/twpayne/go-geom/encoding/geojson"
)

type Restaurants struct {
	Features *geojson.FeatureCollection
}

func ParseRestaurants(encoded []byte) (*Restaurants, error) {
	fc := &geojson.FeatureCollection{}
	err := json.Unmarshal(encoded, fc)
	if err != nil {
		return nil, err
	}

	return &Restaurants{
		Features: fc,
	}, nil
}
