package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parsTrain := strings.Split(data, ",")
	if len(parsTrain) != 3 {
		return 0, "", 0, fmt.Errorf("ошибка формата. Ожидается 'шаги, тип активности, продолжительность', получено '%s'", data)
	}
	quantityStep, err := strconv.Atoi(parsTrain[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}
	if quantityStep <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parsTrain[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка продолжительности: %v", err)
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	activity := parsTrain[1]

	return quantityStep, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	totalDistance := (float64(steps) * stepLength) / float64(mInKm)
	return totalDistance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dis := distance(steps, height) // ПРОВЕРИТЬ
	averageSpeed := dis / duration.Hours()
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64
	switch activity {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	durationHours := duration.Hours()

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, durationHours, dist, speed, calories)
	return result, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("недопустимое количество шагов (<= 0)")
	}
	if weight <= 0 || weight > 300 {
		return 0, fmt.Errorf("недопустимый вес (<=0 или >300)")
	}
	if height <= 0 || height > 3 {
		return 0, fmt.Errorf("недопустимый рост (<=0 или >3)")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("недопустимое время (<=0)")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	result := (weight * meanSpeed * durationInMinutes) / minInH
	return result, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("недопустимое количество шагов (<= 0)")
	}
	if weight <= 0 || weight > 300 {
		return 0, fmt.Errorf("недопустимый вес (<=0 или >300)")
	}
	if height <= 0 || height > 3 {
		return 0, fmt.Errorf("недопустимый рост (<=0 или >3)")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("недопустимое время (<=0)")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	result := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
	return result, nil
}
