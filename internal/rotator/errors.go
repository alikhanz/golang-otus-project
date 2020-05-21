package rotator

import "strconv"

type UnknownError struct {
	Message string
}

func (e UnknownError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return e.Name + ": not found"
}

type SlotHasNoBannersError struct {
	ID uint
}

func (e *SlotHasNoBannersError) Error() string {
	return "Slot [" + strconv.Itoa(int(e.ID)) + "]: has no banners"
}
