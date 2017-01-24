package polycarp

import (
	"reflect"
	"testing"

	"github.com/tidwall/poly"
)

func TestFindSubset(t *testing.T) {
	type args struct {
		ps  []Point
		set poly.Rect
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "1in",
			args: args{
				ps:  []Point{Point{poly.Point{2, 2}}},
				set: poly.Rect{poly.Point{1, 1}, poly.Point{3, 3}},
			},
			want: []Point{Point{poly.Point{2, 2}}},
		},
		{
			name: "0in 0out",
			args: args{
				ps:  []Point{Point{poly.Point{3, 3}}},
				set: poly.Rect{poly.Point{1, 1}, poly.Point{2, 2}},
			},
			want: nil,
		},
		{
			name: "4in 1out",
			args: args{
				ps: []Point{
					Point{poly.Point{3, 3}},
					Point{poly.Point{4, 4}},
					Point{poly.Point{5, 5}},
					Point{poly.Point{6, 6}},
					Point{poly.Point{11, 11}},
				},
				set: poly.Rect{poly.Point{1, 1}, poly.Point{10, 10}},
			},
			want: []Point{
				Point{poly.Point{3, 3}},
				Point{poly.Point{4, 4}},
				Point{poly.Point{5, 5}},
				Point{poly.Point{6, 6}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSubset(tt.args.ps, tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointsToRecords(t *testing.T) {
	type args struct {
		ps []Point
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointsToRecords(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointsToRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordsToPoints(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    []Point
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RecordsToPoints(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordsToPoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecordsToPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadCSVToRecords(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCSVToRecords(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSVToRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSVToRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteRecordsToCSV(t *testing.T) {
	type args struct {
		f       string
		records [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteRecordsToCSV(tt.args.f, tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("WriteRecordsToCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubsetCount(t *testing.T) {
	type args struct {
		ps  []Point
		set poly.Rect
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Riverview, QLD to Bowen Hills, QLD",
			args: args{
				ps: MustReadCSVToPoints("./latlon.csv"),
				set: poly.Rect{
					poly.Point{-27.599009,152.838783}, // riverview qld
					poly.Point{-27.450011,152.035551},
				},
			},
			want: 25705,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsetCount(tt.args.ps, tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubsetCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
