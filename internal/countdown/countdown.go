// Package countdown предоставляет функции для расчёта количества
// календарных дней до ближайшего Нового года.
package countdown

import (
	"fmt"
	"time"
)

// DaysUntilNewYear возвращает количество календарных дней от даты from
// до 1 января следующего календарного года.
//
// Время суток в from игнорируется: расчёт ведётся по календарным суткам.
func DaysUntilNewYear(from time.Time) int {
	next := time.Date(from.Year()+1, time.January, 1, 0, 0, 0, 0, from.Location())
	current := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
	return int(next.Sub(current).Hours() / 24)
}

// DaysUntilNewYearFromString парсит дату в формате RFC 3339 (YYYY-MM-DD)
// и возвращает количество дней до Нового года.
func DaysUntilNewYearFromString(s string) (int, error) {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return 0, fmt.Errorf("parse date %q: %w", s, err)
	}
	return DaysUntilNewYear(t), nil
}