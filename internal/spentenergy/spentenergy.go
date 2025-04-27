package spentenergy

import (
	"errors"
	"log"
	"time"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		log.Println("Weight must be greater than zero")
		return 0, errors.New("Weight must be greater than zero")
	}
	if height <= 0 {
		log.Println("Height must be greater than zero")
		return 0, errors.New("Height must be greater than zero")
	}
	if steps <= 0 {
		log.Println("Steps must be greater than zero")
		return 0, errors.New("Steps must be greater than zero")
	}
	if duration <= 0 {
		log.Println("Duration must be greater than zero")
		return 0, errors.New("Duration must be greater than zero")
	}
	return (MeanSpeed(steps, height, duration) * weight * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		log.Println("Weight must be greater than zero")
		return 0, errors.New("Weight must be greater than zero")
	}
	if height <= 0 {
		log.Println("Height must be greater than zero")
		return 0, errors.New("Height must be greater than zero")
	}
	if steps <= 0 {
		log.Println("Steps must be greater than zero")
		return 0, errors.New("Steps must be greater than zero")
	}
	if duration <= 0 {
		log.Println("Duration must be greater than zero")
		return 0, errors.New("Duration must be greater than zero")
	}
	return (MeanSpeed(steps, height, duration) * weight * duration.Minutes()) / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		log.Println("Duration must be greater than zero")
		return 0
	}
	return Distance(steps, height) / float64(duration.Hours())
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		log.Println("Steps must be greater than zero")
		return 0
	}
	if height <= 0 {
		log.Println("Height must be greater than zero")
		return 0
	}
	return (height * stepLengthCoefficient) * float64(steps) / mInKm
}
