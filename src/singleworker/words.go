package main

import (
  "os"
  "fmt"
  "log"
  "time"
  "strings"
)

const DataFile = "loremipsum.txt"

func clean(str string) string {
  s := strings.ToLower(str)
  var cleaned strings.Builder
  cleaned.Grow(len(s)) // Preallocate using one memory allocation
  for i := 0; i < len(s); i++ {
    b := s[i]
    if('a' <= b && b <= 'z') ||
      (b == ' ') ||
      (b == '\n') {

      cleaned.WriteByte(b)
    }
  }
  return cleaned.String()
}

func WordCount(text string) map[string]int {
  freqs := make(map[string]int)
  clean_text := clean(text)
  words := strings.Fields(clean_text)
  for _, word := range words {
    _, ok := freqs[word]
    if ok {
      freqs[word] += 1
    } else {
      freqs[word] = 1
    }
  }
  return freqs
}

func benchmark(text string, numRuns int) int64 {
  start := time.Now()
  for i := 0; i < numRuns; i++ {
    WordCount(text)
  }
  runtimeMillis := time.Since(start).Nanoseconds() / 1e6
  return runtimeMillis
}

func printResults(runtimeMillis int64, numRuns int) {
  fmt.Printf("amount of runs: %d\n", numRuns)
  fmt.Printf("total time: %d ms\n", runtimeMillis)
  average := float64(runtimeMillis) / float64(numRuns)
  fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
  input, err := os.ReadFile(DataFile)
  if err != nil {
    log.Fatalln(err)
  }
  data := string(input)

  fmt.Printf("%#v", WordCount(string(data)))
  numRuns := 100
  runtimeMillis := benchmark(string(data), numRuns)
  printResults(runtimeMillis, numRuns)
}
