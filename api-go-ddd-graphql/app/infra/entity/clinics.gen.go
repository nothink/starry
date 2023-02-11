// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameClinic = "clinics"

// Clinic mapped from table <clinics>
type Clinic struct {
	ID           string       `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt    time.Time    `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	ClinicDetail ClinicDetail `json:"clinic_detail"`
}

// TableName Clinic's table name
func (*Clinic) TableName() string {
	return TableNameClinic
}
