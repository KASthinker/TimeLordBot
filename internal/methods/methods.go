package methods

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bradfitz/latlong"
	"github.com/markbates/pkger"
)

//GetPath returns the path to the main working directory (if path = "")
//or the path to the file inside the application, if it's path is specified.
func GetPath(path string) string {
	info, err := pkger.Info("")
	if err != nil {
		log.Fatalf("Error GetPath(): %v", err)
	}

	return info.Dir + path
}

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
		loctime = local.Format("15:04")
	} else if timeformat == 12 {
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
		loctime = local.Format("15:04")
	} else if timeformat == 12 {
		loctime = local.Format("03:04 PM")
	}

	tz, _ = local.Zone()
	return loctime, tz, nil
}

// LocTime ...
func LocTime(timezone string) (string, error) {
	inttz, err := strconv.Atoi(timezone)
	if err != nil {
		return "", err
	}

	if inttz < -12 || inttz > 14 {
		return "", err
	}
	inttz *= (-1)

	tz := fmt.Sprintf("Etc/GMT%+d", inttz)
	utc := time.Now().UTC()
	local := utc
	location, err := time.LoadLocation(tz)
	if err == nil {
		local = local.In(location)
	}
	loctime := local.Format("15:04")
	return loctime, nil
}

// LocDate ...
func LocDate(timezone string) (string, error) {
	inttz, err := strconv.Atoi(timezone)
	if err != nil {
		return "", err
	}

	if inttz < -12 || inttz > 14 {
		return "", err
	}
	inttz *= (-1)

	tz := fmt.Sprintf("Etc/GMT%+d", inttz)
	utc := time.Now().UTC()
	local := utc
	location, err := time.LoadLocation(tz)
	if err == nil {
		local = local.In(location)
	}

	year, month, day := local.Date()
	date := fmt.Sprintf("%d-%d-%d", year, int(month), day)

	return date, nil
}

// LocWeekday ...
func LocWeekday(timezone string) (string, error) {
	var shortWeekday = map[time.Weekday]string{
		time.Monday:    "Mon",
		time.Tuesday:   "Tue",
		time.Wednesday: "Wed",
		time.Thursday:  "Thu",
		time.Friday:    "Fri",
		time.Saturday:  "Sat",
		time.Sunday:    "Sun",
	}

	inttz, err := strconv.Atoi(timezone)
	if err != nil {
		return "", err
	}

	if inttz < -12 || inttz > 14 {
		return "", err
	}
	inttz *= (-1)

	tz := fmt.Sprintf("Etc/GMT%+d", inttz)
	utc := time.Now().UTC()
	local := utc
	location, err := time.LoadLocation(tz)
	if err == nil {
		local = local.In(location)
	}

	weekday := shortWeekday[local.Weekday()]

	return weekday, nil
}

// ConvTimeFormat ...
func ConvTimeFormat(strTime string, timeFormat int) (string, error) {
	var layout string
	switch timeFormat {
	case 24:
		layout = "15:04"
	case 12:
		layout = "03:04 PM"
	default:
		err := fmt.Errorf("Wrong time format! -> %d", timeFormat)
		return "", err
	}

	t, err := time.Parse(layout, strTime)
	if err != nil {
		var layout string
		switch timeFormat {
		case 24:
			layout = "03:04 PM"
		case 12:
			layout = "15:04"
		}
		t, err = time.Parse(layout, strTime)
		if err != nil {
			return "", err
		}
	}

	return t.Format(layout), nil

}