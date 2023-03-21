package models

import "errors"

type Profile struct {
	User
	Following bool `json:"following"`
}

func GetProfileByUserId(uid uint) (Profile, error) {

	var p Profile

	if err := DB.First(&p, uid).Error; err != nil {
		return p, errors.New("Profile not found!")
	}

	return p, nil

}
