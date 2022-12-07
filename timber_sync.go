package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gorilla/websocket"
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
    conn    *websocket.Conn
}

// NewTimer creates a new Timer for a given Workout.
func NewTimer(workout Workout, conn *websocket.Conn) *Timer {
    return &Timer{
        workout: workout,
        start:   time.Now(),
        conn:    conn,
    }
}

// Run starts the timer and broadcasts the progress of the workout to all connected clients.
func (t *Timer) Run() {
    for i := 1; i <= t.workout.rounds; i++ {
        message := fmt.Sprintf("Starting round %d\n", i)
        if err := t.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
            log.Println(err)
        }

        //
