package milestones

import (
	"sync"

	"github.com/nidhey27/project-service/dbs"
	"github.com/nidhey27/project-service/models"
	"gorm.io/gorm"
)

type MilestoneService struct {
	db *gorm.DB
}

var milestoneService *MilestoneService
var initOnce sync.Once

func GetService() *MilestoneService {
	initOnce.Do(initService)
	return milestoneService
}

func initService() {
	db := dbs.GetDB()
	milestoneService = NewMilestoneService(db)
}

func NewMilestoneService(db *gorm.DB) *MilestoneService {
	return &MilestoneService{
		db: db,
	}
}

/**
 * Milestone operations
**/

func (ms *MilestoneService) GetAllMilestones() ([]models.Milestone, error) {
	var milestones []models.Milestone
	result := ms.db.Find(&milestones)
	if result.Error != nil {
		return nil, result.Error
	}
	return milestones, nil
}

func (ms *MilestoneService) CreateMilestone(milestone *models.Milestone) error {
	result := ms.db.Create(milestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ms *MilestoneService) EditMilestone(id uint, milestone models.Milestone) error {
	dbMilestone := models.Milestone{}
	dbMilestone.ID = id
	result := ms.db.Model(&dbMilestone).Updates(milestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ms *MilestoneService) DeleteMilestone(id uint) error {
	dbMilestone := models.Milestone{}
	dbMilestone.ID = id
	result := ms.db.Delete(&dbMilestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
