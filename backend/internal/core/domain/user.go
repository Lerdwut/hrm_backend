package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceInfo struct {
	AttendanceID primitive.ObjectID `bson:"attendance_id,omitempty" json:"attendance_id,omitempty"`
	Timestamp    time.Time          `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Type         string             `bson:"type,omitempty" json:"type,omitempty"` // "checkin" หรือ "checkout"
	// Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	// Note string        `bson:"note,omitempty" json:"note,omitempty"`
	GPS []GPSLocation `bson:"gps,omitempty" json:"gps,omitempty"`
}

type GPSLocation struct {
	Latitude     float64 `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude    float64 `bson:"longitude,omitempty" json:"longitude,omitempty"`
	LocationName string  `bson:"location_name,omitempty" json:"location_name,omitempty"`
	DeviceInfo   string  `bson:"device_info,omitempty" json:"device_info,omitempty"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password,omitempty" json:"-"` // optional สำหรับ OAuth
	GoogleID   string             `bson:"google_id,omitempty" json:"-"`
	Avatar     string             `bson:"avatar,omitempty" json:"avatar"`
	Provider   string             `bson:"provider" json:"provider"` // "local", "google"
	IsVerified bool               `bson:"is_verified" json:"is_verified"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`

	Attendance []AttendanceInfo `bson:"attendance,omitempty" json:"attendance,omitempty"`
}
