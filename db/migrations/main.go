package main

import (
	"gosuphu/config"
	"gosuphu/db/seeds"
	"gosuphu/models"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	logger     *logrus.Logger
	conf       *config.Config
	db         *gorm.DB
	uowFactory models.IUoWFactory
)

func init() {
	var err error
	conf = config.NewConfig()

	db, err = gorm.Open(
		postgres.Open(conf.DBConnectString),
		&gorm.Config{},
	)

	if err != nil {
		logger.Panicln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Panicln(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	uowFactory = models.NewUoWFactory(db)
}

const HELP_TEXT = `
You can try:
- migrate : to apply the model changes to database
- seed	  : to seed data to database
- help    : to display this help text
`

func main() {

	args := os.Args
	if len(args) <= 1 {
		println("You have not provide any arguments!" + HELP_TEXT)
		return
	}

	switch args[1] {
	case "migrate":
		db.AutoMigrate(
			&models.Campaign{},
			&models.Rarity{},
			&models.CampaignRarity{},
			&models.User{},
			&models.Company{},
		)
	case "seed":
		seeder := seeds.NewSeeder(uowFactory.GetDefaultUoW())
		if err := seeder.Seed(); err != nil {
			panic(err)
		}
	default:
		println("Argument is not recognized!" + HELP_TEXT)
		return
	}
}
