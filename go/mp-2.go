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

type HotelDetails struct {
	hotelname     string
	negibhborhood string
	cost          int
}

type HotelDetailsList []HotelDetails

func (det HotelDetails) print() {
	fmt.Print("Hotel Name: ", det.hotelname)
	fmt.Print(" Negibhborhood: ", det.negibhborhood)
	fmt.Print(" Cost: ", det.cost)
}

func main() {

	destination := "bengaluru"
	checkinDate := "2021-09-30"
	checkoutDate := "2021-10-01"
	url := "https://www.expedia.co.in/Hotel-Search?adults=2&d1=" + checkinDate + "&d2=" + checkoutDate + "&destination=" + destination + "&directFlights=false&endDate=" + checkoutDate + "&localDateFormat=dd%2FMM%2Fyyyy&partialStay=false&semdtl=&sort=PRICE_LOW_TO_HIGH&startDate=" + checkinDate + "0&theme=&useRewards=false&userIntent="

	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	check(err)

	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	check(err)
	defer wd.Quit()

	println(url)
	if err := wd.Get(url); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	costInfo, err := wd.FindElements(selenium.ByCSSSelector, "span[data-stid='price-lockup-text'")
	check(err)

	costList := []int{}
	for _, we := range costInfo {
		text, err := we.Text()
		check(err)

		text = strings.Trim(text, " ")
		costStr := strings.Split(text, "Rs")[1]
		costStr = strings.ReplaceAll(costStr, ",", "")
		cost, err := strconv.Atoi(costStr)

		costList = append(costList, cost)
	}
	hotelname, err := wd.FindElements(selenium.ByCSSSelector, "h3[data-stid='content-hotel-title'")
	hotelnameList := []string{}

	for _, we := range hotelname {
		text, err := we.Text()
		check(err)

		hotelnameList = append(hotelnameList, text)
	}

	negibhborhood, err := wd.FindElements(selenium.ByCSSSelector, "div[data-test-id='content-hotel-neighborhood'")
	negibhborhoodList := []string{}

	for _, we := range negibhborhood {
		text, err := we.Text()
		check(err)

		negibhborhoodList = append(negibhborhoodList, text)
	}

	hotelsInfoList := HotelDetailsList{}
	for i, cost := range costList {
		hotelInfo := HotelDetails{
			hotelname:     hotelnameList[i],
			negibhborhood: negibhborhoodList[i],
			cost:          cost,
		}
		hotelsInfoList = append(hotelsInfoList, hotelInfo)
	}

	sort.SliceStable(hotelsInfoList, func(i, j int) bool {
		return hotelsInfoList[i].cost < hotelsInfoList[j].cost
	})

	writeToXlx(hotelsInfoList)

	time.Sleep(2 * time.Second)
}

func writeToXlx(hotelsInfoList HotelDetailsList) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Hotels")
	check(err)

	row = sheet.AddRow()

	cell = row.AddCell()
	cell.Value = "Hotel Name"

	cell = row.AddCell()
	cell.Value = "Neighbourhood"

	cell = row.AddCell()
	cell.Value = "Cost"

	for _, det := range hotelsInfoList {
		det.print()
		row = sheet.AddRow()

		cell = row.AddCell()
		cell.Value = det.hotelname

		cell = row.AddCell()
		cell.Value = det.negibhborhood

		cell = row.AddCell()
		cell.Value = strconv.Itoa(det.cost)

	}
	err = file.Save("hotels.xlsx")
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
