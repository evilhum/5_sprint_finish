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
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	base := strings.Split(datastring, ",")
	if len(base) != 3 {
		return fmt.Errorf("Error: wants 3 arg in p")
	}
	steps, err := strconv.Atoi(base[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("Error: zero steps in p")
	}
	duration, err := time.ParseDuration(base[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("Error: zero duration in p")
	}
	t.Steps = steps
	t.TrainingType = base[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	mspeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	if t.TrainingType == "Бег" {
		running, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, mspeed, running), nil
	}
	if t.TrainingType == "Ходьба" {
		walking, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, mspeed, walking), nil
	}
	return "", fmt.Errorf("Error: wrong action type in AI")
}
