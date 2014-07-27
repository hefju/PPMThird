package models
import (
	"time"
)

type Task struct {
	Id int64
	Content string `xorm:"VARCHAR(50)"`
	Finish bool
	FinishTime time.Time
	Version int
}
