package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lanzafame/polycarp"
	"github.com/tidwall/poly"
)

var travelArea poly.Rect
var minPoint polycarp.Point
var maxPoint polycarp.Point
var inputCSV string
var outputCSV string

var usage = `polycarp -min <lat,lon> -max <lat,lon> -i <csv file in> -o <csv file out>`

const (
	inputusage  = "the file to read in"
	outputusage = "the file to write to"
	minusage    = "the min lat/lon point, format: float64,float64, i.e. -27.5999,152.8387"
	maxusage    = "the max lat/lon point, format: float64,float64, i.e. -27.5999,152.8387"
)

func init() {
	travelArea = poly.Rect{minPoint.Point, maxPoint.Point}

	flag.StringVar(&inputCSV, "i", "", inputusage)
	flag.StringVar(&outputCSV, "o", "", outputusage)
	flag.Var(&minPoint, "min", minusage)
	flag.Var(&maxPoint, "max", maxusage)
}

func main() {
	flag.Parse()
	if flag.NFlag() < 4 {
		fmt.Printf("%v\n\n%v%v%v%v\n", usage, minusage, maxusage, inputusage, outputusage)
		os.Exit(-1)
	}

	records, err := polycarp.ReadCSVToRecords(inputCSV)
	if err != nil {
		panic(err)
	}

	ps, err := polycarp.RecordsToPoints(records)
	if err != nil {
		panic(err)
	}
	psInTravelArea := polycarp.FindSubset(ps, travelArea)

	outrecords := polycarp.PointsToRecords(psInTravelArea)
	err = polycarp.WriteRecordsToCSV(outputCSV, outrecords)
}
