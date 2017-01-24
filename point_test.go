package polycarp

import (
	"testing"

	"github.com/tidwall/poly"
)

func TestPoint_String(t *testing.T) {
	type fields struct {
		Point poly.Point
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				Point: tt.fields.Point,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Point.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Set(t *testing.T) {
	type fields struct {
		Point poly.Point
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				Point: tt.fields.Point,
			}
			if err := p.Set(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Point.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
