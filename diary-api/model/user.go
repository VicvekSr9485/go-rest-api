package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"size:255;not null;unique" json:"username"`
    Password string `gorm:"size:255;not null;" json:"-"`
    Entries  []Entry
}

func (user *User) save() (*User, error) {
    err := database.Database.Create(&user).Error

    if err != nil {
        return &User{}, error
    }
    return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
    PasswordHash err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(passwordHash)
    user.Username = html.EscapeString(strings.TrimSpace(user.Username))
    return nil
}
