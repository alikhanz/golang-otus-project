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

func (r *Rotator) AddBannerToSlot(bannerId, slotId uint) error {
	banner, slot, err := r.findBannerAndSlot(bannerId, slotId)

	if err != nil {
		return err
	}

	r.res.Db.Model(banner).Association("Slots").Append(slot)

	return nil
}

func (r *Rotator) HitBanner(bannerId, slotId, sdgId uint) error {
	banner, err := r.findBanner(bannerId)

	if err != nil {
		return err
	}

	stat, err := r.findOrCreateBannerStat(banner.ID, slotId, sdgId)
	if err != nil {
		return err
	}

	res := r.res.Db.Model(stat).
		UpdateColumn("hit_count", gorm.Expr("hit_count + ?", 1))

	if res.Error != nil {
		return UnknownError{"Failed stat update"}
	}

	return nil
}

func (r *Rotator) SelectBanner(slotId, sdgId uint) (*model.Banner, error) {
	slot, err := r.findSlot(slotId)
	if err != nil {
		return nil, err
	}

	r.res.Db.Model(slot).Association("Banners").Find(&slot.Banners)

	if len(slot.Banners) == 0 {
		return nil, &SlotHasNoBannersError{ID: slot.ID}
	}

	sdg, err := r.findSdg(sdgId)
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
	res := r.res.Db.Model(bannerStat[selectedBanner.ID]).
		UpdateColumn("show_count", gorm.Expr("show_count + ?", 1))

	if res.Error != nil {
		return nil, UnknownError{"Failed stat update"}
	}

	return slot.Banners[i], nil
}

func (r *Rotator) RemoveBannerFromSlot(bannerId, slotId uint) error {
	banner, slot, err := r.findBannerAndSlot(bannerId, slotId)
	if err != nil {
		return err
	}

	res := r.res.Db.Model(banner).Association("Slots").Delete(slot)
	if res.Error != nil {
		return UnknownError{"Removing banner from slot failed"}
	}

	return nil
}

func (r *Rotator) findBannerAndSlot(bannerId, slotId uint) (*model.Banner, *model.Slot, error) {
	banner, err := r.findBanner(bannerId)

	if err != nil {
		return nil, nil, err
	}

	slot, err := r.findSlot(slotId)

	if err != nil {
		return nil, nil, err
	}

	return banner, slot, nil
}

func (r *Rotator) findBanner(bannerId uint) (*model.Banner, error) {
	banner := &model.Banner{}

	result := r.res.Db.First(banner, bannerId)

	if result.RecordNotFound() {
		return nil, &NotFoundError{"Banner"}
	}

	if result.Error != nil {
		return nil, &UnknownError{"Banner search failed"}
	}

	return banner, nil
}

func (r *Rotator) findSlot(slotId uint) (*model.Slot, error) {
	slot := &model.Slot{}

	result := r.res.Db.First(slot, slotId)

	if result.RecordNotFound() {
		return nil, &NotFoundError{"Slot"}
	}

	if result.Error != nil {
		return nil, UnknownError{"Slot search failed"}
	}

	return slot, nil
}

func (r *Rotator) findSdg(sdgId uint) (*model.Sdg, error) {
	sdg := &model.Sdg{}

	result := r.res.Db.First(sdg, sdgId)

	if result.RecordNotFound() {
		return nil, &NotFoundError{"Sdg"}
	}

	if result.Error != nil {
		return nil, UnknownError{"Sdg search failed"}
	}

	return sdg, nil
}

func (r *Rotator) findOrCreateBannerStat(bannerId, slotId, sdgId uint) (*model.Stat, error) {
	stat := &model.Stat{}
	result := r.res.Db.FirstOrCreate(stat, &model.Stat{BannerId: bannerId, SlotId: slotId, SdgId: sdgId})

	if result.Error != nil {
		return nil, UnknownError{"Stats search failed"}
	}

	return stat, nil
}