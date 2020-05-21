package rotator

import (
	"github.com/alikhanz/golang-otus-project/pkg/pb"
	"github.com/alikhanz/golang-otus-project/pkg/validator"
	"github.com/pkg/errors"
)

type Validator struct {
	r *Rotator
}

func NewValidator(r *Rotator) *Validator {
	return &Validator{
		r: r,
	}
}

func (*Validator) ValidateAddBannerToSlotRequest(r *pb.AddBannerToSlotRequest) error {
	v := validator.NewValidator()

	if v.HasErrors() {
		return errors.New(v.RenderErrors())
	}

	return nil
}

func (v *Validator) ValidateBannerExists(id uint) error {
	panic("implement me")
}