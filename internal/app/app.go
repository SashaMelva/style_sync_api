package app

import (
	"regexp"

	"github.com/SashaMelva/style_sync_api/internal/storage"
	"go.uber.org/zap"
)

type App struct {
	storage     *storage.Storage
	Logger      *zap.SugaredLogger
	RegexRegNum *regexp.Regexp
}

func New(logger *zap.SugaredLogger, storage *storage.Storage) *App {
	r, _ := regexp.Compile(`^[ABEKMHOPCTYXАВЕКМНОРСТУХ][0-9][0-9][0-9][ABEKMHOPCTYXАВЕКМНОРСТУХ][ABEKMHOPCTYXАВЕКМНОРСТУХ][0-9][0-9][0-9]$`)
	return &App{
		storage:     storage,
		Logger:      logger,
		RegexRegNum: r,
	}
}
