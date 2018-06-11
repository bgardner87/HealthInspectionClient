package main

import (
	"fmt"

	"github.com/CincyGolangMeetup/HealthInspection/Inspection/Query"
)

func main() {
	q := query.NewQuery()
	q = q.QueryWithBusinessName("HABANERO")
	// q = q.QueryWithAddress("ADDRESS")
	// q = q.QueryWithCity("CITY")
	// q = q.QueryWithState("ST")
	//q.Print()
	results, err := q.Execute()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for _, result := range *results {
		fmt.Println("")
		result.Print()
	}
}
