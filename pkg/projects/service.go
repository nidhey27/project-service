package projects

import (
	"sync"

	"github.com/nidhey27/project-service/dbs"
	"github.com/nidhey27/project-service/models"
	"gorm.io/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

var projectService *ProjectService
var initOnce sync.Once

func GetService() *ProjectService {
	initOnce.Do(initService)
	return projectService
}

func initService() {
	db := dbs.GetDB()
	projectService = NewProjectService(db)
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

/**
 * Project operations
**/

func (ps *ProjectService) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	result := ps.db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (ps *ProjectService) CreateProject(project *models.Project) error {
	result := ps.db.Create(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ps *ProjectService) EditProject(id uint, project models.Project) error {
	dbProject := models.Project{}
	dbProject.ID = id
	result := ps.db.Model(&dbProject).Updates(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ps *ProjectService) DeleteProject(id uint) error {
	dbProject := models.Project{}
	dbProject.ID = id
	result := ps.db.Delete(&dbProject)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
