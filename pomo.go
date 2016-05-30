package main

import (
  "fmt"
  "os"
  "time"
  "strconv"
)

// pomo start
// pomo pause
// pomo stop
// pomo list tasks
// pomo start taskname

// currentlyRunning *Pomo
type Pomo struct {
  time *time.Duration
  runCount int
  status string
  timer *time.Timer
}

func NewPomo(totalTime *time.Duration) *Pomo {
  return &Pomo{
    time: totalTime,
    runCount: 0,
    status: "created",
    timer: nil,
  }
}

func (p *Pomo) createTimer() {
  p.timer = time.NewTimer(time.Second * (*p.time))
}

func (p *Pomo) startTimer() {
  // go func() {
    <- p.timer.C
    p.status = "expired"
    fmt.Println("Timer", p.status)
  // }()
}

func main() {
  cmd := os.Args[1]
  strAmt := os.Args[2]
  amt := 10 // default
  if strAmt != "" {
    amt, _ = strconv.Atoi(os.Args[2])
  }

  if cmd == "start" {
    a := time.Duration(amt)
    currentPomo := NewPomo(&a)
    currentPomo.createTimer()
    currentPomo.startTimer()
    // timer := time.NewTimer(time.Second * time.Duration(amt))
    // <- timer.C
  }

}
