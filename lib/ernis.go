package ernis

import (
	"errors"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const ernisURL = "http://www.ernis.at/"

var errClosed = errors.New("ernis is closed")

// GetTodaysMeals returns today's meals
func GetTodaysMeals() (string, error) {
	day := time.Now().Weekday()
	return GetMeals(day)
}

// GetMeals returns the meal of this week for the given weekday
func GetMeals(day time.Weekday) (string, error) {
	doc, err := goquery.NewDocument(ernisURL)
	if err != nil {
		return "", err
	}

	if day == time.Saturday || day == time.Sunday {
		return "", errClosed
	}

	box := doc.Find(".inspeisekarte #accordion > div").Eq(int(day) - 1)
	meal := box.Find(".moduletable .custom").Text()
	return strings.TrimSpace(meal), nil
}
