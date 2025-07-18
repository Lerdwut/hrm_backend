package domain

import (
	"time"
)

type LeaveStatus string

const (
	Pending   LeaveStatus = "pending"
	Approved  LeaveStatus = "approved"
	Rejected  LeaveStatus = "rejected"
)

type Leave struct {
    ID        uint        `json:"id"`
    EmployeeID uint       `json:"employee_id"`
    Reason    string      `json:"reason"`
    FromDate  time.Time   `json:"from_date"`
    ToDate    time.Time   `json:"to_date"`
    Status    LeaveStatus `json:"status"`
    CreatedAt time.Time   `json:"created_at"`
}