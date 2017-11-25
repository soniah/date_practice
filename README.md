# Date Practice

Parsing dates is the programmer's equivalent of doing your taxes - it needs
doing, it's fidly and easy to get wrong, and you want it done quickly and
with as little pain as possible.

Previously I'd read [package time](https://golang.org/pkg/time), got confused
by the **reference time layouti** `Mon Jan 2 15:04:05 MST 2006` and fought my
way through various StackOverflow Questions. Or used
[imetakeule/fmtdate](https://github.com/metakeule/fmtdate).

Hence I wrote these exercises for myself, and I now find Golang's date parsing
very logical and easy (excepting some corner cases). Hopefully the exercises
will be of use to others.

Like many "small" projects it got larger than expected and I ended up rushing
to finish it, so it's a bit rough around the edges. I'm also interested in some
more questions - pull requests welcome for both.

# Instructions

```
git clone https://github.com/soniah/date_practice
go test
```

Edit `dates_test.go` until all tests pass.

Many of the dates were printed out using `time.Format()` with the provided
common layout strings. You could of course use these layouts to parse the
dates but you wouldn't learn much. Instead write your own layouts.

My answers are in the branch `answers` but no peeking!

Where the dates don't have timezones they in are in Hong Kong time.

## Useful Links

* [Martynas - Formatting date and time in Golang](https://medium.com/@Martynas/formatting-date-and-time-in-golang-5816112bf098)
