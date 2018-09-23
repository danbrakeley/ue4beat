package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		u := ParseLine(line)
		out, err := json.Marshal(u)
		if err != nil {
			fmt.Fprintf(os.Stdout, "{\"level\":\"error\",\"message\":\"error marshalling to json: %v\"}\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", string(out))
		}
	}
}
