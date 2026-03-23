package persistence

import (
	"github.com/waiter/back/domain/entity"
	"gorm.io/gorm"
)

type FeedbackRepo struct {
	db *gorm.DB
}

func NewFeedbackRepo(db *gorm.DB) *FeedbackRepo {
	return &FeedbackRepo{db: db}
}

func (r *FeedbackRepo) Create(feedback *entity.Feedback) error {
	return r.db.Create(feedback).Error
}

func (r *FeedbackRepo) FindByTableID(tableID string) ([]entity.Feedback, error) {
	var feedbacks []entity.Feedback
	if err := r.db.Where("table_id = ?", tableID).Order("created_at DESC").Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}
