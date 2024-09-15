package user

import "gorm.io/gorm"

func GetUserByToken(db *gorm.DB, token string) (*User, error) {
	var userRes User
	if err := db.Where("token = ?", token).First(&userRes).Error; err != nil {
		return nil, err
	}

	return &userRes, nil
}
