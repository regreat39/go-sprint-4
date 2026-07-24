package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parsInfo := strings.Split(data, ",")
	if len(parsInfo) != 2 {
		return 0, 0, fmt.Errorf("ошибка формата. Ожидается 'шаги, продолжительность', получено '%s'", data)
	}
	quantityStep, err := strconv.Atoi(parsInfo[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}
	if quantityStep <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parsInfo[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка продолжительности: %v", err)
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	return quantityStep, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	step, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if step <= 0 {
		return ""
	}
	distance := float64(step) * stepLength
	distanceKm := distance / float64(mInKm)

	calories, err := spentcalories.WalkingSpentCalories(step, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d\nДистанция составила %.2f км\nВы сожгли %.2f ккал\n", step, distanceKm, calories)
	return result
}
