package sdv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	daysPerSeason = 28
	daysPerYear   = 4 * daysPerSeason
)

type Date int

func ParseDate(date string) (Date, error) {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	season, _ := strconv.Atoi(parts[1])
	dayOfMonth, _ := strconv.Atoi(parts[2])

	return Date((year-1)*daysPerYear + season*daysPerSeason + dayOfMonth), nil
}

func ParsePrettyDate(date string) (Date, error) {
	regex := regexp.MustCompile(`^(\d+).* of (.+?), Year (\d+)$`)
	match := regex.FindStringSubmatch(date)
	if match == nil {
		return Date(0), fmt.Errorf("Could not parse '%s' as a date string.", date)
	}

	dayOfMonth, _ := strconv.Atoi(match[1])
	season := 0
	year, _ := strconv.Atoi(match[3])

	switch strings.ToLower(match[1]) {
	case "spring":
		season = 0
	case "summer":
		season = 1
	case "fall":
		season = 2
	case "winter":
		season = 3
	}

	return Date((year-1)*daysPerYear + season*daysPerSeason + dayOfMonth), nil
}

func (d Date) DayOfMonth() int {
	return ((int(d) - 1) % 28) + 1
}

func (d Date) Year() int {
	return ((int(d) - d.DayOfMonth()) / (28 * 4)) + 1
}

func (d Date) Season() int {
	return ((int(d) - d.DayOfMonth()) / 28) - ((d.Year() - 1) * 4)
}

func (d Date) ID() string {
	return fmt.Sprintf("%04d-%d-%02d", d.Year(), d.Season(), d.DayOfMonth())
}

func (d Date) String() string {
	dayOfMonth := d.DayOfMonth()
	season := d.Season()
	year := d.Year()
	suffix := getDayOfMonthSuffix(dayOfMonth)

	return fmt.Sprintf("%d%s of %s, Year %d", dayOfMonth, suffix, getSeasonName(season), year)
}

func getDayOfMonthSuffix(n int) string {
	if n >= 11 && n <= 13 {
		return "th"
	}

	switch n % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func getSeasonName(season int) string {
	switch season {
	case 0:
		return "Spring"

	case 1:
		return "Summer"

	case 2:
		return "Fall"

	case 3:
		return "Winter"

	default:
		return "<season ?>"
	}
}
