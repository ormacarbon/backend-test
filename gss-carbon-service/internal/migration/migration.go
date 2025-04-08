package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/icl00ud/backend-test/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, logger *zap.SugaredLogger) error {
	logger.Info("Starting database migrations...")

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20230403_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				logger.Infow("Running migration", "id", "20230403_create_users_table")
				err := tx.AutoMigrate(&model.User{})
				if err != nil {
					logger.Errorw("Migration failed", "id", "20230403_create_users_table", "error", err)
				}
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				logger.Infow("Rolling back migration", "id", "20230403_create_users_table")
				err := tx.Migrator().DropTable("users")
				if err != nil {
					logger.Errorw("Migration rollback failed", "id", "20230403_create_users_table", "error", err)
				}
				return err
			},
		},
	})

	err := m.Migrate()
	if err != nil {
		logger.Errorw("Migration process failed", "error", err)
		return err
	}

	logger.Info("Database migrations finished successfully.")
	return nil
}
