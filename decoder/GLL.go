package decoder

import (
	"github.com/jbaikge/nmea-relay/convert"
)

func init() {
	RegisterDecoder("GLL", GLL)
}

// 0903.892,S,04141.744,W,171742.852,V
func GLL(s []byte) (m Message, err error) {
	m.Type = "GLL"
	m.Description = "Lat/Lon data"
	m.Latitude = convert.NMEALat(s[:10])
	m.Longitude = convert.NMEALon(s[11:22])
	return
}
