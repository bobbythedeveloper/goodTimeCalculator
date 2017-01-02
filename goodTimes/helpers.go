package hello

import "time"

func getRahuPosition(datetime time.Time) int64 {
	switch datetime.Weekday() {
	case time.Monday:
		return 0
	case time.Tuesday:
		return 5
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 2
	case time.Saturday:
		return 1
	case time.Sunday:
		return 6
	}

	return 1
}
