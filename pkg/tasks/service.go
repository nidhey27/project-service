package tasks

import (
	"sync"

	"github.com/nidhey27/project-service/dbs"
	"github.com/nidhey27/project-service/models"
	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

var taskService *TaskService
var initOnce sync.Once

func GetService() *TaskService {
	initOnce.Do(initService)
	return taskService
}

func initService() {
	db := dbs.GetDB()
	taskService = NewTaskService(db)
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

/**
 * Task operations
**/

func (ts *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := ts.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (ts *TaskService) CreateTask(task *models.Task) error {
	result := ts.db.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ts *TaskService) EditTask(id uint, task models.Task) error {
	dbTask := models.Task{}
	dbTask.ID = id
	result := ts.db.Model(&dbTask).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ts *TaskService) DeleteTask(id uint) error {
	dbTask := models.Task{}
	dbTask.ID = id
	result := ts.db.Delete(&dbTask)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
