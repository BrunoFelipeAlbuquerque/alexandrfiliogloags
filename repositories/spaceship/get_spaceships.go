package spaceshiprepositories

import (
	"errors"
	"time"

	"imperial-fleet/dto"
	"imperial-fleet/models"

	"gorm.io/gorm"
)

func GetSpaceships(db *gorm.DB, filter dto.GetSpaceshipFilterDTO) ([]models.Spaceship, int64, error) {
	if err := validateFilters(filter); err != nil {
		return nil, 0, err
	}

	query := db.Model(&models.Spaceship{}).Preload("Armament")
	query = applyTextFilters(query, filter)
	query = applyNumericFilters(query, filter)
	query, err := applyDateFilters(query, filter)
	if err != nil {
		return nil, 0, err
	}

	query = query.Order("id")

	limit := filter.Limit
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	page := filter.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var spaceships []models.Spaceship
	if err := query.Limit(limit).Offset(offset).Find(&spaceships).Error; err != nil {
		return nil, 0, err
	}

	return spaceships, total, nil
}

func validateFilters(filter dto.GetSpaceshipFilterDTO) error {
	if filter.Value != nil && (filter.ValueMin != nil || filter.ValueMax != nil) {
		return errors.New("use either 'value' or 'value_min/value_max', not both")
	}
	if filter.CreatedAt != "" && (filter.CreatedAtMin != "" || filter.CreatedAtMax != "") {
		return errors.New("use either 'created_at' or 'created_at_min/created_at_max', not both")
	}
	if filter.UpdatedAt != "" && (filter.UpdatedAtMin != "" || filter.UpdatedAtMax != "") {
		return errors.New("use either 'updated_at' or 'updated_at_min/updated_at_max', not both")
	}
	if filter.Status != "" && !models.IsValidStatus(filter.Status) {
		return errors.New("invalid status value")
	}
	return nil
}

func applyTextFilters(query *gorm.DB, filter dto.GetSpaceshipFilterDTO) *gorm.DB {
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Class != "" {
		query = query.Where("class ILIKE ?", "%"+filter.Class+"%")
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	return query
}

func applyNumericFilters(query *gorm.DB, filter dto.GetSpaceshipFilterDTO) *gorm.DB {
	if filter.Crew != nil {
		query = query.Where("crew = ?", *filter.Crew)
	}
	if filter.Value != nil {
		query = query.Where("value = ?", *filter.Value)
	}
	if filter.ValueMin != nil {
		query = query.Where("value >= ?", *filter.ValueMin)
	}
	if filter.ValueMax != nil {
		query = query.Where("value <= ?", *filter.ValueMax)
	}
	return query
}

func applyDateFilters(query *gorm.DB, filter dto.GetSpaceshipFilterDTO) (*gorm.DB, error) {
	pattern := "2006-01-02"

	if filter.CreatedAt != "" {
		t, err := time.Parse(pattern, filter.CreatedAt)
		if err != nil {
			return nil, errors.New("invalid format for created_at (expected yyyy-mm-dd)")
		}
		query = query.Where("DATE(created_at) = ?", t)
	}
	if filter.CreatedAtMin != "" {
		t, err := time.Parse(pattern, filter.CreatedAtMin)
		if err != nil {
			return nil, errors.New("invalid format for created_at_min (expected yyyy-mm-dd)")
		}
		query = query.Where("created_at >= ?", t)
	}
	if filter.CreatedAtMax != "" {
		t, err := time.Parse(pattern, filter.CreatedAtMax)
		if err != nil {
			return nil, errors.New("invalid format for created_at_max (expected yyyy-mm-dd)")
		}
		query = query.Where("created_at <= ?", t)
	}

	if filter.UpdatedAt != "" {
		t, err := time.Parse(pattern, filter.UpdatedAt)
		if err != nil {
			return nil, errors.New("invalid format for updated_at (expected yyyy-mm-dd)")
		}
		query = query.Where("DATE(updated_at) = ?", t)
	}
	if filter.UpdatedAtMin != "" {
		t, err := time.Parse(pattern, filter.UpdatedAtMin)
		if err != nil {
			return nil, errors.New("invalid format for updated_at_min (expected yyyy-mm-dd)")
		}
		query = query.Where("updated_at >= ?", t)
	}
	if filter.UpdatedAtMax != "" {
		t, err := time.Parse(pattern, filter.UpdatedAtMax)
		if err != nil {
			return nil, errors.New("invalid format for updated_at_max (expected yyyy-mm-dd)")
		}
		query = query.Where("updated_at <= ?", t)
	}

	return query, nil
}
