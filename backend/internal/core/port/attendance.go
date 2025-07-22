package port

import (
	"hr_management/internal/core/domain"
)

type AttendanceRepository interface {
	RecordCheckIn(info *domain.AttendanceInfo) (*domain.AttendanceInfo, error)
	RecordCheckOut(attendanceID string, checkoutTime int64, checkoutLocation *domain.GPSLocation) (*domain.AttendanceInfo, error)
	GetTodayAttendanceByUserID(id string) (*domain.AttendanceInfo, error)
}

type AttendanceService interface {
	CheckIn(id string, location *domain.GPSLocation) (*domain.AttendanceInfo, error)
	CheckOut(id string, location *domain.GPSLocation) (*domain.AttendanceInfo, error)
	IsWithinAllowedLocation(location *domain.GPSLocation, allowedLocation *domain.GPSLocation, radius float64) (bool, error)
	GetTodayAttendance(id string) (*domain.AttendanceInfo, error)
}
