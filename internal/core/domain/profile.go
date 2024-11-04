package domain

// Profile is an entity representing a user's profile.
type Profile struct {
	UserID            uint `gorm:"primary_key"`
	IsFollowingUserID uint `gorm:"primary_key"`
}
