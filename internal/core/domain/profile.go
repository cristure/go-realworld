package domain

import (
	"errors"
	"reflect"
)

// Profile is an entity representing a user's profile.
type Profile struct {
	UserID            uint `gorm:"primary_key"`
	IsFollowingUserID uint `gorm:"primary_key"`
}

func GetProfileByUserId(uid uint) (*Profile, error) {
	var p *Profile

	if err := DB.First(p, uid).Error; err != nil {
		return p, errors.New("Profile not found!")
	}

	return p, nil
}

func (u *User) IsFollowing(uid uint) (bool, error) {
	var p Profile

	result := DB.Find(&p, Profile{UserID: u.ID, IsFollowingUserID: uid})

	if result.Error != nil {
		return false, result.Error
	}

	if reflect.DeepEqual(p, Profile{}) {
		return false, nil
	}

	return true, nil
}

func (u *User) FollowUser(uid uint) (*Profile, error) {
	p := Profile{UserID: u.ID, IsFollowingUserID: uid}

	if err := DB.Create(&p).Error; err != nil {
		return nil, errors.New("You already follow this user!")
	}

	return &p, nil
}

func (u *User) UnfollowUser(uid uint) error {
	p := Profile{UserID: u.ID, IsFollowingUserID: uid}

	if err := DB.Delete(&p).Error; err != nil {
		return errors.New("You do not follow this user!")
	}

	return nil
}
