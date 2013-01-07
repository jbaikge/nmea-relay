package convert

// Expects []byte("4051.614,N")
func NMEALat(n []byte) (lat float64) {
	if len(n) != 10 {
		return
	}
	lat = toFloat(n[:2]) + toFloat(n[2:4])/60 + toFloat(n[5:8])/3600
	if n[len(n)-1] == 'S' {
		lat *= -1
	}
	return
}

// Expects []byte("11944.397,E")
func NMEALon(n []byte) (lon float64) {
	if len(n) != 11 {
		return
	}
	lon = toFloat(n[:3]) + toFloat(n[3:5])/60 + toFloat(n[6:9])/3600
	if n[len(n)-1] == 'W' {
		lon *= -1
	}
	return
}

func toFloat(b []byte) (i float64) {
	for _, c := range b {
		i = i*10 + float64(c-'0')
	}
	return
}
