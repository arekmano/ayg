package service

import (
	"github.com/arekmano/ayg/activity"
	"github.com/arekmano/ayg/store"
)

type TimeKeeper struct {
	activity.ActivityDetector
	store.TimeStore
}

func (t *TimeKeeper) Monitor() error {

}
