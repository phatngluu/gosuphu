package models

type IRepoManager interface {
	GetCampaignRepo(IUoW) ICampaignRepo
	GetRarityRepo(IUoW) IRarityRepo
	GetCampaignRarityRepo(IUoW) ICampaignRarityRepo
}

type RepoManager struct {
}

func (m RepoManager) GetCampaignRepo(uow IUoW) ICampaignRepo {
	return NewCampaignRepo(uow)
}

func (m RepoManager) GetRarityRepo(uow IUoW) IRarityRepo {
	return NewRarityRepo(uow)
}

func (m RepoManager) GetCampaignRarityRepo(uow IUoW) ICampaignRarityRepo {
	return NewCampaignRarityRepo(uow)
}
