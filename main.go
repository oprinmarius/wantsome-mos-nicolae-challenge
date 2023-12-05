package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DataSource struct {
	Name    string
	Channel chan int
}

func processDataSource(dataSource DataSource, resultChannel chan int) {
	sum := 0
	for value := range dataSource.Channel {
		sum += value
	}
	resultChannel <- sum
}

func aggregateResults(resultChannels []chan int) int {
	totalSum := 0
	for _, ch := range resultChannels {
		totalSum += <-ch
	}
	return totalSum
}

func main() {
	numDataSources := 5

	rand.Seed(time.Now().UnixNano())

	// Dynamically create data sources and result channels
	dataSources := make([]DataSource, numDataSources)
	resultChannels := make([]chan int, numDataSources)
	for i := range dataSources {
		dataSources[i] = DataSource{
			Name:    fmt.Sprintf("Source%d", i+1),
			Channel: make(chan int),
		}
		resultChannels[i] = make(chan int)
		go processDataSource(dataSources[i], resultChannels[i])
	}

	// Simulate data input for each data source
	for _, ds := range dataSources {
		go func(dataSource DataSource) {
			for j := 0; j < 3; j++ { // Simulate 3 data inputs per source
				dataSource.Channel <- rand.Intn(100) // Random integer between 0-99
			}
			close(dataSource.Channel)
		}(ds)
	}

	// Aggregate and display the results
	fmt.Printf("Total Sum: %d\n", aggregateResults(resultChannels))
}
