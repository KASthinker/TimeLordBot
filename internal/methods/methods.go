package methods

import (
	"time"

	"github.com/bradfitz/latlong"
)

//TimeZone ...
func TimeZone(Longitude, Latitude float64) (loctime string,  tz string) {
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
