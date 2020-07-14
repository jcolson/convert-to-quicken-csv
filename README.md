# open24tomint

## Build status

![Go](https://github.com/jcolson/permanenttsb-csv-to-mint-csv/workflows/Go/badge.svg)


## Quicken is a total PITA!

If your bank doesn't pay Quicken/Intuit, they basically try to do everything in their power to disable you from importing transactions.

QIF files can't be imported into an account - as you'll get this error message from Quicken:
```
"Quicken can only import qif files into empty documents."
```

OFX files can't be imported into an account (unless your bank pays quicken) - as you'll get this error message:
```
"Quicken is unable to update this account because Web Connect support for your financial institution has been either temporarily, or permanently discontinued [CC-885]."
```

## Application usage

This app is for converting certain csv formats to something that quicken will 'allow' to import (the Mint CSV format).

Formats that can be converted
* Permanent TSB
* Banktivity (Mac financial mgmt application, like Quicken)
* Revolut exports

## how to build:
```
go build
```

## Help on running

```
Usage of ./permanenttsb-csv-to-mint-csv:
  -h	Help
  -input string
    	The csv input file to read (default "examplepermtsb.csv")
  -type string
    	Input format type.  Available options: open24, banktivity, revolut (default "open24")
```

