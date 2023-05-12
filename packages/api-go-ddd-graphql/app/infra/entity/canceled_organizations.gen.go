// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameCanceledOrganization = "canceled_organizations"

// CanceledOrganization mapped from table <canceled_organizations>
type CanceledOrganization struct {
	OrganizationID string    `gorm:"column:organization_id;primaryKey" json:"organization_id"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName CanceledOrganization's table name
func (*CanceledOrganization) TableName() string {
	return TableNameCanceledOrganization
}
