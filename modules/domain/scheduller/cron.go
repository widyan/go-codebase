package scheduller

func (c *SchedullerImpl) InitJob() {
	c.CronWorker.AddJob(c.Ctx, "TestScheduller", "*/1 * * * *", c.TestScheduller)
	c.CronWorker.AddJob(c.Ctx, "TestLagiAh", "*/1 * * * *", c.TestLagiAh)
}
