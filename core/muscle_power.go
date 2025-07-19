package core

import (
	"math"
	"time"
)

// KeystrokeEvent represents a single keystroke with timing
type KeystrokeEvent struct {
	Character   rune
	Timestamp   time.Time
	IsCorrect   bool
	IsBackspace bool
}

// MusclePowerIndicator tracks typing stamina and power
type MusclePowerIndicator struct {
	keystrokes      []KeystrokeEvent
	sessionStart    time.Time
	lastKeystroke   time.Time
	windowSize      time.Duration // sliding window for calculations
	maxWindowEvents int           // max events to keep in memory

	// Power metrics
	currentPower     float64
	peakPower        float64
	powerHistory     []PowerSnapshot
	consistencyScore float64
	staminaLevel     float64
	fatigueDetected  bool

	// Streak tracking
	correctStreak    int
	maxCorrectStreak int

	// Rhythm tracking
	avgKeystrokeDelay time.Duration
	rhythmVariance    float64
}

// PowerSnapshot represents power at a specific time
type PowerSnapshot struct {
	Timestamp time.Time
	Power     float64
	Status    PowerStatus
}

// PowerStatus represents different power states
type PowerStatus int

const (
	PowerStatusZenMode   PowerStatus = iota // Perfect consistency
	PowerStatusFullPower                    // High performance
	PowerStatusGoodFlow                     // Normal performance
	PowerStatusFatigue                      // Performance declining
	PowerStatusBurnout                      // Need rest
)

// PowerLevel represents the visual power level
type PowerLevel struct {
	Status     PowerStatus
	Percentage float64
	Message    string
	Icon       string
	Color      string
}

// NewMusclePowerIndicator creates a new MPI tracker
func NewMusclePowerIndicator() *MusclePowerIndicator {
	return &MusclePowerIndicator{
		keystrokes:      make([]KeystrokeEvent, 0),
		sessionStart:    time.Now(),
		windowSize:      30 * time.Second, // 30-second sliding window
		maxWindowEvents: 1000,             // Keep last 1000 keystrokes max
		powerHistory:    make([]PowerSnapshot, 0),
		currentPower:    1.0,
		peakPower:       1.0,
		staminaLevel:    1.0,
	}
}

// RecordKeystroke adds a new keystroke event
func (mpi *MusclePowerIndicator) RecordKeystroke(char rune, isCorrect bool, isBackspace bool) {
	now := time.Now()

	event := KeystrokeEvent{
		Character:   char,
		Timestamp:   now,
		IsCorrect:   isCorrect,
		IsBackspace: isBackspace,
	}

	// Add to keystrokes
	mpi.keystrokes = append(mpi.keystrokes, event)

	// Update streak
	if isCorrect && !isBackspace {
		mpi.correctStreak++
		if mpi.correctStreak > mpi.maxCorrectStreak {
			mpi.maxCorrectStreak = mpi.correctStreak
		}
	} else {
		mpi.correctStreak = 0
	}

	// Update timing
	if !mpi.lastKeystroke.IsZero() {
		delay := now.Sub(mpi.lastKeystroke)
		mpi.updateRhythm(delay)
	}
	mpi.lastKeystroke = now

	// Clean old events (keep within window and memory limits)
	mpi.cleanOldEvents()

	// Recalculate power
	mpi.calculatePower()
}

// cleanOldEvents removes events outside the sliding window
func (mpi *MusclePowerIndicator) cleanOldEvents() {
	cutoff := time.Now().Add(-mpi.windowSize)

	// Remove events older than window
	filtered := make([]KeystrokeEvent, 0)
	for _, event := range mpi.keystrokes {
		if event.Timestamp.After(cutoff) {
			filtered = append(filtered, event)
		}
	}

	// Also limit total events to prevent memory issues
	if len(filtered) > mpi.maxWindowEvents {
		start := len(filtered) - mpi.maxWindowEvents
		filtered = filtered[start:]
	}

	mpi.keystrokes = filtered
}

// updateRhythm calculates rhythm consistency
func (mpi *MusclePowerIndicator) updateRhythm(delay time.Duration) {
	// Simple moving average of keystroke delays
	if mpi.avgKeystrokeDelay == 0 {
		mpi.avgKeystrokeDelay = delay
	} else {
		// Exponential moving average
		alpha := 0.1
		mpi.avgKeystrokeDelay = time.Duration(float64(mpi.avgKeystrokeDelay)*(1-alpha) + float64(delay)*alpha)
	}

	// Calculate variance (simplified)
	variance := math.Abs(float64(delay - mpi.avgKeystrokeDelay))
	mpi.rhythmVariance = mpi.rhythmVariance*0.9 + variance*0.1
}

// calculatePower computes the current muscle power
func (mpi *MusclePowerIndicator) calculatePower() {
	if len(mpi.keystrokes) < 2 {
		mpi.currentPower = 1.0
		return
	}

	now := time.Now()
	windowStart := now.Add(-mpi.windowSize)

	var correctCount, incorrectCount, backspaceCount int
	var totalDelay time.Duration
	var eventCount int

	// Analyze events in current window
	for i := 0; i < len(mpi.keystrokes); i++ {
		event := mpi.keystrokes[i]
		if event.Timestamp.Before(windowStart) {
			continue
		}

		eventCount++

		if event.IsBackspace {
			backspaceCount++
		} else if event.IsCorrect {
			correctCount++
		} else {
			incorrectCount++
		}

		// Calculate delay between keystrokes
		if i > 0 {
			delay := event.Timestamp.Sub(mpi.keystrokes[i-1].Timestamp)
			totalDelay += delay
		}
	}

	if eventCount == 0 {
		return
	}

	// Calculate base metrics
	elapsedSeconds := time.Since(mpi.sessionStart).Seconds()
	if elapsedSeconds < 1 {
		elapsedSeconds = 1
	}

	// Core power calculation
	netCorrectKeystrokes := float64(correctCount - backspaceCount)
	if netCorrectKeystrokes < 0 {
		netCorrectKeystrokes = 0
	}

	// Power = (correct keystrokes - penalties) / time, with bonuses
	basePower := netCorrectKeystrokes / (elapsedSeconds / 60) // per minute

	// Consistency bonus (lower variance = higher bonus)
	consistencyBonus := 1.0
	if mpi.rhythmVariance > 0 {
		consistencyBonus = math.Max(0.5, 1.0-(mpi.rhythmVariance/1000000000)) // normalize variance
	}

	// Streak bonus
	streakBonus := 1.0 + math.Min(0.5, float64(mpi.correctStreak)/100.0)

	// Accuracy penalty
	accuracy := 1.0
	if correctCount+incorrectCount > 0 {
		accuracy = float64(correctCount) / float64(correctCount+incorrectCount)
	}

	// Final power calculation
	mpi.currentPower = basePower * consistencyBonus * streakBonus * accuracy

	// Update peak power
	if mpi.currentPower > mpi.peakPower {
		mpi.peakPower = mpi.currentPower
	}

	// Calculate stamina (power relative to peak, with time decay)
	if mpi.peakPower > 0 {
		mpi.staminaLevel = mpi.currentPower / mpi.peakPower

		// Add time-based fatigue after 2 minutes
		if elapsedSeconds > 120 {
			fatigueFactor := math.Max(0.3, 1.0-(elapsedSeconds-120)/600) // gradual decline
			mpi.staminaLevel *= fatigueFactor
		}
	}

	// Detect fatigue
	mpi.fatigueDetected = mpi.staminaLevel < 0.6 && elapsedSeconds > 30

	// Update consistency score
	mpi.consistencyScore = consistencyBonus

	// Record power snapshot
	mpi.powerHistory = append(mpi.powerHistory, PowerSnapshot{
		Timestamp: now,
		Power:     mpi.currentPower,
		Status:    mpi.getCurrentStatus(),
	})

	// Limit history size
	if len(mpi.powerHistory) > 100 {
		mpi.powerHistory = mpi.powerHistory[1:]
	}
}

// getCurrentStatus determines the current power status
func (mpi *MusclePowerIndicator) getCurrentStatus() PowerStatus {
	// If no keystrokes yet, show starting status
	if len(mpi.keystrokes) == 0 {
		return PowerStatusGoodFlow
	}

	// Zen Mode: Perfect consistency + good speed
	if mpi.consistencyScore > 0.95 && mpi.currentPower > 80 && mpi.correctStreak > 50 {
		return PowerStatusZenMode
	}

	// Full Power: High performance
	if mpi.staminaLevel > 0.8 && mpi.currentPower > 60 {
		return PowerStatusFullPower
	}

	// Good Flow: Normal performance
	if mpi.staminaLevel > 0.6 && mpi.currentPower > 30 {
		return PowerStatusGoodFlow
	}

	// Fatigue: Performance declining
	if mpi.staminaLevel > 0.3 {
		return PowerStatusFatigue
	}

	// Burnout: Need rest
	return PowerStatusBurnout
}

// GetCurrentPowerLevel returns the current power state with visual info
func (mpi *MusclePowerIndicator) GetCurrentPowerLevel() PowerLevel {
	status := mpi.getCurrentStatus()
	percentage := math.Min(100, mpi.staminaLevel*100)

	// Special case for initial state
	if len(mpi.keystrokes) == 0 {
		return PowerLevel{
			Status:     PowerStatusGoodFlow,
			Percentage: 100,
			Message:    "🚀 Ready to Type",
			Icon:       "🚀",
			Color:      "blue",
		}
	}

	switch status {
	case PowerStatusZenMode:
		return PowerLevel{
			Status:     status,
			Percentage: percentage,
			Message:    "🧘 Zen Mode Activated",
			Icon:       "🧘",
			Color:      "purple",
		}
	case PowerStatusFullPower:
		return PowerLevel{
			Status:     status,
			Percentage: percentage,
			Message:    "💪 Full Power",
			Icon:       "💪",
			Color:      "green",
		}
	case PowerStatusGoodFlow:
		return PowerLevel{
			Status:     status,
			Percentage: percentage,
			Message:    "⚡ Good Flow",
			Icon:       "⚡",
			Color:      "blue",
		}
	case PowerStatusFatigue:
		return PowerLevel{
			Status:     status,
			Percentage: percentage,
			Message:    "💤 Fatigue Mode",
			Icon:       "💤",
			Color:      "yellow",
		}
	case PowerStatusBurnout:
		return PowerLevel{
			Status:     status,
			Percentage: percentage,
			Message:    "🔥 Rest Needed",
			Icon:       "🔥",
			Color:      "red",
		}
	default:
		return PowerLevel{
			Status:     status,
			Percentage: percentage,
			Message:    "⚡ Starting Up",
			Icon:       "⚡",
			Color:      "gray",
		}
	}
}

// GetStats returns detailed MPI statistics
func (mpi *MusclePowerIndicator) GetStats() map[string]interface{} {
	elapsedSeconds := time.Since(mpi.sessionStart).Seconds()

	return map[string]interface{}{
		"current_power":       mpi.currentPower,
		"peak_power":          mpi.peakPower,
		"stamina_level":       mpi.staminaLevel,
		"consistency_score":   mpi.consistencyScore,
		"correct_streak":      mpi.correctStreak,
		"max_streak":          mpi.maxCorrectStreak,
		"fatigue_detected":    mpi.fatigueDetected,
		"session_duration":    elapsedSeconds,
		"total_keystrokes":    len(mpi.keystrokes),
		"avg_keystroke_delay": mpi.avgKeystrokeDelay.Milliseconds(),
	}
}

// Reset resets the MPI for a new session
func (mpi *MusclePowerIndicator) Reset() {
	mpi.keystrokes = make([]KeystrokeEvent, 0)
	mpi.sessionStart = time.Now()
	mpi.lastKeystroke = time.Time{}
	mpi.powerHistory = make([]PowerSnapshot, 0)
	mpi.currentPower = 1.0
	mpi.peakPower = 1.0
	mpi.consistencyScore = 1.0
	mpi.staminaLevel = 1.0
	mpi.fatigueDetected = false
	mpi.correctStreak = 0
	mpi.maxCorrectStreak = 0
	mpi.avgKeystrokeDelay = 0
	mpi.rhythmVariance = 0
}

// GetPowerBar generates a visual power bar
func (mpi *MusclePowerIndicator) GetPowerBar(width int) string {
	if width <= 0 {
		width = 20
	}

	level := mpi.GetCurrentPowerLevel()
	filled := int((level.Percentage / 100.0) * float64(width))

	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	return bar
}
