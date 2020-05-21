// +build integration

package test

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/alikhanz/golang-otus-project/internal/helper"
	"github.com/alikhanz/golang-otus-project/internal/model"
	"github.com/alikhanz/golang-otus-project/internal/resources"
	"github.com/alikhanz/golang-otus-project/internal/rotator"
	"github.com/alikhanz/golang-otus-project/internal/server"
	"github.com/alikhanz/golang-otus-project/pkg/os_context"
	"github.com/alikhanz/golang-otus-project/pkg/pb"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	res := resources.Get(ctx)

	prepareDb(res)
	seed(res)
	go runServer(res, ctx)
	time.Sleep(2*time.Second)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func runServer(res *resources.Resources, ctx context.Context) {
	r := rotator.NewRotator(res)
	s := server.NewServer(
		r,
		res,
	)
	log.Fatal().Err(s.Run(ctx))
}

func prepareDb(res *resources.Resources) {
	helper.Migrate(res)

	models := []interface{} {
		&model.Banner{},
		&model.Slot{},
		&model.Sdg{},
		&model.Stat{},
	}

	res.DB.Exec("TRUNCATE TABLE banner_slots")

	for _, m := range models {
		table := res.DB.NewScope(m).TableName()
		res.DB.Exec("TRUNCATE TABLE " + table)
	}

}

var slot1 *model.Slot
var slot2 *model.Slot
var banner1 *model.Banner
var banner2 *model.Banner
var sdg1 *model.Sdg
var sdg2 *model.Sdg

func seed(res *resources.Resources) {
	slot1 = &model.Slot{Description: "Слот 1"}
	slot2 = &model.Slot{Description: "Слот 2"}
	banner1 = &model.Banner{Description: "Баннер 1"}
	banner2 = &model.Banner{Description: "Баннер 2"}
	sdg1 = &model.Sdg{Description: "Бабушки"}
	sdg2 = &model.Sdg{Description: "Дедушки"}

	res.DB.Create(slot1)
	res.DB.Create(slot2)
	res.DB.Create(banner1)
	res.DB.Create(banner2)
	res.DB.Create(sdg1)
	res.DB.Create(sdg2)

	res.DB.Model(banner1).Association("Slots").Append(slot1)
	res.DB.Model(banner1).Association("Slots").Append(slot2)

	res.DB.Model(banner2).Association("Slots").Append(slot1)
}

func TestBannerSelection(t *testing.T) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("localhost:"+os.Getenv("GRPC_PORT"), opts...)

	if err != nil {
		log.Fatal().Msgf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewRotatorClient(conn)
	ctx := os_context.Context(context.Background())

	rand.Seed(time.Now().Unix())
	selectedBanners := make(map[int]int)

	// В 1 соц. дем. группе 1 слота должен гораздо чаще отдаваться 1-й баннер.
	// Имитируем хиты по 2 баннерам в одном слоте в одной соц. дем. группе,
	// проверяем, что впоследствии первый станет выбираться значительно чаще, т.к. по нему будет значимо больше хитов.
	const TRIALS = 500
	for i := 0; i < TRIALS; i++ {
		// Выбираем баннер. Если выпадает 1-й - сразу же делаем хит по нему. Если 2-й, хитуем в 30% случаях.
		req := &pb.SelectBannerRequest{
			SlotId: uint32(slot1.ID),
			SdgId:  uint32(sdg1.ID),
		}
		res, err := client.SelectBanner(ctx, req)

		assert.NoError(t, err)
		selectedBanners[int(res.BannerId)]++

		switch res.BannerId {
			case uint32(banner1.ID):
				req := &pb.HitBannerRequest{
					BannerId: uint32(banner1.ID),
					SlotId:   uint32(slot1.ID),
					SdgId:    uint32(sdg1.ID),
				}
				_, err = client.HitBanner(ctx, req)
				assert.NoError(t, err)
				break
			case uint32(banner2.ID):
				r := rand.Float32()
				if r <= 0.3 {
					req := &pb.HitBannerRequest{
						BannerId: uint32(banner2.ID),
						SlotId:   uint32(slot1.ID),
						SdgId:    uint32(sdg1.ID),
					}
					_, err = client.HitBanner(ctx, req)
					assert.NoError(t, err)
				}
				break
		}
	}

	// Проверяем, что баннер1 был показан более чем в 70% случаях
	assert.GreaterOrEqual(
		t,
		float32(selectedBanners[int(banner1.ID)])/TRIALS*100,
		float32(70),
		"first arm must have more than 70% hits",
	)
	// Проверяем, что баннер2 был показан хотя бы в 1% или более случаях
	assert.GreaterOrEqual(
		t,
		float32(selectedBanners[int(banner2.ID)])/TRIALS*100,
		float32(1),
		"second arm must have more than 1% hits",
	)
}
