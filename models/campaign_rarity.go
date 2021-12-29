package models

type CampaignRarity struct {
	Model
	CampaignID uint     `json:"campaign_id"`
	RarityID   uint     `json:"rarity_id"`
	Campaign   Campaign `json:"campaign"    gorm:"foreignKey: "`
	Rarity     Rarity   `json:"rarity"`
}

type ICampaignRarityRepo interface {
	IRepo
	CreateCampaignRarity(r CampaignRarity) error
	GetCampaignRarityById(id uint) (CampaignRarity, error)
}

type CampaignRarityRepo struct {
	*Repo
}

func NewCampaignRarityRepo(uow IUoW) ICampaignRarityRepo {
	repo := NewRepo(uow)
	return &CampaignRarityRepo{Repo: repo}
}

func (r *CampaignRarityRepo) CreateCampaignRarity(campaignRarity CampaignRarity) error {
	return r.DB.Model(&Campaign{}).Create(&campaignRarity).Error
}

func (r *CampaignRarityRepo) GetCampaignRarityById(id uint) (CampaignRarity, error) {
	var campaignRarity CampaignRarity

	err := r.DB.Where("id = ?", id).First(&campaignRarity).Error
	if err != nil {
		return CampaignRarity{}, err
	}

	return campaignRarity, nil
}
