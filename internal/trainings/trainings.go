package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// разделяем строку на части
	parts := strings.Split(datastring, ",")
	// проверяем что частей ровно 3
	if len(parts) != 3 {
		return err
	}
	// парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	// сохраняем тип тренировки
	t.TrainingType = parts[1]
	// парсинг длительности тренировки
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	// вычисляем среднюю скорость
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var calories float64
	var err error

	// рассчитываем калории в зависимости от типа тренировки
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	// формируем итоговую строку
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

	return result, nil
}
