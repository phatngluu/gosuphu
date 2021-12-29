package seeds

import (
	"gosuphu/models"
)

type ISeeder interface {
	Seed()
}

type Seeder struct {
	uow models.IUoW
}

func NewSeeder(uow models.IUoW) *Seeder {
	return &Seeder{uow}
}

func (s *Seeder) Seed() error {
	var err error

	db := s.uow.GetDB()

	campaign := models.Campaign{
		CampaignCode: "animalia",
		CampaignName: "Animalia INO",
	}

	rarity := models.Rarity{
		RarityLevel: "Level 1",
	}

	campaign_rarity := models.CampaignRarity{
		CampaignID: campaign.ID,
		Campaign:   campaign,
		RarityID:   rarity.ID,
		Rarity:     rarity,
	}

	db.Create(&campaign_rarity)

	return err
}
