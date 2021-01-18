package callsforservice

import (
	"strconv"
	"time"

	"github.com/cridenour/go-postgis"
)

type ServiceCall struct {
	NopdItem         string      `json:"nopd_item"`
	TypeText         string      `json:"typetext"`
	Priority         string      `json:"priority"`
	InitialType      string      `json:"initialtype"`
	InitialTypeText  string      `json:"initialtypetext"`
	InitialPriority  string      `json:"initialpriority"`
	MapX             string      `json:"mapx"`
	MapY             string      `json:"mapy"`
	TimeCreate       ServiceTime `json:"timecreate"`
	TimeDispatch     ServiceTime `json:"timedispatch"`
	TimeArrive       ServiceTime `json:"timearrive"`
	TimeClosed       ServiceTime `json:"timeclosed"`
	Disposition      string      `json:"disposition"`
	DispositionText  string      `json:"dispositiontext"`
	SelfInitiatedRaw string      `json:"selfinitiated"`
	SelfInitiated    bool
	Beat             string `json:"beat"`
	BlockAddress     string `json:"block_address"`
	Zip              string `json:"zip"`
	PoliceDistrict   string `json:"policedistrict"`
	LngLatPoint      postgis.PointS
	Location         struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"location"`
}

func (c *ServiceCall) AfterParse() {
	if c.Location.Latitude != "" && c.Location.Longitude != "" {
		x, err := strconv.ParseFloat(c.MapX, 64)
		y, err := strconv.ParseFloat(c.MapY, 64)
		if err != nil {
			log.Warn("Could not parse coords x=%s y=%s", c.MapX, c.MapY)
		} else {
			// taking a guess that it's NAD83
			if x > 1000 && y > 1000 {
				c.LngLatPoint = postgis.PointS{SRID: 3452, X: x, Y: y}
			} else {
				c.LngLatPoint = postgis.PointS{SRID: 4326, X: x, Y: y}
			}
		}
	} else {
		if c.MapX != "" && c.MapY != "" {
			x, err := strconv.ParseFloat(c.MapX, 64)
			y, err := strconv.ParseFloat(c.MapY, 64)
			if err != nil || x == 0 || y == 0 {
				log.Warn(err)
				log.Warnf("Could not parse coords x=%s y=%s", c.MapX, c.MapY)
			} else {
				c.LngLatPoint = postgis.PointS{SRID: 3452, X: x, Y: y}
			}
		}
	}

}

type ServiceTime struct {
	time.Time
}

func (m *ServiceTime) UnmarshalJSON(data []byte) error {
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
	*m = ServiceTime{tt}
	return err
}
