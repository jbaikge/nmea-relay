package decoder

import (
	"github.com/jbaikge/nmea-relay/convert"
)

func init() {
	RegisterDecoder("GLL", GLL)
}

// 0903.892,S,04141.744,W,171742.852,V
// 3848.3469,N,07702.9541,W,221534,A
func GLL(s []byte) (m Message, err error) {
	var latIdx, lonIdx, commas int
	for i, b := range s {
		if b == ',' {
			commas++
			switch commas {
			case 2:
				latIdx = i
			case 4:
				lonIdx = i
			}
		}
	}

	m.Type = "GLL"
	m.Description = "Lat/Lon data"
	m.Latitude = convert.NMEALat(s[:latIdx])
	m.Longitude = convert.NMEALon(s[latIdx+1 : lonIdx])
	return
}
