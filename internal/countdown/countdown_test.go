package countdown

import (
	"testing"
	"time"
)

func TestDaysUntilNewYear(t *testing.T) {
	tests := []struct {
		name string
		from time.Time
		want int
	}{
		{
			name: "начало календарного года",
			from: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: 366, // 2024 — високосный
		},
		{
			name: "конец календарного года",
			from: time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "день перед 29 февраля високосного года",
			from: time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC),
			want: 308,
		},
		{
			name: "29 февраля високосного года",
			from: time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			want: 307,
		},
		{
			name: "обычный год, середина",
			from: time.Date(2023, time.June, 15, 0, 0, 0, 0, time.UTC),
			want: 200,
		},
		{
			name: "время суток игнорируется",
			from: time.Date(2024, time.December, 31, 23, 59, 59, 0, time.UTC),
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DaysUntilNewYear(tt.from)
			if got != tt.want {
				t.Errorf("DaysUntilNewYear(%v) = %d, want %d", tt.from, got, tt.want)
			}
		})
	}
}

func TestDaysUntilNewYearFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "корректная дата", input: "2024-12-31", want: 1},
		{name: "некорректный формат", input: "31.12.2024", wantErr: true},
		{name: "пустая строка", input: "", wantErr: true},
		{name: "несуществующая дата", input: "2023-02-29", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DaysUntilNewYearFromString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}