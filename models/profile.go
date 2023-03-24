package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	UserID            uint
	IsFollowingUserID uint
}

func GetProfileByUserId(uid uint) (Profile, error) {
	var p Profile

	if err := DB.First(&p, uid).Error; err != nil {
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

	return true, nil
}

func (u *User) FollowUser(uid uint) (*Profile, error) {
	p := Profile{UserID: u.ID, IsFollowingUserID: uid}

	if err := DB.Create(&p).Error; err != nil {
		return nil, errors.New("User not found!")
	}

	return &p, nil
}
