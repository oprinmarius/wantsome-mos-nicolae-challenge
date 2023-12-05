You have a data processing system that reads data from multiple sources concurrently, processes it, and then aggregates the results. Each data source provides a stream of integers. Your task is to design a solution using Go that concurrently processes data from these sources, sums the integers, and aggregates the results into a final total.

```
type DataSource struct {
	Name    string
	Channel chan int
}

func processDataSource(dataSource DataSource, resultChannel chan int) {
}

func aggregateResults(resultChannels []chan int) {
}

func main() {
}
```
