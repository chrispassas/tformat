package tformat

/*
--------------- + ------------ + ------------
Type            | Placeholder  |        Value
--------------- + ------------ + ------------
Year            | 2006         | 1609
Year            | 06           | 09
Month           | 01           | 09
Month           | 1            | 9
Month           | Jan          | Sep
Month           | January      | September
Day             | 02           | 12
Day             | 2            | 12
Week day        | Mon          | Sat
Week day        | Monday       | Saturday
Hours           | 03           | 07
Hours           | 3            | 7
Hours           | 15           | 19
Minutes         | 04           | 02
Minutes         | 4            | 2
Seconds         | 05           | 35
Seconds         | 5            | 35
AM or PM        | PM           | PM
Milliseconds     | .000         | .123
Microseconds    | .000000      | .123456
Nanoseconds     | .000000000   | .123456789
Timezone offset | -0700        | +0300
Timezone offset | -07:00       | +03:00
Timezone offset | Z0700        | +0300
Timezone offset | Z07:00       | +03:00
Timezone        | MST          | PDT
--------------- + ------------ + ------------
*/

const (
	// PGTIMESTAMP PostgreSQL timestamp format
	PGTIMESTAMP  = "2006-01-02 15:04:05.000000"
	YYYYMM       = "200601"
	YYYYMMDD     = "20060102"
	YYYYMMDDHH   = "2006010215"
	YYYYMMDDHHMI = "200601021504"
	// YYYY_MM_DD common date format YYYY-MM-DD
	YYYY_MM_DD  = "2006-01-02" //nolint
	YYYY        = "2006"
	YY          = "06"
	MM          = "01"
	Month       = "1"
	MonthAbbr   = "Jan"
	MonthName   = "January"
	WeekDay     = "Monday"
	WeekDayAbbr = "Mon"
	DD          = "02"
	Day         = "2"
	HH24        = "15"
	HH12        = "03"
	H12         = "3"
	MI          = "04"
	Minute      = "4"
	SS          = "05"
	Second      = "5"
	Milli       = ".000"
	Micro       = ".000000"
	Nano        = ".000000000"
	Timezone    = "MST"
	TZOffset1   = "-0700"
	TZOffset2   = "-07:00"
	TZOffset3   = "Z0700"
	TZOffset4   = "Z07:00"
)
