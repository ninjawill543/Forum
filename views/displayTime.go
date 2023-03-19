package forum

import (
	"fmt"
	"strings"
	"time"
)

func DisplayTime(creationDate string, separator string) string {
	timeToday := time.Now().String()
	timeTodayHour := strings.Split(timeToday, " ")[0]
	timeDate := strings.Split(creationDate, separator)[0]

	if timeDate == timeTodayHour {
		hour := strings.Split(creationDate, separator)[1]
		return strings.Split(hour, ".")[0]
	} else {
		checkYearNow := strings.Split(timeToday, "-")[0]
		checkYearCreationDate := strings.Split(creationDate, "-")[0]

		checkMonthNow := strings.Split(timeToday, "-")[1]
		checkMonthCreationDate := strings.Split(creationDate, "-")[1]

		checkDayNow := strings.Split(timeToday, "-")[2]
		checkDayNow = strings.Split(checkDayNow, " ")[0]

		checkDayCreationDate := strings.Split(creationDate, "-")[2]
		checkDayCreationDate = strings.Split(checkDayCreationDate, separator)[0]
		if checkYearCreationDate != checkYearNow {
			returnDate := (Atoi(checkYearNow) - Atoi(checkYearCreationDate))
			if returnDate != 1 {
				returnDateString := fmt.Sprintf("%d years ago", returnDate)
				return returnDateString
			} else {
				returnDateString := fmt.Sprintf("%d year ago", returnDate)
				return returnDateString
			}
		} else if checkMonthNow != checkMonthCreationDate {
			returnDate := (Atoi(checkMonthNow) - Atoi(checkMonthCreationDate))
			if returnDate != 1 {
				returnDateString := fmt.Sprintf("%d months ago", returnDate)
				return returnDateString
			} else {
				returnDateString := fmt.Sprintf("%d month ago", returnDate)
				return returnDateString
			}
		} else {
			returnDate := (Atoi(checkDayNow) - Atoi(checkDayCreationDate))
			if returnDate != 1 {
				returnDateString := fmt.Sprintf("%d days ago", returnDate)
				return returnDateString
			} else {
				returnDateString := fmt.Sprintf("%d day ago", returnDate)
				return returnDateString
			}
		}
	}
}
