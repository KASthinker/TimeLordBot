package methods

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bradfitz/latlong"
)

//TimeZoneGPS ...
func TimeZoneGPS(Longitude, Latitude float64) (loctime string, tz string) {
	tz = latlong.LookupZoneName(Latitude, Longitude)
	utc := time.Now().UTC()
	local := utc
	location, err := time.LoadLocation(tz)
	if err == nil {
		local = local.In(location)
	}
	loctime = local.Format("15:04")
	tz, _ = local.Zone()
	return loctime, tz
}

//TimeZoneManualy ...
func TimeZoneManualy(strtz string) (loctime string, tz string, err error) {
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
	loctime = local.Format("15:04")
	tz, _ = local.Zone()
	return loctime, tz, nil
}
