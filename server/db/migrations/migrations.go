package migrations

import (
	"log"
	"server/hepers"
	"server/types"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	migrations := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: hepers.GenUuidStr(),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&types.Post{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("posts")
			},
		},
	})

	if err := migrations.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
		return err
	}
	log.Println("Migrations ran successfully")
	return nil
}
