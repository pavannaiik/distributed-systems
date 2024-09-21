package cos418_hw1_1

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"sync"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
func sumWorker(nums chan int, out chan int) {
	sum := 0
	for num := range nums {
		sum += num
	}
	// Output the sum to the out channel
	out <- sum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
func sum(num int, fileName string) int {
	// Open the file
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	// Read all integers from the file
	ints, err := readInts(file)
	checkError(err)

	// Create channels for workers
	nums := make(chan int, len(ints)) // Buffered channel for integers
	out := make(chan int, num)        // Buffered channel for sum from each worker
	var wg sync.WaitGroup             // WaitGroup to wait for all workers to finish

	// Launch workers
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sumWorker(nums, out)
		}()
	}

	// Feed the numbers into the nums channel
	for _, num := range ints {
		nums <- num
	}
	close(nums) // Close nums channel after all numbers have been sent

	// Wait for workers to finish summing
	go func() {
		wg.Wait()
		close(out) // Close the out channel after all workers are done
	}()

	// Collect the results from the out channel
	totalSum := 0
	for partialSum := range out {
		totalSum += partialSum
	}

	return totalSum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
