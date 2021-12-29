package models

type Campaign struct {
	Model
	ID           uint   `json:"id"`
	CampaignCode string `json:"campaign_code"`
	CampaignName string `json:"campaign_name"`
}

type ICampaignRepo interface {
	IRepo
	CreateCampaign(c Campaign) error
	GetCampaignById(id uint) (Campaign, error)
}

type CampaignRepo struct {
	*Repo
}

func NewCampaignRepo(uow IUoW) ICampaignRepo {
	repo := NewRepo(uow)
	return &CampaignRepo{Repo: repo}
}

func (r *CampaignRepo) CreateCampaign(campaign Campaign) error {
	return r.DB.Model(&Campaign{}).Create(&campaign).Error
}

func (r *CampaignRepo) GetCampaignById(id uint) (Campaign, error) {
	var campaign Campaign

	err := r.DB.Where("id = ?", id).First(&campaign).Error
	if err != nil {
		return Campaign{}, err
	}

	return campaign, nil
}
