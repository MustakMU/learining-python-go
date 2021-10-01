package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
	"github.com/tebeka/selenium"
)

const (
	//Set constants separately chromedriver.exe Address and local call port of
	seleniumPath = `/Users/mustakmuntawale/Downloads/chromedriver`
	port         = 9515
)

type FlightDetails struct {
	arrivalDeparture string
	departureTime    string
	cost             int
	journeyDuration  string
}

type FlightDetailsList []FlightDetails

func (det FlightDetails) print() {
	fmt.Print("Arrival: ", det.arrivalDeparture)
	fmt.Print(" Departure: ", det.departureTime)
	fmt.Print(" Cost: ", det.cost)
	fmt.Println(" Journey time: ", det.journeyDuration)
}

func main() {

	from_loc := "bengaluru"
	to_loc := "chennai"
	date := "30/10/2021"
	url := "https://www.expedia.co.in/Flights-Search?trip=oneway&leg1=from:" + from_loc + ",to:" + to_loc + ",departure:" + date + "TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass:economy&mode=search"

	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	check(err)

	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//Call browser urlPrefix: test reference: defaulturlprefix =“ http://127.0.0.1 :4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	check(err)

	//Delay exiting chrome
	defer wd.Quit()

	println(url)
	err = wd.Get(url)
	check(err)

	time.Sleep(4 * time.Second)
	//Find Baidu input box id
	costInfo, err := wd.FindElements(selenium.ByCSSSelector, "span[class='uitk-price-a11y is-visually-hidden'")
	check(err)

	costList := []int{}
	for _, we := range costInfo {
		text, err := we.Text()
		check(err)

		costStr := strings.Split(strings.Split(text, " ")[0], "Rs")[1]
		costStr = strings.ReplaceAll(costStr, ",", "")
		cost, err := strconv.Atoi(costStr)
		costList = append(costList, cost)
	}

	arrivalInfo, err := wd.FindElements(selenium.ByCSSSelector, "div[class='uitk-text truncate uitk-type-200 uitk-spacing uitk-spacing-margin-blockstart-one uitk-text-emphasis-theme'")
	arrivalDepartureList := []string{}

	for _, we := range arrivalInfo {
		text, err := we.Text()
		check(err)

		arrivalDepartureList = append(arrivalDepartureList, text)
	}

	departure, err := wd.FindElements(selenium.ByCSSSelector, "span[class='uitk-text uitk-type-400 uitk-type-bold uitk-text-emphasis-theme'")
	departureList := []string{}

	for _, we := range departure {
		text, err := we.Text()
		check(err)

		departureList = append(departureList, text)
	}

	journeyDuration, err := wd.FindElements(selenium.ByCSSSelector, "div[class='uitk-text uitk-type-200 uitk-text-emphasis-theme'")
	journeyDurationList := []string{}

	for _, we := range journeyDuration {
		text, err := we.Text()
		check(err)

		journeyDurationList = append(journeyDurationList, text)
	}

	filghtsInfoList := FlightDetailsList{}
	for i, cost := range costList {
		flightInfo := FlightDetails{
			arrivalDeparture: arrivalDepartureList[i],
			departureTime:    departureList[i],
			cost:             cost,
			journeyDuration:  journeyDurationList[i],
		}
		filghtsInfoList = append(filghtsInfoList, flightInfo)
	}

	sort.SliceStable(filghtsInfoList, func(i, j int) bool {
		return filghtsInfoList[i].cost < filghtsInfoList[j].cost
	})

	writeToXlx(filghtsInfoList)

	time.Sleep(2 * time.Second)
}

func writeToXlx(filghtsInfoList FlightDetailsList) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Flights")
	check(err)

	row = sheet.AddRow()

	cell = row.AddCell()
	cell.Value = "Arrival Departure"

	cell = row.AddCell()
	cell.Value = "Journey Duration"

	cell = row.AddCell()
	cell.Value = "Cost"

	cell = row.AddCell()
	cell.Value = "Departure Time"

	for _, det := range filghtsInfoList {
		det.print()
		row = sheet.AddRow()

		cell = row.AddCell()
		cell.Value = det.arrivalDeparture

		cell = row.AddCell()
		cell.Value = det.journeyDuration

		cell = row.AddCell()
		cell.Value = strconv.Itoa(det.cost)

		cell = row.AddCell()
		cell.Value = det.departureTime

	}
	err = file.Save("flights.xlsx")
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
