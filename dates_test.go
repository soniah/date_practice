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
	parseFunction func(string) (time.Time, error)

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
		"",
		1511263707,
		// YOU might also need to write the following
		nil,
		// YOU might also need to write the following
		nil,
	},
	{
		"bandicoot",
		"Tuesday, 21 November 2017 7:28:27 PM GMT+08:00",
		"",
		1511263707,
		nil,
		nil,
	},
	{
		"dingo",
		"Fri Oct 26 11:41:59 HKT 1979",
		"",
		309757319,
		nil,
		nil,
	},
	{
		"quokka",
		"Tue Jan 27 07:30:41 +0800 1970",
		"",
		2244641,
		nil,
		nil,
	},
	{
		"echidna",
		"Monday, 21-Apr-75 11:02:18 HKST",
		"",
		167277738,
		nil,
		nil,
	},
	{
		"emu",
		"Thu, 21 Sep 2000 05:45:12 HKT",
		"",
		969486312,
		nil,
		nil,
	},
	{
		"quoll",
		"Sun, 14 Sep 1997 20:18:04 +0800",
		"",
		874239484,
		nil,
		nil,
	},
	{
		"wallaby",
		"1991-11-13T00:08:18+08:00",
		"",
		689962098,
		nil,
		nil,
	},
	{
		"kookaburra",
		"1973-11-10T23:42:42+08:00",
		"",
		121794162,
		nil,
		nil,
	},
	{
		"platypus",
		"Fri May  7 01:04:53 1982",
		"",
		389552693,
		nil,
		nil,
	},
	{
		"bilby",
		"21 Apr 87 20:11 HKT",
		"",
		546005494,
		nil,
		nil,
	},
	{
		"cassowary",
		"08 Jan 70 14:59 +0800",
		"",
		629954,
		nil,
		nil,
	},
	{
		"numbat",
		"2:54PM",
		"",
		28104869,
		nil,
		nil,
	},
	// if you find the following difficult (I did), see for ideas:
	// https://stackoverflow.com/questions/47471071/parse-dates-with-ordinal-date-fields/47475260#47475260
	{
		"wombat",
		"Sunday 23rd January 2033 04:38:25 AM",
		"",
		1990067905,
		nil,
		nil,
	},
	{
		"kangaroo",
		"Tuesday 7th November 2017 03:18:25 PM",
		"",
		1510067905,
		nil,
		nil,
	},
}

func TestExercises(t *testing.T) {
	for _, row := range testsFoo {

		var parsedTime time.Time
		var err error

		if row.parseFunction == nil {
			parsedTime, err = time.Parse(row.answer, row.question)
		} else {
			parsedTime, err = row.parseFunction(row.question)
		}

		if err != nil {
			t.Errorf("\n%s:\n  question: %s\n  answer:   %s\nerror:\n  %v\n\n", row.animal, row.question, row.answer, err)
		}

		var equal bool
		unixTime := time.Unix(row.unixTime, 0)

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

