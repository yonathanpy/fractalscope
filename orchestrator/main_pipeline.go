package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Event struct {
	Type string `json:"type"`
	Value int    `json:"value"`
}

func main() {
	fmt.Println("[fractalscope] pipeline start")

	event := Event{Type: "net", Value: 7}

	data, _ := json.Marshal(event)

	// send to python analyzer
	py := exec.Command("python3", "../python-analysis/anomaly_eval.py", string(data))
	outPy, _ := py.Output()

	// send to rust engine
	rs := exec.Command("../rust-core/target/release/correlation_engine", string(data))
	outRs, _ := rs.Output()

	fmt.Println("python:", string(outPy))
	fmt.Println("rust:", string(outRs))

	fmt.Println("[fractalscope] pipeline complete")
}
