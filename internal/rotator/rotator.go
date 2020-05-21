package rotator

import (
	"github.com/alikhanz/golang-otus-project/internal/model"
	"github.com/alikhanz/golang-otus-project/internal/resources"
	"github.com/alikhanz/golang-otus-project/pkg/ucb1"
	"github.com/jinzhu/gorm"
)

type Rotator struct {
	res *resources.Resources
}

func NewRotator(res *resources.Resources) *Rotator {
	return &Rotator{res: res}
}

func (r *Rotator) AddBannerToSlot(bannerID, slotID uint) error {
	banner, slot, err := r.findBannerAndSlot(bannerID, slotID)

	if err != nil {
		return err
	}

	r.res.DB.Model(banner).Association("Slots").Append(slot)

	return nil
}

func (r *Rotator) HitBanner(bannerID, slotID, sdgID uint) error {
	banner, err := r.findBanner(bannerID)

	if err != nil {
		return err
	}

	stat, err := r.findOrCreateBannerStat(banner.ID, slotID, sdgID)
	if err != nil {
		return err
	}

	res := r.res.DB.Model(stat).
		UpdateColumn("hit_count", gorm.Expr("hit_count + ?", 1))

	if res.Error != nil {
		return UnknownError{"Failed stat update"}
	}

	return nil
}

func (r *Rotator) SelectBanner(slotID, sdgID uint) (*model.Banner, error) {
	slot, err := r.findSlot(slotID)
	if err != nil {
		return nil, err
	}

	r.res.DB.Model(slot).Association("Banners").Find(&slot.Banners)

	if len(slot.Banners) == 0 {
		return nil, &SlotHasNoBannersError{ID: slot.ID}
	}

	sdg, err := r.findSdg(sdgID)
	if err != nil {
		return nil, err
	}

	arms := make(ucb1.Arms, 0, len(slot.Banners))
	bannerStat := make(map[uint]*model.Stat)

	for _, banner := range slot.Banners {
		stat, err := r.findOrCreateBannerStat(banner.ID, slot.ID, sdg.ID)

		if err != nil {
			return nil, nil
		}
		bannerStat[banner.ID] = stat
		arms = append(
			arms,
			ucb1.Arm{
				Count:  stat.ShowCount,
				Reward: float64(stat.HitCount),
			},
		)
	}

	i := ucb1.UCB1(arms)
	selectedBanner := slot.Banners[i]
	res := r.res.DB.Model(bannerStat[selectedBanner.ID]).
		UpdateColumn("show_count", gorm.Expr("show_count + ?", 1))

	if res.Error != nil {
		return nil, UnknownError{"Failed stat update"}
	}

	return slot.Banners[i], nil
}

func (r *Rotator) RemoveBannerFromSlot(bannerID, slotID uint) error {
	banner, slot, err := r.findBannerAndSlot(bannerID, slotID)
	if err != nil {
		return err
	}

	res := r.res.DB.Model(banner).Association("Slots").Delete(slot)
	if res.Error != nil {
		return UnknownError{"Removing banner from slot failed"}
	}

	return nil
}

func (r *Rotator) findBannerAndSlot(bannerID, slotID uint) (*model.Banner, *model.Slot, error) {
	banner, err := r.findBanner(bannerID)

	if err != nil {
		return nil, nil, err
	}

	slot, err := r.findSlot(slotID)

	if err != nil {
		return nil, nil, err
	}

	return banner, slot, nil
}

func (r *Rotator) findBanner(bannerID uint) (*model.Banner, error) {
	banner := &model.Banner{}

	result := r.res.DB.First(banner, bannerID)

	if result.RecordNotFound() {
		return nil, &NotFoundError{"Banner"}
	}

	if result.Error != nil {
		return nil, &UnknownError{"Banner search failed"}
	}

	return banner, nil
}

func (r *Rotator) findSlot(slotID uint) (*model.Slot, error) {
	slot := &model.Slot{}

	result := r.res.DB.First(slot, slotID)

	if result.RecordNotFound() {
		return nil, &NotFoundError{"Slot"}
	}

	if result.Error != nil {
		return nil, UnknownError{"Slot search failed"}
	}

	return slot, nil
}

func (r *Rotator) findSdg(sdgID uint) (*model.Sdg, error) {
	sdg := &model.Sdg{}

	result := r.res.DB.First(sdg, sdgID)

	if result.RecordNotFound() {
		return nil, &NotFoundError{"Sdg"}
	}

	if result.Error != nil {
		return nil, UnknownError{"Sdg search failed"}
	}

	return sdg, nil
}

func (r *Rotator) findOrCreateBannerStat(bannerID, slotID, sdgID uint) (*model.Stat, error) {
	stat := &model.Stat{}
	result := r.res.DB.FirstOrCreate(stat, &model.Stat{BannerID: bannerID, SlotID: slotID, SdgID: sdgID})

	if result.Error != nil {
		return nil, UnknownError{"Stats search failed"}
	}

	return stat, nil
}
