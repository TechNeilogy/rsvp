# rsvp

## Reservior Sampling Utility

A command line utility for reservoir sampling from stdin.

Submitted for entry in the development programming contest of Fall, 2022: "Pick three random entries from a file."

### Goals

I had two goals for this project:

1. Learn to create a Linux-style command line utility using Go, including:
   1. Use of command line parameters.
   2. Use of standard IO.
   3. Creating **man** pages.
2. Experiment with the reservoir sampling algorithm.

### Building

This project assumes a standard Go setup with some recent version of Go.  That step is beyond this documentation; see the Go language home page, or the appropriate page of your cloud provider.  

Once properly configured, the project can be built by running the "go build" command in the project root directory:

```shell
go build .
```

### Reservoir Sampling Algorithm

#### Summary
Chooses a random sample, without replacement, of **k** items from a population of unknown size, in a single pass, without storing previously seen unselected values.  It is ideal for use where the population size is (or could be) large, and it is undesirable to store the entire population in memory.

#### More Information
https://en.wikipedia.org/wiki/Reservoir_sampling

### Usage

#### Summary
Chooses **k** random lines from **stdin**, or from a file, stopping on **EOF** (**ctrl-d**) or some maximum number of lines.

#### Examples
(Assumes the path to rsvp is not in PATH.  If it is, you can use "rsvp".)
```shell
# Display 5 random lines from test100.txt
cat test100.txt | ./rsvp -k=5

# Take input from a file (pipe is ignored)
./rsvp test100.txt

# Select 5 random lines from test100.txt
# sort them, and write them to samples.txt
cat test100.txt | rsvp -k=5 | sort >samples.txt

# It works with anything that writes to stdout:
ls -l | rsvp -ml=100 -sk=1 -s

# Print help and version.
rsvp --help
rsvp -v

# Tries to handle edge conditions gracefully.
cat missingfile.txt | ./rsvp -k=5
```
#### Parameters
- **--help** - Print brief help text and exit.  (Auto-generated by the Go flags parser.)


- **-k** - Number of samples to return.  If this number is larger than the number of lines from **stdin**, the program will return a number of samples equal to the number of lines.  Defaults to 3.  (Values less than 0 revert to 0.)


- **-ml** - Maximum number of lines to process, *not including the number of lines skipped by* ***-sk*** (below).  The program will stop and return a sample after processing this many lines, regardless of the value of **-k** or the number of lines in **stdin**.  Facilitates sampling from very large or infinite streams.  Defaults to unlimited.  (Values less than 0 revert to unlimited.)


- **-s** - Silent mode.  If the program exits in error, no error message will be printed to **stderr**.


- **-sk** - Initial number of lines to skip.  This many lines will be skipped regardless of the values of **-k** or **-ml**.  Defaults to none.  (Values less than 0 revert to none.)


- **-v** - Print version and exit.
#### Return
Results are printed to **stout**, errors are printed to **stderr**.  Exits with code 0 for success, code 1 for any error.

### Files
- go.mod - Go module file.
- README.md - This file.
- rsvp.go - Source code
- rsvp.1 - Linux "man" page.
- test100.txt - A file of 100 indexed lines for testing and experimentation.

### Notes
- Name stands for ReSerVoir Program.


- Ideal for installation as a general utility.  Compile the project using "go build .", and copy the resulting executable (rsvp) to some directory in your PATH.  
  - Optionally, copy the man page file (rsvp.1) to the appropriate directory for your distro.  You can preview it locally by entering:
    ```shell
    man ./rsvp.1
    ```


- Credits: neil.carrier@healthcatalyst, Fall 2022