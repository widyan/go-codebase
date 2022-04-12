package scheduller

func (c *SchedullerImpl) InitJob() {
	c.CronWorker.AddJob(c.Ctx, "TestScheduller", "*/3 * * * *", c.TestScheduller)
	c.CronWorker.AddJob(c.Ctx, "TestLagiAh", "*/1 * * * *", c.TestLagiAh)
	c.CronWorker.AddJob(c.Ctx, "TestTambahWorker", "*/1 * * * *", c.TestTambahWorker)
	c.CronWorker.AddJob(c.Ctx, "TestTambahWorkerLagiDuh", "/5 * * * *", c.TestTambahWorkerLagiDuh)
	c.CronWorker.AddJob(c.Ctx, "TestDasarKampret", "*/1 * * * *", c.TestDasarKampret)

	c.CronWorker.SetListWorker(c.Ctx)
}
