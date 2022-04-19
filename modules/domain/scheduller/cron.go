package scheduller

func (c *SchedullerImpl) InitJob() {
	c.CronWorker.AddJob("*/3 * * * *", c.TestScheduller)
	c.CronWorker.AddJob("*/1 * * * *", c.TestLagiAh)
	c.CronWorker.AddJob("*/1 * * * *", c.TestTambahWorker)
	c.CronWorker.AddJob("/5 * * * *", c.TestTambahWorkerLagiDuh)
	c.CronWorker.AddJob("*/1 * * * *", c.TestDasarKampret)
	c.CronWorker.AddJob("*/1 * * * *", c.TestDasarKampret)

	c.CronWorker.SetListWorker(c.Ctx)
}
