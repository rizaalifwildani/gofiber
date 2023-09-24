package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository struct {
	db      *gorm.DB
	Preload []string
}

type FilterType struct {
	Key   string
	Value interface{}
}

type FindType struct {
	Model   interface{}
	Filters []FilterType
	OrderBy string
}

func NewBaseRepository(db *gorm.DB, preload []string) *BaseRepository {
	return &BaseRepository{db: db, Preload: preload}
}

func (r *BaseRepository) Count(model interface{}) error {
	var totalData int64
	err := r.db.Model(&model).Count(&totalData).Error
	return err
}

func (r *BaseRepository) Create(model interface{}) error {
	err := r.db.Create(model).Error
	return err
}

func (r *BaseRepository) Find(model interface{}, filters []FilterType, orderBy string) error {
	query := r.db.Model(&model)
	for _, v := range filters {
		if v.Value != "" {
			query.Where("LOWER("+v.Key+")"+" ILIKE ?", fmt.Sprintf("%%%s%%", v.Value))
		}
	}

	if r.Preload != nil {
		for _, preload := range r.Preload {
			query.Preload(preload)
		}
	}

	if len(orderBy) > 0 {
		query.Order(fmt.Sprintf("%v DESC", orderBy))
	} else {
		query.Order("created_at DESC")
	}

	if err := query.Find(model).Error; err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository) FindOne(model interface{}, id string) error {
	query := r.db.Model(&model)
	if r.Preload != nil {
		for _, preload := range r.Preload {
			query.Preload(preload)
		}
	}
	if err := query.First(model, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository) Update(model interface{}, id string) error {
	err := r.db.Model(&model).Where("id = ?", id).Updates(model).Error
	return err
}

func (r *BaseRepository) UpdateAssociation(model interface{}, associationName string, data interface{}) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(model).Association(associationName).Unscoped().Clear(); err != nil {
			return err
		}

		if err := tx.Create(data).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *BaseRepository) Delete(model interface{}, ids []uuid.UUID) error {
	err := r.db.Unscoped().Delete(&model, ids).Error
	return err
}
