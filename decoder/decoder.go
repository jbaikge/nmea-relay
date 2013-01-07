package decoder

import (
	"bytes"
	"encoding/hex"
	"errors"
	"time"
)

// AAM - Waypoint Arrival Alarm
// ALM - Almanac data
// APA - Auto Pilot A sentence
// APB - Auto Pilot B sentence
// BOD - Bearing Origin to Destination
// BWC - Bearing using Great Circle route
// DTM - Datum being used.
// GGA - Fix information
// GLL - Lat/Lon data
// GRS - GPS Range Residuals
// GSA - Overall Satellite data
// GST - GPS Pseudorange Noise Statistics
// GSV - Detailed Satellite data
// MSK - send control for a beacon receiver
// MSS - Beacon receiver status information.
// RMA - recommended Loran data
// RMB - recommended navigation data for gps
// RMC - recommended minimum data for gps
// RTE - route message
// TRF - Transit Fix Data
// STN - Multiple Data ID
// VBW - dual Ground / Water Spped
// VTG - Vector track an Speed over the Ground
// WCV - Waypoint closure velocity (Velocity Made Good)
// WPL - Waypoint Location information
// XTC - cross track error
// XTE - measured cross track error
// ZTG - Zulu (UTC) time and time to go (to destination)
// ZDA - Date and Time 

type Decoder struct {
	Type []byte
	Func DecoderFunc
}

// Expects portion of sentence after the type declaration and before the splat.
//
// Example sentence: $GPGLL,0903.892,S,04141.744,W,171742.852,V*2B
// Portion sent to DecoderFunc: 0903.892,S,04141.744,W,171742.852,V
type DecoderFunc func([]byte) (Message, error)

type Message struct {
	Course      float64
	Description string
	Latitude    float64
	Longitude   float64
	Time        time.Time
	Type        string
}

var decoders = []Decoder{}

func Decode(s []byte) (m Message, err error) {
	if !ValidChecksum(s) {
		err = errors.New("Checksum mismatch")
		return
	}
	firstComma := bytes.IndexByte(s, ',')
	t := s[3:firstComma]
	data := s[firstComma+1 : len(s)-3]
	for _, d := range decoders {
		if bytes.Equal(t, d.Type) {
			m, err = d.Func(data)
			break
		}
	}
	return
}

func Checksum(s []byte) (sum byte) {
	for i := 0; i < len(s); i++ {
		sum ^= s[i]
	}
	return
}

func RegisterDecoder(t string, f DecoderFunc) {
	decoders = append(decoders, Decoder{
		Type: []byte(t),
		Func: f,
	})
}

func Split(s []byte) (sentence []byte, checksum byte, err error) {
	// Convert last 2 characters of sentence from hex
	buf := []byte{0}
	if _, err = hex.Decode(buf, s[len(s)-2:]); err != nil {
		return
	}
	checksum = buf[0]
	// Trim leading $ and everything after *
	sentence = s[1 : len(s)-3]
	return
}

func ValidChecksum(s []byte) bool {
	s, c, err := Split(s)
	if err != nil {
		return false
	}
	return c == Checksum(s)
}
