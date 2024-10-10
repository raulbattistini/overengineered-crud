package context

import (
	"context"
	"server/configs"
	"server/db"
	"server/enums"
	"server/hepers"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type AppContext struct {
	Connection         *mongo.Client
	PostgresConnection *gorm.DB
	MongoConnection    *mongo.Client
}

var (
	appContext *AppContext
	cancelFunc context.CancelFunc
)

func InitContext() {
	connection := configs.GetConnection()
	potsgresDb := db.DB
	mongoDb := db.MongoDB

	appContext = &AppContext{
		Connection:         connection,
		PostgresConnection: potsgresDb,
		MongoConnection:    mongoDb,
	}
}

func GetContext() *AppContext {
	return appContext
}

func GetCancelFunc() context.CancelFunc {
	return cancelFunc
}

func SetCancelFunc(cf context.CancelFunc) {
	cancelFunc = cf
}

func unplugContextVariables() error {
	if appContext.MongoConnection != nil {
		if err := db.MongoDB.Disconnect(context.TODO()); err != nil {
			return err
		}
	}
	if appContext.PostgresConnection != nil {
		appDb, err := db.DB.DB()
		if err != nil {
			return err
		}
		if err = appDb.Close(); err != nil {
			return err
		}
	}
	return nil
}

func CancelFunc() {
	defer func() {
		if err := unplugContextVariables(); err != nil {
			hepers.Log(err.Error(), &err, enums.Error)
		}
	}()

	if cancelFunc != nil {
		cancelFunc()
	}
}
