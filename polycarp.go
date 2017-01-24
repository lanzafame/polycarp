package polycarp

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/tidwall/poly"
)

func FindSubset(ps []Point, set poly.Rect) []Point {
	var subset []Point
	for _, p := range ps {
		if p.InsideRect(set) {
			subset = append(subset, p)
		}
	}
	return subset
}

func SubsetCount(ps []Point, set poly.Rect) int {
	return len(FindSubset(ps, set))
}

func PointsToRecords(ps []Point) [][]string {
	var records [][]string
	records = append(records, []string{"lat", "lon"})
	for _, p := range ps {
		lat := strconv.FormatFloat(p.X, 'f', -1, 64)
		lon := strconv.FormatFloat(p.Y, 'f', -1, 64)
		records = append(records, []string{lat, lon})
	}
	return records
}

func RecordsToPoints(records [][]string) ([]Point, error) {
	records = records[1:]
	var ps []Point
	for i, v := range records {
		lat, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			return nil, fmt.Errorf("Line %v contains non-float", i)
		}
		lon, err := strconv.ParseFloat(v[1], 64)
		if err != nil {
			return nil, fmt.Errorf("Line %v contains non-float", i)
		}
		p := Point{poly.Point{lat, lon}}
		ps = append(ps, p)
	}
	return ps, nil
}

func ReadCSVToRecords(f string) ([][]string, error) {
	csvf, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(csvf)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func MustReadCSVToRecords(f string) [][]string {
	records, err := ReadCSVToRecords(f)
	if err != nil {
		panic(err)
	}
	return records
}

func MustRecordsToPoints(records [][]string) []Point {
	ps, err := RecordsToPoints(records)
	if err != nil {
		panic(err)
	}
	return ps
}

func MustReadCSVToPoints(f string) []Point {
	records := MustReadCSVToRecords(f)
	ps := MustRecordsToPoints(records)
	return ps
}

func WriteRecordsToCSV(f string, records [][]string) error {
	outf, err := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	w := csv.NewWriter(outf)
	w.WriteAll(records)
	if err := w.Error(); err != nil {
		return fmt.Errorf("error writing csv:", err)
	}
	return nil
}
