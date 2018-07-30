package main

import "fmt"
import	"time"

var week time.Duration
func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Day())
	fmt.Println(t.Minute())
	fmt.Printf("%02d.%02d.%4d\n", t.Month(), t.Day(), t.Year()) 
	fmt.Println(t.Format("Jan 02 2006 15:04")) 
	
	week = 60 * 60 * 24 * 7 * 1e9
	t = time.Now().UTC()
	fmt.Println(t)

	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
}