package main

import (
  "os"
  "fmt"
  "log"
  "sync"
  "time"
  "strings"
)

const DataFile = "loremipsum.txt"

func clean(str string) string {
  s := strings.ToLower(str)
  var cleaned strings.Builder
  cleaned.Grow(len(s))
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

func WordCount(text string) (freq map[string]int) {
  freq = make(map[string]int)
  clean_text := clean(text)
  words := strings.Fields(clean_text)

  const go_size = 2048
  str_len := len(words)
  subgo_channel := make(chan map[string]int, (str_len / go_size) + 1)

  var wg sync.WaitGroup

  for i, j := 0, go_size; i < str_len; i, j = j, (j + go_size) {
    if j > str_len {
      j = str_len
    }
    wg.Add(1)
    go func(x, y int) {
      subgo_map := make(map[string]int)
      for z := x; z < y; z++ {
        subgo_map[words[z]]++
      }
      subgo_channel <- subgo_map
      wg.Done()
    }(i, j)
  }

  wg.Wait()
  close(subgo_channel)
  for sub := range subgo_channel {
    for word, val := range sub {
      freq[word] += val
    }
  }
  return
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
