package service

import (
	"hr_management/internal/core/domain"
	"hr_management/internal/core/port"
	"math"
	"time"
)

type AttendanceService struct {
	repo port.AttendanceRepository
}

func NewAttendanceService(repo port.AttendanceRepository) *AttendanceService {
	return &AttendanceService{repo}
}

func (as *AttendanceService) CheckIn(id string, location *domain.GPSLocation) (*domain.AttendanceInfo, error) {
	info := &domain.AttendanceInfo{
		Timestamp: time.Now(),
		Type:      "check_in",
		GPS:       []domain.GPSLocation{*location},
	}
	return as.repo.RecordCheckIn(info)
}

func (as *AttendanceService) CheckOut(id string, location *domain.GPSLocation) (*domain.AttendanceInfo, error) {
	info := &domain.AttendanceInfo{
		Timestamp: time.Now(),
		Type:      "check_out",
		GPS:       []domain.GPSLocation{*location},
	}
	return as.repo.RecordCheckOut(info.AttendanceID.Hex(), info.Timestamp.Unix(), location)
}

func (as *AttendanceService) IsWithinAllowedLocation(location *domain.GPSLocation, allowedLocation *domain.GPSLocation, radius float64) (bool, error) {
	const earthRadius = 6371000.0 // หน่วยเมตร

	lat1 := location.Latitude * math.Pi / 180
	lon1 := location.Longitude * math.Pi / 180
	lat2 := allowedLocation.Latitude * math.Pi / 180
	lon2 := allowedLocation.Longitude * math.Pi / 180

	// dlat = delta latitude
	// dlon = delta longitude

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance <= radius, nil
}

func (as *AttendanceService) GetTodayAttendance(id string) (*domain.AttendanceInfo, error) {
	return as.repo.GetTodayAttendanceByUserID(id)
}
