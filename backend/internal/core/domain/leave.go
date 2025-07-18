package domain

import (
	"time"
)

// LeaveStatus represents the status of a leave request
type LeaveStatus string

const (
	Pending  LeaveStatus = "pending"  // Pending represents a pending leave request
	Approved LeaveStatus = "approved" // Approved represents an approved leave request
	Rejected LeaveStatus = "rejected" // Rejected represents a rejected leave request
)

// Leave represents a leave request
type Leave struct {
	ID         uint        `json:"id" gorm:"primaryKey" example:"1"`
	EmployeeID uint        `json:"employee_id" example:"123"`
	Reason     string      `json:"reason" example:"Family vacation"`
	FromDate   time.Time   `json:"from_date" example:"2024-01-15T00:00:00Z"`
	ToDate     time.Time   `json:"to_date" example:"2024-01-20T00:00:00Z"`
	Status     LeaveStatus `json:"status" example:"pending"`
	CreatedAt  time.Time   `json:"created_at" example:"2024-01-01T00:00:00Z"`
}
