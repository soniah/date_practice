package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var epoch time.Time

var testsFoo = []struct {
	// an Australian animal name, to help navigate test output
	animal string

	// the "question" - the time string to be parsed. Times without a
	// timezone are Hong Kong times
	question string

	// your "answer" - the layout string used by time.Parse().
	// If you have a parseFuncion "answer" will instead be passed
	// into that function
	answer string

	// the unixTime value of the "question", used by the tests to check
	// "question" and "answer" equality
	unixTime int64

	// some questions can't be parsed only using a layout string; for
	// these write a parsing function
	parseFunction func(answer string, question string) (time.Time, error)

	// some questions can't be parsed exactly; for these write a function
	// to test equality
	equalityFunction func(unixTime, answerTime time.Time) bool
}{

	{
		"bandicoot",
		"Tuesday, 21 November 2017 7:28:27 PM GMT+08:00",

		// YOU need to write the following
		"",

		1511263707,

		// YOU _might_ need to write the following
		nil,

		// YOU _might_ need to write the following
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
		"koala",
		"Tuesday, 21 November 2017 7:28:27 PM HKT",
		"",
		1511263707,
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
		"dingo",
		"Fri Oct 26 11:41:59 HKT 1979",
		"",
		309757319,
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
		"kookaburra",
		"1973-11-10T23:42:42+08:00",
		"",
		121794162,
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

	// if you find the following difficult (I did) see:
	// https://stackoverflow.com/questions/47471071/parse-dates-with-ordinal-date-fields/47475260#47475260
	{
		"wombat",
		"Sunday 23rd January 2033 04:38:25 AM",
		"",
		1990039105,
		nil,
		nil,
	},

	{
		"kangaroo",
		"Tuesday 7th November 2017 03:18:25 PM",
		"",
		1510039105,
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
			parsedTime, err = row.parseFunction(row.answer, row.question)
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

