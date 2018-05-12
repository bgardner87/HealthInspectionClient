package main

import (
	"github.com/CincyGolangMeetup/HealthInspection"
)

func main() {
	q := healthinspection.NewQuery()
	q = q.QueryWithBusinessName("NAME")
	q = q.QueryWithAddress("ADDRESS")
	q = q.QueryWithCity("CITY")
	q = q.QueryWithState("ST")
	q.Print()
	for _, r := range *q.Execute() {
		r.Print()
	}
}
