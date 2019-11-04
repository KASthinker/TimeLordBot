package methods

import (
	"fmt"
	"strconv"
	"time"
	"log"
	"github.com/bradfitz/latlong"
)

//TimeZoneGPS ...
func TimeZoneGPS(Longitude, Latitude float64, timeformat int) (loctime string, tz string) {
	tz = latlong.LookupZoneName(Latitude, Longitude)
	utc := time.Now().UTC()
	local := utc
	location, err := time.LoadLocation(tz)
	if err == nil {
		local = local.In(location)
	}
	if timeformat == 24 {
		log.Printf("\n\n24-hour clock\n%v\n\n\n", timeformat)
		loctime = local.Format("15:04")
	} else if timeformat == 12 {
		log.Printf("\n\n12-hour clock\n%v\n\n\n", timeformat)
		loctime = local.Format("03:04 PM")
	}
	tz, _ = local.Zone()
	return loctime, tz
}

//TimeZoneManually ...
func TimeZoneManually(strtz string, timeformat int) (loctime string, tz string, err error) {
	inttz, err := strconv.Atoi(strtz)
	if err != nil {
		return "", "", err
	}

	if inttz < -12 || inttz > 14 {
		return "", "", err
	}
	inttz *= (-1)

	tz = fmt.Sprintf("Etc/GMT%+d", inttz)
	utc := time.Now().UTC()
	local := utc
	location, err := time.LoadLocation(tz)
	if err == nil {
		local = local.In(location)
	}
	
	if timeformat == 24 {
		log.Printf("\n\n24-hour clock\n%v\n\n\n", timeformat)
		loctime = local.Format("15:04")
	} else if timeformat == 12 {
		log.Printf("\n\n12-hour clock\n%v\n\n\n", timeformat)
		loctime = local.Format("03:04 PM")
	}
	
	tz, _ = local.Zone()
	return loctime, tz, nil
}
