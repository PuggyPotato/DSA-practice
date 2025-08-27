package main

import (
    "fmt"
    "sort"
    "time"
)

type Weather struct {
    Id          int
    RecordDate  string
    Temperature int
}

func risingTemperature(weather []Weather) []int {
    // Sort by RecordDate
    sort.Slice(weather, func(i, j int) bool {
        return weather[i].RecordDate < weather[j].RecordDate
    })

    result := []int{}
    layout := "2006-01-02"

    for i := 1; i < len(weather); i++ {
        prevDate, _ := time.Parse(layout, weather[i-1].RecordDate)
        currDate, _ := time.Parse(layout, weather[i].RecordDate)

        // Check if current day is exactly yesterday + 1 day
        if currDate.Sub(prevDate).Hours() == 24 && weather[i].Temperature > weather[i-1].Temperature {
            result = append(result, weather[i].Id)
        }
    }
    return result
}

func main() {
    data := []Weather{
        {1, "2015-01-01", 10},
        {2, "2015-01-02", 25},
        {3, "2015-01-03", 20},
        {4, "2015-01-04", 30},
    }

    ans := risingTemperature(data)
    fmt.Println(ans) // Output: [2 4]
}
