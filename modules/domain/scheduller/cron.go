package scheduller

func (c *SchedullerImpl) InitJob() {
	c.CronWorker.AddJob("TestScheduller", "*/3 * * * *", c.TestScheduller)
	c.CronWorker.AddJob("TestLagiAh", "*/1 * * * *", c.TestLagiAh)
	c.CronWorker.AddJob("TestTambahWorker", "*/1 * * * *", c.TestTambahWorker)
	c.CronWorker.AddJob("TestTambahWorkerLagiDuh", "/5 * * * *", c.TestTambahWorkerLagiDuh)
	c.CronWorker.AddJob("TestDasarKampret", "*/1 * * * *", c.TestDasarKampret)
	c.CronWorker.AddJob("TambahWorkerLagi", "*/1 * * * *", c.TestDasarKampret)

	c.CronWorker.SetListWorker(c.Ctx)
}
