//go:build !windows
// +build !windows

package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func InitSysNotify() {
	c = make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP /*1*/, syscall.SIGINT /*2*/, syscall.SIGQUIT /*3*/, syscall.SIGTERM /*15*/, syscall.SIGUSR1 /*10*/, syscall.SIGUSR2 /*12*/)
	WaitCbk = func() {
		once.Do(waitCbk)
	}
}

func waitCbk() {
	log.Info("wait sys notify")
	for s := range c {
		log.Infof("receive sys notify: %v", s)
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			tq, ok := TaskGroup.Load(NotifyTypeEXIT)
			if ok {
				log.Info("task: NotifyTypeEXIT running...")
				runTask(tq)
			}
			return
		case syscall.SIGUSR1, syscall.SIGUSR2:
			tq, ok := TaskGroup.Load(NotifyTypeRELOAD)
			if ok {
				log.Info("task: NotifyTypeRELOAD running...")
				runTask(tq)
			}
		}
		log.Info("task: all done")
	}
}
