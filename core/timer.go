package core

import (
	"time"
)

// Timer handles timing functionality for typing sessions
type Timer struct {
	startTime time.Time
	endTime   time.Time
	running   bool
}

// NewTimer creates a new timer instance
func NewTimer() *Timer {
	return &Timer{}
}

// Start starts the timer
func (t *Timer) Start() {
	t.startTime = time.Now()
	t.running = true
	t.endTime = time.Time{} // Reset end time
}

// Stop stops the timer
func (t *Timer) Stop() {
	if t.running {
		t.endTime = time.Now()
		t.running = false
	}
}

// Reset resets the timer
func (t *Timer) Reset() {
	t.startTime = time.Time{}
	t.endTime = time.Time{}
	t.running = false
}

// IsRunning returns true if timer is currently running
func (t *Timer) IsRunning() bool {
	return t.running
}

// Elapsed returns the elapsed time
func (t *Timer) Elapsed() time.Duration {
	if t.startTime.IsZero() {
		return 0
	}

	if t.running {
		return time.Since(t.startTime)
	}

	if !t.endTime.IsZero() {
		return t.endTime.Sub(t.startTime)
	}

	return 0
}

// TotalTime returns the total session time (for completed sessions)
func (t *Timer) TotalTime() time.Duration {
	if t.startTime.IsZero() {
		return 0
	}

	if !t.endTime.IsZero() {
		return t.endTime.Sub(t.startTime)
	}

	// If still running, return current elapsed time
	return time.Since(t.startTime)
}
