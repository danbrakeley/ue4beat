package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	// Version is the app version
	Version = "0.0.4"
)

func printVersion() {
	fmt.Printf("ue4beat v%s\n", Version)
}

func printError(err error) {
	fmt.Printf("ue4beat: %v\n", err)
}

func printUsage() {
	fmt.Printf("\n")
	fmt.Printf("Usage: ue4beat [-f <name> <value> [-f <name> <value> [...]]]\n")
	fmt.Printf("\n")
	fmt.Printf("Command Line Parameters:\n")
	fmt.Printf("  -f, --field <name> <value>   Add name:value field to each output line\n")
	fmt.Printf("  -h, --help                   Print this information\n")
	fmt.Printf("  -v, --version                Print just the app name and version\n")
}

func main() {
	// parse command line parameters
	extraFields := make(map[string]interface{})
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-v", "version", "-version", "--version":
			printVersion()
			os.Exit(0)
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		case "-f", "--field":
			if i+2 >= len(os.Args) {
				printError(fmt.Errorf("-f requires both <name> and <value>"))
				os.Exit(-1)
			}
			i++
			name := os.Args[i]
			if !strings.HasPrefix(name, "fields.") {
				name = fmt.Sprintf("fields.%s", name)
			}
			i++
			value := os.Args[i]
			if number, err := strconv.ParseFloat(value, 64); err == nil {
				extraFields[name] = number
			} else {
				extraFields[name] = value
			}
		default:
			printError(fmt.Errorf("unrecognized parameter: %s", os.Args[i]))
			os.Exit(-1)
		}
	}

	curLine := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		curLine++

		u := ParseLine(s.Text())
		fields := make(map[string]interface{})

		// copy fields from command line
		for k, v := range extraFields {
			fields[k] = v
		}

		// copy fields from UE4Line
		if u.Timestamp != nil {
			fields["@timestamp"] = *u.Timestamp
		}
		if len(u.Category) > 0 {
			fields["fields.category"] = u.Category
		}
		if u.Frame >= 0 {
			fields["fields.frame"] = u.Frame
		}
		fields["fields.level"] = u.Level
		fields["message"] = u.Message
		fields["fields.log_line"] = curLine

		// marshal the results
		out, err := json.Marshal(fields)
		if err != nil {
			fmt.Fprintf(os.Stdout, "{\"level\":\"error\",\"message\":\"error marshalling to json\",\"fields.error\":\"%v\",\"fields.line\":%d}\n", err, curLine)
			continue
		}

		fmt.Fprintf(os.Stdout, "%s\n", string(out))
	}
}
