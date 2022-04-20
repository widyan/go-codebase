package scheduller

func (c *SchedullerImpl) InitJob() {
	c.CronWorker.AddJob("*/1 * * * *", c.TestScheduller)
	// c.CronWorker.AddJob("*/2 * * * *", c.TestLagiAh)
	// c.CronWorker.AddJob("*/3 * * * *", c.TestTambahWorker)
	// c.CronWorker.AddJob("*/4 * * * *", c.TestTambahWorkerLagiDuh)
	// c.CronWorker.AddJob("*/5 * * * *", c.TestDasarKampret)
	// c.CronWorker.AddJob("*/6 * * * *", c.TambahWorkerLagi)
	c.CronWorker.AddJob("*/7 * * * *", c.RetestWorker)

	c.CronWorker.SetListWorker(c.Ctx)
}
