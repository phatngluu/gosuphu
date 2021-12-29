package models

type Rarity struct {
	Model
	ID          uint   `json:"id"`
	RarityLevel string `json:"rarity_level"`
}

type IRarityRepo interface {
	IRepo
	CreateRarity(r Rarity) error
	GetRarityById(id uint) (Rarity, error)
}

type RarityRepo struct {
	*Repo
}

func NewRarityRepo(uow IUoW) IRarityRepo {
	repo := NewRepo(uow)
	return &RarityRepo{Repo: repo}
}

func (r *RarityRepo) CreateRarity(rarity Rarity) error {
	return r.DB.Model(&Campaign{}).Create(&rarity).Error
}

func (r *RarityRepo) GetRarityById(id uint) (Rarity, error) {
	var rarity Rarity

	err := r.DB.Where("id = ?", id).First(&rarity).Error
	if err != nil {
		return Rarity{}, err
	}

	return rarity, nil
}
