package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/icl00ud/backend-test/internal/model"
	"gorm.io/gorm"
)

// Migrate executa as migrações usando gormigrate.
func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230403_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	})
	return m.Migrate()
}
