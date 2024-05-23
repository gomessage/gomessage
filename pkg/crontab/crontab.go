package crontab

import (
	"gomessage/pkg/models"

	"github.com/robfig/cron/v3"
)

var c *cron.Cron

var jobs = map[int]cron.EntryID{}

// 必须在main函数里面调用
func InitCrontab() {
	c = cron.New()
	initJobs()
	c.Run()
}

func initJobs() {
	list, err := models.ListCrontabs()
	if err != nil {
		// TODO: error
		return
	}

	for _, ct := range list {
		if !ct.CrontabActivate {
			continue
		}
		err := AddJob(ct)
		if err != nil {
			// TODO: error
			return
		}
	}
}

// 添加定时任务
func AddJob(ct models.Crontab) error {
	_, ok := jobs[ct.ID]
	if ok {
		return nil
	}
	f := genJobFunc(ct)
	entryID, err := c.AddFunc(ct.CrontabRule, f)
	if err != nil {
		return err
	}
	jobs[ct.ID] = entryID
	return nil
}

// 删除定时任务
func RemoveJob(ct models.Crontab) {
	entryID, ok := jobs[ct.ID]
	if !ok {
		return
	}
	c.Remove(entryID)
	delete(jobs, ct.ID)
}

// 删除定时任务
func RemoveJobByCrontabID(crontabID int) {
	entryID, ok := jobs[crontabID]
	if !ok {
		return
	}
	c.Remove(entryID)
	delete(jobs, crontabID)
}
