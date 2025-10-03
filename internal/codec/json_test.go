package codec

import (
	"bytes"
	"strings"
	"testing"
)

type TestOrder struct {
	OrderID string  `json:"orderId"`
	Status  string  `json:"status"`
	Total   float64 `json:"total"`
}

func TestNewEncoder(t *testing.T) {
	tests := []struct {
		name string
		opts []EncoderOption
	}{
		{
			name: "default encoder",
			opts: nil,
		},
		{
			name: "encoder with indent",
			opts: []EncoderOption{WithIndent("", "  ")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewEncoder(tt.opts...)
			if encoder == nil {
				t.Error("NewEncoder() returned nil")
			}
		})
	}
}

func TestEncoder_Marshal(t *testing.T) {
	order := TestOrder{
		OrderID: "123",
		Status:  "Shipped",
		Total:   99.99,
	}

	tests := []struct {
		name    string
		encoder *Encoder
		value   interface{}
		wantErr bool
		check   func([]byte) bool
	}{
		{
			name:    "marshal struct",
			encoder: NewEncoder(),
			value:   order,
			wantErr: false,
			check: func(data []byte) bool {
				return strings.Contains(string(data), `"orderId":"123"`)
			},
		},
		{
			name:    "marshal with indent",
			encoder: NewEncoder(WithIndent("", "  ")),
			value:   order,
			wantErr: false,
			check: func(data []byte) bool {
				return strings.Contains(string(data), "\n") &&
					strings.Contains(string(data), `"orderId"`)
			},
		},
		{
			name:    "marshal nil",
			encoder: NewEncoder(),
			value:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.encoder.Marshal(tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.check != nil {
				if !tt.check(data) {
					t.Errorf("Marshal() data check failed: %s", string(data))
				}
			}
		})
	}
}

func TestNewDecoder(t *testing.T) {
	tests := []struct {
		name string
		opts []DecoderOption
	}{
		{
			name: "default decoder",
			opts: nil,
		},
		{
			name: "decoder with disallow unknown fields",
			opts: []DecoderOption{WithDisallowUnknownFields()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := NewDecoder(tt.opts...)
			if decoder == nil {
				t.Error("NewDecoder() returned nil")
			}
		})
	}
}

func TestDecoder_Unmarshal(t *testing.T) {
	validJSON := []byte(`{"orderId":"123","status":"Shipped","total":99.99}`)
	jsonWithUnknownField := []byte(`{"orderId":"123","status":"Shipped","total":99.99,"unknown":"field"}`)

	tests := []struct {
		name    string
		decoder *Decoder
		data    []byte
		wantErr bool
		check   func(*TestOrder) bool
	}{
		{
			name:    "unmarshal valid JSON",
			decoder: NewDecoder(),
			data:    validJSON,
			wantErr: false,
			check: func(order *TestOrder) bool {
				return order.OrderID == "123" &&
					order.Status == "Shipped" &&
					order.Total == 99.99
			},
		},
		{
			name:    "unmarshal with unknown fields (allowed)",
			decoder: NewDecoder(),
			data:    jsonWithUnknownField,
			wantErr: false,
			check: func(order *TestOrder) bool {
				return order.OrderID == "123"
			},
		},
		{
			name:    "unmarshal with unknown fields (disallowed)",
			decoder: NewDecoder(WithDisallowUnknownFields()),
			data:    jsonWithUnknownField,
			wantErr: true,
		},
		{
			name:    "unmarshal empty data",
			decoder: NewDecoder(),
			data:    []byte{},
			wantErr: true,
		},
		{
			name:    "unmarshal into nil",
			decoder: NewDecoder(),
			data:    validJSON,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var order *TestOrder
			if tt.name != "unmarshal into nil" {
				order = &TestOrder{}
			}

			err := tt.decoder.Unmarshal(tt.data, order)

			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.check != nil {
				if !tt.check(order) {
					t.Errorf("Unmarshal() result check failed: %+v", order)
				}
			}
		})
	}
}

func TestDecoder_UnmarshalFromReader(t *testing.T) {
	validJSON := `{"orderId":"123","status":"Shipped","total":99.99}`

	tests := []struct {
		name    string
		decoder *Decoder
		json    string
		wantErr bool
		check   func(*TestOrder) bool
	}{
		{
			name:    "unmarshal from reader",
			decoder: NewDecoder(),
			json:    validJSON,
			wantErr: false,
			check: func(order *TestOrder) bool {
				return order.OrderID == "123"
			},
		},
		{
			name:    "unmarshal from nil reader",
			decoder: NewDecoder(),
			json:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader *bytes.Reader
			if tt.json != "" {
				reader = bytes.NewReader([]byte(tt.json))
			}

			var order *TestOrder
			if tt.name != "unmarshal into nil" {
				order = &TestOrder{}
			}

			var err error
			if reader == nil {
				err = tt.decoder.UnmarshalFromReader(nil, order)
			} else {
				err = tt.decoder.UnmarshalFromReader(reader, order)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalFromReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.check != nil {
				if !tt.check(order) {
					t.Errorf("UnmarshalFromReader() result check failed: %+v", order)
				}
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	order := TestOrder{
		OrderID: "123",
		Status:  "Shipped",
		Total:   99.99,
	}

	data, err := MarshalJSON(order)
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}

	if !strings.Contains(string(data), `"orderId":"123"`) {
		t.Errorf("MarshalJSON() data check failed: %s", string(data))
	}
}

func TestUnmarshalJSON(t *testing.T) {
	validJSON := []byte(`{"orderId":"123","status":"Shipped","total":99.99}`)

	var order TestOrder
	err := UnmarshalJSON(validJSON, &order)
	if err != nil {
		t.Fatalf("UnmarshalJSON() error = %v", err)
	}

	if order.OrderID != "123" || order.Status != "Shipped" || order.Total != 99.99 {
		t.Errorf("UnmarshalJSON() result check failed: %+v", order)
	}
}

func TestMarshalIndentJSON(t *testing.T) {
	order := TestOrder{
		OrderID: "123",
		Status:  "Shipped",
		Total:   99.99,
	}

	data, err := MarshalIndentJSON(order, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndentJSON() error = %v", err)
	}

	if !strings.Contains(string(data), "\n") {
		t.Error("MarshalIndentJSON() should contain newlines")
	}

	if !strings.Contains(string(data), `"orderId"`) {
		t.Errorf("MarshalIndentJSON() data check failed: %s", string(data))
	}
}

func TestEncoder_Marshal_RoundTrip(t *testing.T) {
	original := TestOrder{
		OrderID: "123",
		Status:  "Shipped",
		Total:   99.99,
	}

	// Encode
	encoder := NewEncoder()
	data, err := encoder.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	// Decode
	decoder := NewDecoder()
	var decoded TestOrder
	err = decoder.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	// Compare
	if original.OrderID != decoded.OrderID ||
		original.Status != decoded.Status ||
		original.Total != decoded.Total {
		t.Errorf("Round trip failed: original = %+v, decoded = %+v", original, decoded)
	}
}

func BenchmarkEncoder_Marshal(b *testing.B) {
	order := TestOrder{
		OrderID: "123",
		Status:  "Shipped",
		Total:   99.99,
	}

	encoder := NewEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = encoder.Marshal(order)
	}
}

func BenchmarkDecoder_Unmarshal(b *testing.B) {
	data := []byte(`{"orderId":"123","status":"Shipped","total":99.99}`)
	decoder := NewDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var order TestOrder
		_ = decoder.Unmarshal(data, &order)
	}
}
