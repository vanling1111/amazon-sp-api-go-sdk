package utils

import (
	"strings"
	"testing"
)

func TestGenerateRequestID(t *testing.T) {
	// 生成多个 ID 并验证唯一性
	ids := make(map[string]bool)
	for i := 0; i < 100; i++ {
		id, err := GenerateRequestID()
		if err != nil {
			t.Fatalf("GenerateRequestID() error = %v", err)
		}

		// 验证长度（32个字符）
		if len(id) != 32 {
			t.Errorf("GenerateRequestID() length = %d, want 32", len(id))
		}

		// 验证是否为十六进制字符串
		for _, c := range id {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				t.Errorf("GenerateRequestID() contains non-hex character: %c", c)
			}
		}

		// 验证唯一性
		if ids[id] {
			t.Errorf("GenerateRequestID() generated duplicate ID: %s", id)
		}
		ids[id] = true
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		maxLen int
		want   string
	}{
		{
			name:   "short string",
			s:      "hello",
			maxLen: 10,
			want:   "hello",
		},
		{
			name:   "exact length",
			s:      "hello",
			maxLen: 5,
			want:   "hello",
		},
		{
			name:   "truncate long string",
			s:      "very long string here",
			maxLen: 10,
			want:   "very lo...",
		},
		{
			name:   "maxLen too small",
			s:      "hello",
			maxLen: 2,
			want:   "he",
		},
		{
			name:   "maxLen 3",
			s:      "hello world",
			maxLen: 3,
			want:   "...",
		},
		{
			name:   "empty string",
			s:      "",
			maxLen: 10,
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TruncateString(tt.s, tt.maxLen)
			if got != tt.want {
				t.Errorf("TruncateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinNonEmpty(t *testing.T) {
	tests := []struct {
		name string
		sep  string
		strs []string
		want string
	}{
		{
			name: "all non-empty",
			sep:  ", ",
			strs: []string{"a", "b", "c"},
			want: "a, b, c",
		},
		{
			name: "with empty strings",
			sep:  ", ",
			strs: []string{"a", "", "b", "", "c"},
			want: "a, b, c",
		},
		{
			name: "all empty",
			sep:  ", ",
			strs: []string{"", "", ""},
			want: "",
		},
		{
			name: "single non-empty",
			sep:  ", ",
			strs: []string{"", "a", ""},
			want: "a",
		},
		{
			name: "different separator",
			sep:  "|",
			strs: []string{"a", "", "b"},
			want: "a|b",
		},
		{
			name: "no strings",
			sep:  ", ",
			strs: []string{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinNonEmpty(tt.sep, tt.strs...)
			if got != tt.want {
				t.Errorf("JoinNonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizeLogString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "token with equals",
			s:    "token=abc123xyz",
			want: "token=***",
		},
		{
			name: "secret with equals",
			s:    "secret=mysecret",
			want: "secret=***",
		},
		{
			name: "password with colon",
			s:    "password:mypassword",
			want: "password:***",
		},
		{
			name: "key with equals",
			s:    "api_key=12345",
			want: "api_key=***",
		},
		{
			name: "safe string",
			s:    "user=john",
			want: "user=john",
		},
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "just word token",
			s:    "token",
			want: "***",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeLogString(tt.s)
			if got != tt.want {
				t.Errorf("SanitizeLogString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaskString(t *testing.T) {
	tests := []struct {
		name         string
		s            string
		visibleStart int
		visibleEnd   int
		want         string
	}{
		{
			name:         "normal case",
			s:            "1234567890",
			visibleStart: 2,
			visibleEnd:   2,
			want:         "12******90",
		},
		{
			name:         "short string",
			s:            "123",
			visibleStart: 1,
			visibleEnd:   1,
			want:         "1*3",
		},
		{
			name:         "too short string",
			s:            "12",
			visibleStart: 2,
			visibleEnd:   2,
			want:         "12",
		},
		{
			name:         "no masking start",
			s:            "1234567890",
			visibleStart: 0,
			visibleEnd:   3,
			want:         "*******890",
		},
		{
			name:         "no masking end",
			s:            "1234567890",
			visibleStart: 3,
			visibleEnd:   0,
			want:         "123*******",
		},
		{
			name:         "access token",
			s:            "Atza|IwEBIDdL8LKjQp_abcdefghijk1234567890",
			visibleStart: 4,
			visibleEnd:   4,
			want:         "Atza*********************************7890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaskString(tt.s, tt.visibleStart, tt.visibleEnd)
			if got != tt.want {
				t.Errorf("MaskString() = %v, want %v", got, tt.want)
			}

			// 验证长度没有变化
			if len(got) != len(tt.s) {
				t.Errorf("MaskString() length = %d, want %d", len(got), len(tt.s))
			}
		})
	}
}

func TestGenerateRequestID_Uniqueness(t *testing.T) {
	// 高并发测试唯一性
	const goroutines = 10
	const idsPerGoroutine = 100
	results := make(chan string, goroutines*idsPerGoroutine)

	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < idsPerGoroutine; j++ {
				id, err := GenerateRequestID()
				if err != nil {
					t.Errorf("GenerateRequestID() error = %v", err)
					return
				}
				results <- id
			}
		}()
	}

	// 收集所有 ID
	ids := make(map[string]bool)
	for i := 0; i < goroutines*idsPerGoroutine; i++ {
		id := <-results
		if ids[id] {
			t.Errorf("Duplicate ID generated: %s", id)
		}
		ids[id] = true
	}

	if len(ids) != goroutines*idsPerGoroutine {
		t.Errorf("Expected %d unique IDs, got %d", goroutines*idsPerGoroutine, len(ids))
	}
}

func TestMaskString_EdgeCases(t *testing.T) {
	tests := []struct {
		name         string
		s            string
		visibleStart int
		visibleEnd   int
	}{
		{
			name:         "visible larger than string",
			s:            "abc",
			visibleStart: 10,
			visibleEnd:   10,
		},
		{
			name:         "zero visible",
			s:            "abcdefghijk",
			visibleStart: 0,
			visibleEnd:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaskString(tt.s, tt.visibleStart, tt.visibleEnd)

			// 确保不会 panic 且有返回值
			if got == "" && tt.s != "" {
				t.Error("MaskString() returned empty string for non-empty input")
			}

			// 验证返回字符串只包含原字符和星号
			for _, c := range got {
				if c != '*' && !strings.ContainsRune(tt.s, c) {
					t.Errorf("MaskString() contains unexpected character: %c", c)
				}
			}
		})
	}
}
