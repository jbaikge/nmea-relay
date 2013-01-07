package decoder

import (
	"strconv"
)

/*
const (
	AAM int = iota // Waypoint Arrival Alarm
	ALM            // Almanac data
	APA            // Auto Pilot A sentence
	APB            // Auto Pilot B sentence
	BOD            // Bearing Origin to Destination
	BWC            // Bearing using Great Circle route
	DTM            // Datum being used.
	GGA            // Fix information
	GLL            // Lat/Lon data
	GRS            // GPS Range Residuals
	GSA            // Overall Satellite data
	GST            // GPS Pseudorange Noise Statistics
	GSV            // Detailed Satellite data
	MSK            // send control for a beacon receiver
	MSS            // Beacon receiver status information.
	RMA            // recommended Loran data
	RMB            // recommended navigation data for gps
	RMC            // recommended minimum data for gps
	RTE            // route message
	TRF            // Transit Fix Data
	STN            // Multiple Data ID
	VBW            // dual Ground / Water Spped
	VTG            // Vector track an Speed over the Ground
	WCV            // Waypoint closure velocity (Velocity Made Good)
	WPL            // Waypoint Location information
	XTC            // cross track error
	XTE            // measured cross track error
	ZTG            // Zulu (UTC) time and time to go (to destination)
	ZDA            // Date and Time 
)
*/

type Location struct {
	Heading   float64
	Latitude  float64
	Longitude float64
}

func Decode(s string) (l Location, err error) {
	return
}

func Checksum(s string) (sum int64) {
	for i := 0; i < len(s); i++ {
		sum ^= int64(s[i])
	}
	return
}

func Split(s string) (sentence string, checksum int64, err error) {
	checksum, err = strconv.ParseInt(s[len(s)-2:], 16, 0)
	// Trim leading $ and everything after *
	sentence = s[1 : len(s)-3]
	return
}

func ValidChecksum(s string) bool {

	return true
}
