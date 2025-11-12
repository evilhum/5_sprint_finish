package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	base := strings.Split(datastring, ",")
	if len(base) != 2 {
		return fmt.Errorf("Error: wants 2 nums in p")
	}
	steps, err := strconv.Atoi(base[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("Error: zero steps in p")
	}
	dur, err := time.ParseDuration(base[1])
	if err != nil {
		return err
	}
	if dur <= 0 {
		return fmt.Errorf("Error: zero duration in p")
	}
	ds.Steps = steps
	ds.Duration = dur
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calory, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calory), err
}
