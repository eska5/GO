package main

import (
	"errors"
	"fmt"
)

type WeightFodder interface {
	FodderAmount() (float64, error)
}
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	foodWeight, err := weightFodder.FodderAmount()
	var ErrScaleMalfunction = errors.New("sensor error")
	var sillyNephewError SillyNephewError
	sillyNephewError.cows = cows

	if foodWeight >= 0 && err == ErrScaleMalfunction {
		foodWeight *= 2
		fmt.Println(foodWeight)
		return foodWeight / float64(cows), nil
	} else if foodWeight < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errors.New("negative fodder")
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, errors.New(sillyNephewError.Error())
	} else if err != nil {
		return 0, err
	}
	return foodWeight / float64(cows), nil
}
func main() {
	fmt.Println("1234")

}
