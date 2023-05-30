package entity

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

// TableName is a gorm hook.
func (r Row) TableName() string {
	//return "Data_" + strconv.FormatInt(r.OrgID, 10)
	return "Data"
}

// BeforeCreate is a gorm hook.
func (r *Row) BeforeCreate(tx *gorm.DB) (err error) {
	if err = r.basicValidate(); err != nil {
		return
	}

	if strings.TrimSpace(r.CUser) == "" {
		// TODO cuser not empty?
	}

	now := time.Now()
	if r.CTime.IsZero() {
		r.CTime = now
	}
	if r.MTime.IsZero() {
		r.MTime = now
	}
	return
}

// BeforeUpdate is a gorm hook.
func (r *Row) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = r.basicValidate(); err != nil {
		return
	}

	if r.ID < 1 {
		return ErrRowEmptyID
	}

	if strings.TrimSpace(r.MUser) == "" {
		// TODO muser not empty?
	}

	if r.MTime.IsZero() {
		r.MTime = time.Now()
	}
	return
}

// BeforeDelete is hook of gorm.
func (r *Row) BeforeDelete(tx *gorm.DB) (err error) {
	if r.ModelID < 1 {
		// a row must belong to a model
		return ErrRowEmptyModelID
	}

	return
}

func (r *Row) basicValidate() error {
	if r.orgID < 1 {
		// orgID is the sharding policy: each org has all model data inside a single data table
		return ErrRowEmptyOrgID
	}

	if r.ModelID < 1 {
		// a row must belong to a model
		return ErrRowEmptyModelID
	}

	return nil
}
