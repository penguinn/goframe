package cron

import (
	"github.com/robfig/cron"
)

type Cron struct {
	c *cron.Cron
}

var cr Cron

func GetCron() *Cron {
	return &cr
}

func (p Cron) RegisterCron(spec string, cmd func()) error {
	return p.c.AddFunc(spec, cmd)
}

func (p Cron) RegisterJob(spec string, cmd cron.Job) error {
	return p.c.AddJob(spec, cmd)
}

func (p Cron) StartCron() {
	p.c.Start()
}

func (p Cron) StopCron() {
	p.c.Stop()
}
