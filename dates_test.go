package main

/*
Fill in the questions and run 'go test'.

https://golang.org/src/time/format.go

Mon Jan 2 15:04:05 MST 2006    01/02 03:04:05PM '06 -0700

Hint: when parsing difficult dates, you can build up the layout chunk by chunk
- time.Parse() will print out the remaining unmatched text.
*/

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

var zero time.Time

var testsFoo = []struct {
	// an Australian animal name, to help navigate test output
	animal string

	// the "question" - the time string to be parsed. Times without a
	// timezone are Hong Kong times
	question string

	// your "answer" - the layout string to be used by time.Parse();
	// ignored if parseFunction is not nil
	answer string

	// the unixTime value of the "question", used by the tests to check
	// "question" and "answer" equality
	unixTime int64

	// some questions can't be parsed using only a layout string; for
	// these write a parsing function
	parseFunction func(string, string, *time.Location) (time.Time, error)

	// some questions can't be parsed exactly; for these write an equality
	// function to test if unixTime is "equal" to your answer
	//
	// you may only need to write parseFunction, only equalityFunction, or
	// both functions
	equalityFunction func(unixTime, answer time.Time) bool
}{
	{
		"koala",
		"Tuesday, 21 November 2017 7:28:27 PM HKT",
		// YOU write the following
		"Monday, 02 January 2006 3:04:05 PM MST", // REPLACE_EMPTY_STRING
		1511263707,
		// YOU might also need to write the following
		nil,
		// YOU might also need to write the following
		nil,
	},
	{
		"bandicoot",
		"Tuesday, 21 November 2017 7:28:27 PM GMT+08:00",
		"Monday, 02 January 2006 3:04:05 PM GMT-07:00", // REPLACE_EMPTY_STRING
		1511263707,
		nil,
		nil,
	},
	{
		"dingo",
		"Fri Oct 26 11:41:59 HKT 1979",
		"Mon Jan 02 15:04:05 MST 2006", // REPLACE_EMPTY_STRING
		309757319,
		nil,
		nil,
	},
	{
		"quokka",
		"Tue Jan 27 07:30:41 +0800 1970",
		"Mon Jan 02 15:04:05 -0700 2006", // REPLACE_EMPTY_STRING
		2244641,
		nil,
		nil,
	},
	{
		"echidna",
		"Monday, 21-Apr-75 11:02:18 HKST",
		"Monday, 02-Jan-06 15:04:05 MST", // REPLACE_EMPTY_STRING
		167277738,
		nil,
		nil,
	},
	{
		"emu",
		"Thu, 21 Sep 2000 05:45:12 HKT",
		"Mon, 02 Jan 2006 15:04:05 MST", // REPLACE_EMPTY_STRING
		969486312,
		nil,
		nil,
	},
	{
		"quoll",
		"Sun, 14 Sep 1997 20:18:04 +0800",
		"Mon, 02 Jan 2006 15:04:05 -0700", // REPLACE_EMPTY_STRING
		874239484,
		nil,
		nil,
	},
	{
		"wallaby",
		"1991-11-13T00:08:18+08:00",
		"2006-01-02T15:04:05-07:00", // REPLACE_EMPTY_STRING
		689962098,
		nil,
		nil,
	},
	{
		"kookaburra",
		"1973-11-10T23:42:42+08:00",
		"2006-01-02T15:04:05-07:00", // REPLACE_EMPTY_STRING
		121794162,
		nil,
		nil,
	},
	{
		"platypus",
		"Fri May  7 01:04:53 1982",
		"Mon Jan _2 15:04:05 2006", // REPLACE_EMPTY_STRING
		389552693,
		nil,
		nil,
	},
	{
		"bilby",
		"21 Apr 87 20:11 HKT",
		"02 Jan 06 15:04 MST", // REPLACE_EMPTY_STRING
		546005494,
		nil,
		minuteEquality, // REPLACE_NIL
	},
	{
		"cassowary",
		"08 Jan 70 14:59 +0800",
		"02 Jan 06 15:04 -0700", // REPLACE_EMPTY_STRING
		629954,
		nil,
		minuteEquality, // REPLACE_NIL
	},
	{
		"numbat",
		"2:54PM",
		"3:04PM", // REPLACE_EMPTY_STRING
		28104869,
		nil,
		kitchenEquality, // REPLACE_NIL
	},
	// if you find the following difficult (I did), see for ideas:
	// https://stackoverflow.com/questions/47471071/parse-dates-with-ordinal-date-fields/47475260#47475260
	{
		"wombat",
		"Sunday 23rd January 2033 04:38:25 AM",
		"Monday 02 January 2006 15:04:05 PM",
		1990067905,
		parseOrdinalDate, // REPLACE_NIL
		nil,
	},
	{
		"kangaroo",
		"Tuesday 7th November 2017 03:18:25 PM",
		"Monday 02 January 2006 15:04:05 PM",
		1510067905,
		parseOrdinalDate, // REPLACE_NIL
		nil,
	},
}

func TestExercises(t *testing.T) {

	location, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range testsFoo {

		var parsedTime time.Time
		var err error

		if row.parseFunction == nil {
			parsedTime, err = time.ParseInLocation(row.answer, row.question, location)
		} else {
			parsedTime, err = row.parseFunction(row.answer, row.question, location)
		}

		if err != nil {
			t.Errorf("\n%s:\n  question: %s\n  answer:   %s\nerror:\n  %v\n\n", row.animal, row.question, row.answer, err)
		}

		var equal bool
		unixTime := time.Unix(row.unixTime, 0).In(location)

		if row.equalityFunction == nil {
			equal = unixTime.Equal(parsedTime)
		} else {
			equal = row.equalityFunction(unixTime, parsedTime)
		}
		if !equal {
			t.Errorf("\n%s: unix timestamps don't match:\n  you   : %+v\n  golang: %+v\n\n", row.animal, parsedTime, unixTime)
		}

	}
}

// ----------- TRUNCATE

// test equality to the minute by rounding down; can't use
// duration.Round() as it rounds to the *nearest* minute
func minuteEquality(unixTime, answer time.Time) bool {
	seconds := unixTime.Second()
	duration := time.Duration(seconds) * time.Second
	unixTime = unixTime.Add(-duration)
	return unixTime.Equal(answer)
}

// test equality for "kitchen times" like "3:04PM", by
// comparing the hour and minute values of the inputs
func kitchenEquality(unixTime, answer time.Time) bool {
	uhours := unixTime.Hour()
	uminutes := unixTime.Minute()
	ahours := answer.Hour()
	aminutes := answer.Minute()

	if (uhours == ahours) && (uminutes == aminutes) {
		return true
	}
	return false
}

var ordinals = map[string]string{
	"01st": "01", "02nd": "02", "03rd": "03", "04th": "04", "05th": "05",
	"06th": "06", "07th": "07", "08th": "08", "09th": "09", "10th": "10",
	"11th": "11", "12th": "12", "13th": "13", "14th": "14", "15th": "15",
	"16th": "16", "17th": "17", "18th": "18", "19th": "19", "20th": "20",
	"21st": "21", "22nd": "22", "23rd": "23", "24th": "24", "25th": "25",
	"26th": "26", "27th": "27", "28th": "28", "29th": "29", "30th": "30",
	"31st": "31",
	"1st":  "01", "2nd": "02", "3rd": "03", "4th": "04", "5th": "05",
	"6th": "06", "7th": "07", "8th": "08", "9th": "09",
}

// convert dates with days of month like "2nd" to "02" or "15th" to "15"
// assume days of month are delimited by spaces eg "Jan 1st 2017" not
// "Jan1st2017" - a reasonable assumption.
func ordinalToCardinal(ordinalDate string) (string, error) {
	var found bool

	splits := strings.Split(ordinalDate, " ")
	if len(splits) == 0 {
		return "", fmt.Errorf("no spaces in date %s", ordinalDate)
	}

	for i, split := range splits {
		if cardinal, ok := ordinals[split]; ok {
			found = true
			splits[i] = cardinal
			break
		}
	}

	if !found {
		return "", fmt.Errorf("ordinal day not found in %s", ordinalDate)
	}

	return strings.Join(splits, " "), nil
}

func parseOrdinalDate(layout, question string, location *time.Location) (time.Time, error) {
	cardinalQuestion, err := ordinalToCardinal(question)
	if err != nil {
		return zero, err
	}

	result, err := time.Parse(layout, cardinalQuestion)
	if err != nil {
		return zero, err
	}
	// TODO vs platypus - why In() vs ParseInLocation()?
	return result.In(location), nil
}
