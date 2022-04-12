package scheduller

func (c *SchedullerImpl) InitJob() {
	c.CronWorker.AddJob(c.Ctx, "TestScheduller", "*/3 * * * *", c.TestScheduller)
	c.CronWorker.AddJob(c.Ctx, "TestLagiAh", "*/1 * * * *", c.TestLagiAh)
	c.CronWorker.AddJob(c.Ctx, "TestTambahWorker", "* * * * *", c.TestTambahWorker)
	c.CronWorker.AddJob(c.Ctx, "TestTambahWorkerLagiDuh", "* * * * *", c.TestTambahWorkerLagiDuh)
	c.CronWorker.AddJob(c.Ctx, "TestDasarKampret", "* * * * *", c.TestDasarKampret)

	c.CronWorker.SetListWorker(c.Ctx)
}
