package methods

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
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

// CheckTime ...
func CheckTime(strTime string) error {
	strTime = strings.TrimSpace(strTime)
	_, err := time.Parse("15:04", strTime)
	if err != nil {
		return err
	}

	return nil
}

// CheckDate ...
func CheckDate(strDate string) (string, error) {
	strDate = strings.TrimSpace(strDate)
	tm, err := time.Parse("02/01/2006", strDate)
	if err != nil {
		tm, err = time.Parse("02.01.2006", strDate)
		if err != nil {
			return "", err
		}
	}

	return tm.Format("02-01-2006"), nil
}

// ConvDate ...
func ConvDate(strDate, language string) string {
	strDate = strings.TrimSpace(strDate)
	tm, _ := time.Parse("02-01-2006", strDate)
	if language == "en_EN" {
		return tm.Format("02/01/2006")
	} else if language == "ru_RU" {
		return tm.Format("02.01.2006")
	}
	return ""
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] != true {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

// CheckWeekday ...
func CheckWeekday(weekday string) string {
	v := strings.Split(weekday, ",")
	temp := []string{}

	for i := 0; i < len(v); i++ {
		_, ok := data.Weekday[strings.TrimSpace(v[i])]
		if ok {
			temp = append(temp, strings.TrimSpace(v[i]))
		}
	}

	temp = removeDuplicates(temp)
	str := strings.Join(temp[:], ",")

	if len(temp) > 0 {
		return str
	}
	return ""
}
