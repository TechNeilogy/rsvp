package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// Sample lines from a stream using the reservoir sampling algorithm.
//
// input - Input reader.
// k - Sample count.
// maxLines - Maximum number of lines to process (not including skipLines).
// skipLines - Initial number of lines to skip.
//
// Returns an array of samples, and an error (which may be nil).
func reservoirSampler(
	input io.Reader,
	k int,
	maxLines int,
	skipLines int,
) ([]string, error) {

	// Normalize k and maxLines.

	if maxLines <= -1 {
		maxLines = -1
	} else {
		if maxLines < k {
			k = maxLines
		}
	}

	if k <= 0 {
		return make([]string, 0), nil
	}

	scanner := bufio.NewScanner(input)

	// Skip initial lines.

	if skipLines > 0 {
		for ; skipLines > 0; skipLines-- {
			scan := scanner.Scan()
			if !scan {
				return make([]string, 0), nil
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	samples := make([]string, k)

	// Copy initial batch.

	i := 0
	for ; i < k; i++ {
		scan := scanner.Scan()
		if !scan {
			break
		}
		samples[i] = scanner.Text()
	}
	k = i
	linesTotal := k

	// Reservoir algorithm.

	if maxLines < 0 || linesTotal < maxLines {

		fk := float64(k)

		w := math.Exp(math.Log(rand.Float64()) / fk)

		for true {

			skip := math.Floor(math.Log(rand.Float64())/math.Log(1-w)) + 1

			halt := false

			for ; skip > 0; skip-- {
				if maxLines >= 0 && linesTotal >= maxLines {
					halt = true
					break
				}
				scan := scanner.Scan()
				linesTotal++
				if !scan {
					halt = true
					break
				}
			}
			if halt {
				break
			}

			samples[rand.Intn(k)] = scanner.Text()

			w = w * math.Exp(math.Log(rand.Float64())/fk)
		}
	}

	return samples[0:k], nil
}

func main() {

	k := flag.Int("k", 3, "number of samples")
	ml := flag.Int("ml", -1, "maximum number of lines to read")
	s := flag.Bool("s", false, "silent mode (no error message on failure)")
	sk := flag.Int("sk", -1, "lines to skip initially")
	v := flag.Bool("v", false, "print version and exit")

	flag.Parse()

	if *v {
		_, _ = fmt.Fprintln(os.Stdout, "1.0.0")
		os.Exit(0)
	}

	samples, err := reservoirSampler(
		os.Stdin,
		*k,
		*ml,
		*sk,
	)

	if err != nil {
		if !*s {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}

	for _, sample := range samples {
		_, err := fmt.Fprintln(os.Stdout, sample)
		if err != nil {
			if !*s {
				_, _ = fmt.Fprintln(os.Stderr, err)
			}
			os.Exit(1)
		}
	}

	os.Exit(0)
}
