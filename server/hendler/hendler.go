package hendler

import (
	"sync"

	"github.com/SashaMelva/style_sync_api/internal/app"
	"go.uber.org/zap"
)

type Service struct {
	log zap.SugaredLogger
	app app.App
	sync.RWMutex
}

func NewHendler(log *zap.SugaredLogger, application *app.App) *Service {
	return &Service{
		log: *log,
		app: *application,
	}
}
