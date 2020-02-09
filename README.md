# open24tomint

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

## This app

This app is primarily used to convert csv files from the Permanent TSB (open24.ie) to the Mint CSV format,
which is importable into Quicken.

It also allows the conversion of Banktivity (Mac financial mgmt application, like Quicken) CSV exports into mint csv format for importing into quicken.

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
    	Input format type.  Available options: open24, banktivity (default "open24")
```

