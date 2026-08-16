package main

import (
    "fmt"
    "learn/go/day-15/config"
    "learn/go/day-15/counter"
    "learn/go/day-15/report"
)

func main() {
    cfg, err := config.Load("data/config.json")
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }
    if err := cfg.Validate(); err != nil {
        fmt.Println("Invalid config:", err)
        return
    }

    words, err := counter.CounterWords(cfg.InputFile)
    if err != nil {
        fmt.Println("Error counting words:", err)
        return
    }
    fmt.Println("Words:", words)

    summary := report.Build(cfg.ShopName, words)
    if err := report.Write(cfg.OutputFile, summary); err != nil {
        fmt.Println("Error saving summary:", err)
        return
    }
    fmt.Println("Summary saved to:", cfg.OutputFile)
}