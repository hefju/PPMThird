package controller

import( "models"
 "time"
"log")

func GetTask() []models.Task{
	var task_list []models.Task  //任务列表
	task_list=make([]models.Task,0,10)

	var t models.Task     //struct初始化1
	t.Id=1
	t.Content="学习使用template 1"
	t.Finish=false
	t.Version=0

	t1:= models.Task{}    //struct初始化2
	t1.Id=2
	t1.Content="学习使用struct"
	t1.Finish=false
	t1.Version=0

	time,_:=time.Parse("2006-01-02", "2014-02-03")
	log.Println(time)
	t2:=models.Task{
		Id:3 ,     //struct初始化3
		Content:"学习使用time 1",
		Finish:true,
		FinishTime:time,
		Version:0,
	}

	task_list=append(task_list,t)
	task_list=append(task_list,t1)
	task_list=append(task_list,t2)

	return task_list
}

