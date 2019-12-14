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

var helpFlag = flag.Bool("h", false, "Help")
var csvInputFlag = flag.String("input", "examplepermtsb.csv", "The csv input file to read")
var inputTypeFlag = flag.String("type", "open24", "Input format type.  Available options: open24, banktivity")

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
	if *inputTypeFlag == "open24" {
		open24CsvStuff(filename)
	} else if *inputTypeFlag == "banktivity" {
		banktivityCsvStuff(filename)
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

func open24CsvStuff(filename string) {
	inLayout := "02 Jan 06"
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
		date, err := time.Parse(inLayout, record[0])
		check(err)
		formattedDate := date.Format(outLayout)
		creditDebit := "debit"
		moneyOut, err := strconv.ParseFloat(strings.Replace(record[3], ",", "", -1), 64)
		if err != nil {
			creditDebit = "credit"
		}
		moneyIn, err := strconv.ParseFloat(strings.Replace(record[2], ",", "", -1), 64)
		amount := moneyIn - moneyOut
		fmt.Printf("\"")
		fmt.Printf(formattedDate)
		fmt.Printf("\",\"")
		fmt.Printf(record[1])
		fmt.Printf("\",\"")
		fmt.Printf(record[1])
		fmt.Printf("\",\"")
		fmt.Printf("%f", amount)
		fmt.Printf("\",\"")
		fmt.Printf(creditDebit)
		fmt.Printf("\",\"")
		fmt.Printf("\"")
		fmt.Printf("\n")
	}
}
