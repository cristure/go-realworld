package domain

import (
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/go-realworld/token"
)

// User is an entity representing a user.
type User struct {
	gorm.Model
	Username         string     `gorm:"size:255;not null;unique" json:"username"`
	Password         string     `gorm:"size:255;not null;" json:"password"`
	Email            string     `gorm:"size;255 not null;" json:"email"`
	Bio              string     `gorm:"size:255, not null;" json:"bio"`
	Image            string     `gorm:"size:255; nullable" json:"image"`
	FavoriteArticles []*Article `gorm:"many2many:favorite_article_user"`
	Articles         []Article
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func LoginAdmin(username, password string) (string, error) {
	if username == "admin" && password == "password" {
		return token.GenerateToken(0)
	}

	return "", errors.New("not admin")
}

func GetUserByID(uid uint) (User, error) {
	var u User

	if uid == 0 {
		return User{Username: "admin", Password: "password"}, nil
	}

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func GetUserByName(username string) (*User, error) {
	var user User

	if err := DB.Model(&User{}).Preload("FavoriteArticles").First(&user, "username = ?", username).Error; err != nil {
		return nil, errors.New("user not found!")
	}

	return &user, nil
}

func (u *User) UpdateUser(uid uint, newUser User) (*User, error) {
	if err := DB.First(&u, uid).Error; err != nil {
		return nil, errors.New("User was not found")
	}

	var err error
	if newUser.Password, err = MakePassword(u.Password); err != nil {
		return nil, err
	}

	DB.Save(newUser)
	return u, err
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	password, err := MakePassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = password

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func MakePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) FavoriteArticle(article *Article) error {
	var user User
	if err := DB.First(&user, u.ID).Error; err != nil {
		return err
	}

	user.FavoriteArticles = append(user.FavoriteArticles, article)
	DB.Save(user)
	return nil
}

func (u *User) FeedArticles() ([]*Article, error) {
	var user User
	list := make([]*Article, 0)

	articles, err := ListArticles()
	if err != nil {
		return nil, err
	}

	for _, a := range articles {
		if ok, _ := user.IsFollowing(a.UserID); !ok {
			return nil, err
		} else {
			list = append(list, a)
		}
	}

	return list, nil
}
