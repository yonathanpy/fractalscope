package control

import "fmt"

func Monitor(score float64) {
    if score > 0.8 {
        fmt.Println("[alert] anomaly threshold exceeded")
    }
}
