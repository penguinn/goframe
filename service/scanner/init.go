package scanner

import (
	"github.com/penguinn/goframe/component/cron"
)

func Init() error {
	c := cron.GetCron()
	//err := c.RegisterCron("*/10 * * * * ?", someServices)
	//if err != nil {
	//	return err
	//}
	c.StartCron()

	return nil
}
