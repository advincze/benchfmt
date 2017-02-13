package main

import "fmt"

import "os"
import "bufio"
import "strings"
import "time"
import "strconv"
import "flag"

func main() {
	dur := flag.Duration("du", time.Second, "duration unit")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line) // Println will add back the final '\n'
		if strings.HasPrefix(line, "Benchmark") {
			fields := strings.Fields(line)
			if len(fields) < 4 || fields[3] != "ns/op" {
				fmt.Println(line)
				continue
			}
			nsPerOp, err := strconv.ParseInt(fields[2], 10, 64)
			if err != nil {
				fmt.Println(line)
				continue
			}
			opsPerDur := dur.Nanoseconds() / nsPerOp
			fmt.Printf("%s\t%d ops/%v\n", line, opsPerDur, dur)
			continue
		}
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
