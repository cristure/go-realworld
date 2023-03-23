package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	Following    User
	FollowingID  uint
	FollowedBy   User
	FollowedByID uint
}

func GetProfileByUserId(uid uint) (Profile, error) {

	var p Profile

	if err := DB.First(&p, uid).Error; err != nil {
		return p, errors.New("Profile not found!")
	}

	return p, nil

}
