package assessor

import (
	"github.com/cridenour/go-postgis"
	"time"
)

type Property struct {
	Id int
	AssessorId string
	OwnerName string
	LocationAddress string
	MailingAddress string
	PropertyClass string
	TaxBillNumber string
	ParcelNo string
	AssessmentArea string
	LandAreaSqFt *int
	BuildingAreaSqFt *int
	LngLatPoint postgis.PointS
	InsertedAt time.Time
	UpdatedAt time.Time
}

type PropertyValue struct {
	Year *int
	LandValue *int
	BuildingValue *int
	TotalValue *int
	AssessedLandValue *int
	AssessedBuildingValue *int
	TotalAssessedValue *int
	HomesteadExemptionValue *int
	TaxableAssessment *int
	AgeFreeze *int
	DisabilityFreeze *int
	AssessmentChange *int
	TaxContract *int
	InsertedAt time.Time
	UpdatedAt time.Time
}

type PropertySale struct {
	Grantor string
	Grantee string
	NotarialArchiveNumber string
	InstrumentNumber string
	Date string
	Price *int
	InsertedAt time.Time
	UpdatedAt time.Time
}

type AsssessorId struct {
	AsssessorId string
	InsertedAt *time.Time
	ScrapedAt *time.Time
}