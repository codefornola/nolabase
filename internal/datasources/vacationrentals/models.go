package vacationrentals

import (
	"github.com/cridenour/go-postgis"
	"time"
)

type VacationRental struct {
	Name           string `json:"name"`
	AddressName    string `json:"address"`
	Type           string `json:"type"`
	ExpirationDate VacationRentalTime `json:"expiration_date"`
	X              *float64 `json:",string"`
	Y              *float64 `json:",string"`
	BedroomLimit   *int64 `json:"bedroom_limit,string"`
	GuestLimit     *int64 `json:"guest_limit,string"`
	LngLatPoint    postgis.PointS
	Location       struct {
		Latitude     *float64 `json:",string"`
		Longitude    *float64 `json:",string"`
		HumanAddress string `json:"human_address"`
	} `json:"location"`
}

func (v *VacationRental) AfterParse() {
	if v.X != nil && v.Y != nil {
		v.LngLatPoint = postgis.PointS{SRID: 3452, X: *v.X, Y: *v.Y}
	} else if v.Location.Latitude != nil && v.Location.Longitude != nil {
		x := *v.X
		y := *v.Y
		// taking a guess that it's NAD83 if above 1000
		if x > 1000 && y > 1000 {
			v.LngLatPoint = postgis.PointS{SRID: 3452, X: x, Y: y}
		} else {
			v.LngLatPoint = postgis.PointS{SRID: 4326, X: x, Y: y}
		}
	}
}

type VacationRentalTime struct {
	time.Time
}

func (m *VacationRentalTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	dataS := string(data)[1:]
	dataS = dataS[:len(dataS)-5] + "Z"
	tt, err := time.Parse(time.RFC3339, dataS)
	if err != nil {
		return err
	}
	*m = VacationRentalTime{tt}
	return err
}

