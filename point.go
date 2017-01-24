package polycarp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tidwall/poly"
)

func String(poly.Point) string {
	return fmt.Sprintf("%s,%s", p.X, p.Y)
}

func Set(s string) (poly.Point, error) {
	p := poly.Point{}
	ps := strings.Split(s, ",")
	x, err := strconv.ParseFloat(ps[0], 64)
	if err != nil {
		return nil, fmt.Errorf("parsefloat: %v is not a float", ps[0])
	}
	y, err := strconv.ParseFloat(ps[1], 64)
	if err != nil {
		return nil, fmt.Errorf("parsefloat: %v is not a float", ps[1])
	}
	p.X = x
	p.Y = y
	return p, nil
}
