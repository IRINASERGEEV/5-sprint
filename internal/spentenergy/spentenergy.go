package spentenergy

import (
	"errors"
	"time"
)

var ErrNull = errors.New("error: steps, duration, height, weight cannot be less than or equal to 0")

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 || height <= 0 || weight <= 0 {
		return 0, ErrNull
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 || height <= 0 || weight <= 0 {
		return 0, ErrNull
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationMinutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps < 0 || duration <= 0 || height <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	durationHours := duration.Hours()
	meanSpeed := distance / durationHours
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps < 0 || height <= 0 {
		return 0
	}
	step_length := height * stepLengthCoefficient
	distance := float64(steps) * step_length / mInKm
	return distance
}
