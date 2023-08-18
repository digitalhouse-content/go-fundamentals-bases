package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2023, 07, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	then = then.Add(30 * time.Minute)
	p(then)

	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	p(now.Nanosecond())
	p(now.Location())
	p(now.Weekday())

	p("Then es antes que Now:", then.Before(now))
	p("Then es despues que Now:", then.After(now))
	p("Then es igual a Now:", then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p((diff.Hours() / 24) / 365)
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
}
