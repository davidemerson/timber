package main

import (
    "fmt"
    "time"
)

// Workout represents a HIIT workout with intervals of work and rest.
type Workout struct {
    workInterval time.Duration
    restInterval time.Duration
    rounds       int
}

// Timer represents a timer for a HIIT workout.
type Timer struct {
    workout Workout
    start   time.Time
}

// NewTimer creates a new Timer for a given Workout.
func NewTimer(workout Workout) *Timer {
    return &Timer{
        workout: workout,
        start:   time.Now(),
    }
}

// Run starts the timer and prints the progress of the workout to the console.
func (t *Timer) Run() {
    for i := 1; i <= t.workout.rounds; i++ {
        fmt.Printf("Starting round %d\n", i)

        // Work interval.
        time.Sleep(t.workout.workInterval)
        fmt.Println("Work interval done, time for a rest!")

        // Rest interval.
        time.Sleep(t.workout.restInterval)
        fmt.Println("Rest interval done, time to work again!")
    }

    fmt.Println("Workout complete!")
}

func main() {
    workout := Workout{
        workInterval: 30 * time.Second,
        restInterval: 10 * time.Second,
        rounds:       4,
    }
    timer := NewTimer(workout)
    timer.Run()
}
