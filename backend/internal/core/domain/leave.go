package domain

import (
	"time"
)

type LeaveStatus string

const (
	Pending  LeaveStatus = "pending"
	Approved LeaveStatus = "approved"
	Rejected LeaveStatus = "rejected"
)

type Leave struct {
	ID         uint        `json:"id" bson:"_id,omitempty" gorm:"primaryKey" example:"1"`
	EmployeeID uint        `json:"employee_id" bson:"employee_id" example:"123"`
	Reason     string      `json:"reason" bson:"reason" example:"Family vacation"`
	FromDate   time.Time   `json:"from_date" bson:"from_date" example:"2024-01-15T00:00:00Z"`
	ToDate     time.Time   `json:"to_date" bson:"to_date" example:"2024-01-20T00:00:00Z"`
	Status     LeaveStatus `json:"status" bson:"status" example:"pending"`
	CreatedAt  time.Time   `json:"created_at" bson:"created_at" example:"2024-01-01T00:00:00Z"`
}
