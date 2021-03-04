package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var year, month, day = time.Now().AddDate(0, -1, 0).Date()
var defaultMonthYear = fmt.Sprintf("%d%s%d", year, "/", month)
var monthYearLayout = "2006/1"
var helpFlag = flag.Bool("h", false, "Help")
var csvInputFlag = flag.String("input", "examplepermtsb.csv", "The csv input file to read")
var inputTypeFlag = flag.String("type", "open24", "Input format type.  Available options: open24, banktivity, revolut")
var monthYearFlag = flag.String("dates", defaultMonthYear, "Year/Month for open24 files, default is previous month")

func check(err error) {
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(1)
	}
	filename := *csvInputFlag
	// fmt.Printf("%s\n", *monthYearFlag)
	// outLayout := "2006/02"
	date, err := time.Parse(monthYearLayout, *monthYearFlag)
	check(err)
	// fmt.Printf("%s\n", date.Format(outLayout))
	if *inputTypeFlag == "open24" {
		open24CsvStuff(filename, date)
	} else if *inputTypeFlag == "banktivity" {
		banktivityCsvStuff(filename)
	} else if *inputTypeFlag == "revolut" {
		revolutCsvStuff(filename)
	}
}

func revolutCsvStuff(filename string) {
	inLayout := "Jan 2, 2006"
	outLayout := "01/02/2006"
	f, err := os.Open(filename)
	check(err)
	fileReader := bufio.NewReader(f)
	r := csv.NewReader(fileReader)
	r.Comma = ';'
	//read header line first and ignore it
	r.Read()
	fmt.Println("Date,Payee,FI Payee,Amount,CreditDebit,Category")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		date, err := time.Parse(inLayout, strings.TrimSpace(record[0]))
		check(err)
		formattedDate := date.Format(outLayout)
		creditDebit := "x"
		moneyOut, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(record[2]), ",", "", -1), 64)
		if err != nil {
			creditDebit = "credit"
		}
		moneyIn, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(record[3]), ",", "", -1), 64)
		if err != nil {
			creditDebit = "debit"
		}
		if creditDebit == "x" {
			moneyOut, err = strconv.ParseFloat(strings.Replace(strings.TrimSpace(record[4]), ",", "", -1), 64)
			if err != nil {
				creditDebit = "credit"
				moneyIn, err = strconv.ParseFloat(strings.Replace(strings.TrimSpace(record[5]), ",", "", -1), 64)
			}
		}
		amount := moneyIn - moneyOut
		fmt.Printf("\"")
		fmt.Printf(formattedDate)
		fmt.Printf("\",\"")
		fmt.Printf(strings.TrimSpace(record[1]))
		fmt.Printf("\",\"")
		fmt.Printf(strings.TrimSpace(record[8]))
		fmt.Printf("\",\"")
		fmt.Printf("%f", amount)
		fmt.Printf("\",\"")
		fmt.Printf(creditDebit)
		fmt.Printf("\",\"")
		fmt.Printf(strings.TrimSpace(record[7]))
		fmt.Printf("\"")
		fmt.Printf("\n")
	}
}

func banktivityCsvStuff(filename string) {
	inLayout := "1/2/06"
	outLayout := "01/02/2006"
	f, err := os.Open(filename)
	check(err)
	fileReader := bufio.NewReader(f)
	r := csv.NewReader(fileReader)
	//read header line first and ignore it
	r.Read()
	fmt.Println("Date,Payee,FI Payee,Amount,CreditDebit,Category")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		date, err := time.Parse(inLayout, record[2])
		check(err)
		formattedDate := date.Format(outLayout)
		creditDebit := string([]rune(record[5])[0])
		currency := "$"
		amount := float64(0)
		if creditDebit == "-" {
			currency = string([]rune(record[5])[1])
			creditDebit = "debit"
			amountString := string([]rune(record[5])[2:])
			amount, err = strconv.ParseFloat(strings.Replace(amountString, ",", "", -1), 64)
		} else {
			currency = creditDebit
			creditDebit = "credit"
			amountString := string([]rune(record[5])[1:])
			amount, err = strconv.ParseFloat(strings.Replace(amountString, ",", "", -1), 64)
		}
		// fmt.Printf("currency %s\n", currency)
		fmt.Printf("\"")
		fmt.Printf(formattedDate)
		fmt.Printf("\",\"")
		fmt.Printf(record[3])
		fmt.Printf("\",\"")
		fmt.Printf("%s - %s", record[6], currency)
		fmt.Printf("\",\"")
		fmt.Printf("%f", amount)
		fmt.Printf("\",\"")
		fmt.Printf(creditDebit)
		fmt.Printf("\",\"")
		fmt.Printf(record[4])
		fmt.Printf("\"")
		fmt.Printf("\n")
	}
}

func open24CsvStuff(filename string, monthYear time.Time) {
	inLayout := "02 Jan 06"
	outLayout := "01/02/2006"
	f, err := os.Open(filename)
	check(err)
	fileReader := bufio.NewReader(f)
	r := csv.NewReader(fileReader)
	// ignore field lengths (because open24 is a pita and doesn't output valid csv)
	r.FieldsPerRecord = -1
	//read header line first and ignore it
	r.Read()
	fmt.Println("Date,Payee,FI Payee,Amount,CreditDebit,Category")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if len(record) < 4 {
			continue // skip too short records
		}
		check(err)
		date, err := time.Parse(inLayout, record[0])
		check(err)
		// fmt.Printf("%s == %s\n", monthYear.Format(monthYearLayout), date.Format(monthYearLayout))
		if monthYear.Format(monthYearLayout) == date.Format(monthYearLayout) {
			formattedDate := date.Format(outLayout)
			creditDebit := "credit"
			columnIncrement := 0
			moneyOut := float64(0)
			_, err = time.Parse(inLayout, record[1])
			if err == nil {
				// if this is a credit card file then there will be another date in column 1
				columnIncrement = 1
			}
			moneyIn, err := strconv.ParseFloat(strings.Replace(record[2+columnIncrement], ",", "", -1), 64)
			fmt.Printf("moneyIn: %f\n", moneyIn)
			if err != nil || moneyIn == 0 {
				moneyOut, err = strconv.ParseFloat(strings.Replace(record[3+columnIncrement], ",", "", -1), 64)
				check(err)
				creditDebit = "debit"
			}
			amount := moneyIn - moneyOut
			fmt.Printf("\"")
			fmt.Printf(formattedDate)
			fmt.Printf("\",\"")
			fmt.Printf(record[1+columnIncrement])
			fmt.Printf("\",\"")
			fmt.Printf(record[1+columnIncrement])
			fmt.Printf("\",\"")
			fmt.Printf("%f", amount)
			fmt.Printf("\",\"")
			fmt.Printf(creditDebit)
			fmt.Printf("\",\"")
			fmt.Printf("\"")
			fmt.Printf("\n")
		}
	}
}
