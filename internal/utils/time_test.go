package utils

import (
	"testing"
	"time"
)

func TestFormatISO8601(t *testing.T) {
	testTime := time.Date(2025, 1, 15, 10, 30, 45, 0, time.UTC)
	got := FormatISO8601(testTime)
	want := "2025-01-15T10:30:45Z"

	if got != want {
		t.Errorf("FormatISO8601() = %v, want %v", got, want)
	}
}

func TestFormatAMZDate(t *testing.T) {
	testTime := time.Date(2025, 1, 15, 10, 30, 45, 0, time.UTC)
	got := FormatAMZDate(testTime)
	want := "20250115T103045Z"

	if got != want {
		t.Errorf("FormatAMZDate() = %v, want %v", got, want)
	}
}

func TestParseISO8601(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		check   func(time.Time) bool
	}{
		{
			name:    "valid ISO8601",
			input:   "2025-01-15T10:30:45Z",
			wantErr: false,
			check: func(t time.Time) bool {
				return t.Year() == 2025 && t.Month() == 1 && t.Day() == 15 &&
					t.Hour() == 10 && t.Minute() == 30 && t.Second() == 45
			},
		},
		{
			name:    "with timezone offset",
			input:   "2025-01-15T10:30:45+08:00",
			wantErr: false,
			check: func(t time.Time) bool {
				// 转换为 UTC 应该是 02:30:45
				utc := t.UTC()
				return utc.Hour() == 2 && utc.Minute() == 30
			},
		},
		{
			name:    "invalid format",
			input:   "2025-01-15 10:30:45",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseISO8601(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseISO8601() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.check != nil {
				if !tt.check(got) {
					t.Errorf("ParseISO8601() result check failed for %v", got)
				}
			}
		})
	}
}

func TestIsExpired(t *testing.T) {
	tests := []struct {
		name   string
		time   time.Time
		buffer time.Duration
		want   bool
	}{
		{
			name:   "expired without buffer",
			time:   time.Now().Add(-1 * time.Hour),
			buffer: 0,
			want:   true,
		},
		{
			name:   "not expired without buffer",
			time:   time.Now().Add(1 * time.Hour),
			buffer: 0,
			want:   false,
		},
		{
			name:   "expired with buffer",
			time:   time.Now().Add(30 * time.Second),
			buffer: 60 * time.Second,
			want:   true,
		},
		{
			name:   "not expired with buffer",
			time:   time.Now().Add(2 * time.Minute),
			buffer: 60 * time.Second,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsExpired(tt.time, tt.buffer)
			if got != tt.want {
				t.Errorf("IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateExpireTime(t *testing.T) {
	tests := []struct {
		name      string
		expiresIn int
		checkFunc func(time.Time) bool
	}{
		{
			name:      "3600 seconds (1 hour)",
			expiresIn: 3600,
			checkFunc: func(t time.Time) bool {
				expected := time.Now().Add(3600 * time.Second)
				diff := t.Sub(expected)
				// 允许 1 秒的误差
				return diff >= -time.Second && diff <= time.Second
			},
		},
		{
			name:      "60 seconds (1 minute)",
			expiresIn: 60,
			checkFunc: func(t time.Time) bool {
				expected := time.Now().Add(60 * time.Second)
				diff := t.Sub(expected)
				return diff >= -time.Second && diff <= time.Second
			},
		},
		{
			name:      "0 seconds",
			expiresIn: 0,
			checkFunc: func(t time.Time) bool {
				diff := time.Now().Sub(t)
				return diff >= 0 && diff <= time.Second
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateExpireTime(tt.expiresIn)
			if !tt.checkFunc(got) {
				t.Errorf("CalculateExpireTime() check failed for result %v", got)
			}
		})
	}
}

func TestFormatISO8601_RoundTrip(t *testing.T) {
	// 测试 Format 和 Parse 的往返转换
	original := time.Date(2025, 1, 15, 10, 30, 45, 0, time.UTC)
	formatted := FormatISO8601(original)
	parsed, err := ParseISO8601(formatted)

	if err != nil {
		t.Fatalf("ParseISO8601() error = %v", err)
	}

	if !original.Equal(parsed) {
		t.Errorf("Round trip failed: original = %v, parsed = %v", original, parsed)
	}
}

