package dateutils

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"

func GetNowMx() time.Time {
	loc, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		// TODO: Handle time location error
	}
	return time.Now().In(loc)
}

func GetNowMxString() string {
	return GetNowMx().Format(apiDateLayout)
}
