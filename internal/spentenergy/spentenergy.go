package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("amount of steps mustn't be 0")
	}
	if weight <= 30 || weight >= 300 {
		return 0, fmt.Errorf("weight must be between 30 and 300 kg, got %.2f", weight)
	}
	if height <= 0.5 || height >= 3.0 {
		return 0, fmt.Errorf("height must be between 0.5 and 3.0 m, got %.2f", height)
	}
	if duration <= 0 {
		return 0, errors.New("duration mustn't be 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	kkal := (weight * meanSpeed * duration.Minutes()) / minInH
	finalKkal := kkal * walkingCaloriesCoefficient

	return finalKkal, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("amount of steps mustn't be 0")
	}
	if weight <= 30 || weight >= 300 {
		return 0, fmt.Errorf("weight must be between 30 and 300 kg, got %.2f", weight)
	}
	if height <= 0.5 || height >= 3.0 {
		return 0, fmt.Errorf("height must be between 0.5 and 3.0 m, got %.2f", height)
	}
	if duration <= 0 {
		return 0, errors.New("duration mustn't be 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	kkal := (weight * meanSpeed * durationInMinutes) / minInH
	return kkal, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := (stepLength * float64(steps)) / mInKm
	return distance
}
