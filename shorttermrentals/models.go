package shorttermrentals

import (
	"github.com/cridenour/go-postgis"
	"time"
)

type Permit struct {
	Address            string `json:"address"`
	LicenseNumber      string `json:"license_number"`
	LicenseType        string `json:"license_type"`
	ResidentialSubtype string `json:"residential_subtype"`
	CurrentStatus      string `json:"current_status"`
	Expired            string `json:"expired"`
	Link               string `json:"link"`
	ReferenceCode      string `json:"reference_code"`
	X              	   *float64 `json:",string"`
	Y                  *float64 `json:",string"`
	LngLatPoint        postgis.PointS
	ContactName        string `json:"contact_name"`
	ContactPhone       string `json:"contact_phone"`
	ContactEmail       string `json:"contact_email"`
	LicenseHolderName  string `json:"license_holder_name"`

	ApplicationDate    PermitTime `json:"application_date"`
	ExpirationDate     PermitTime `json:"expiration_date"`
	IssueDate    	   PermitTime `json:"issue_date"`
	OperatorPermitNumber       string `json:"operator_permit_number"`

	BedroomLimit   	    *int64 `json:"bedroom_limit,string"`
	GuestOccupancyLimit *int64 `json:"guest_occupancy_limit,string"`

	Location struct {
		Latitude      *float64 `json:",string"`
		Longitude     *float64 `json:",string"`
	} `json:"location"`
}

func (p *Permit) AfterParse() {
	if p.X != nil && p.Y != nil {
		p.LngLatPoint = postgis.PointS{SRID: 3452, X: *p.X, Y: *p.Y}
	} else if p.Location.Latitude != nil && p.Location.Longitude != nil {
		x := *p.X
		y := *p.Y
		// taking a guess that it's NAD83 if above 1000
		if x > 1000 && y > 1000 {
			p.LngLatPoint = postgis.PointS{SRID: 3452, X: x, Y: y}
		} else {
			p.LngLatPoint = postgis.PointS{SRID: 4326, X: x, Y: y}
		}
	}
}

type PermitTime struct {
	time.Time
}

func (m *PermitTime) UnmarshalJSON(data []byte) error {
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
	*m = PermitTime{tt}
	return err
}

