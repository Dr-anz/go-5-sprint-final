package spentenergy

import (
	"errors"
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
	// проверка на некорректные параметры
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("the parameters are incorrect")
	}
	// расчет средней скорости
	meanSpeed := MeanSpeed(steps, height, duration)
	// перевод продолжительности в минуты
	durationInMinutes := duration.Minutes()
	// расчет количетва калорий
	calories := (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка на некорректные параметры
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("the parameters are incorrect")
	}
	// расчет средней скорости
	meanSpeed := MeanSpeed(steps, height, duration)
	// перевод продолжительности в минуты
	durationInMinutes := duration.Minutes()
	// расчет количества калорий
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверка на отрицательную продолжительность или шаги
	if duration <= 0 || steps <= 0 {
		return 0
	}
	// вычислении дистанции
	distance := Distance(steps, height)
	// перевод продолжительности в часы
	durationInHours := duration.Hours()
	// расчет и возврат средней скорости
	return distance / durationInHours
}

func Distance(steps int, height float64) float64 {
	// приобразование steps в float64
	stepsFloat := float64(steps)
	// расчет длины шага
	stepLength := height * stepLengthCoefficient
	// вычисление общей длины в метрах
	totalLengthMeters := stepsFloat * stepLength
	// перевод в км и возврат результата
	return totalLengthMeters / mInKm
}
