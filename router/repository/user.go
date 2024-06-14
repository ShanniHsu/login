package repository

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"login/model"
	"time"
)

//放撈取對資料庫操作的方法
type UserRepository interface {
	GetByID(ID int64) (user *model.User, err error)
	GetByAccount(account string) (user *model.User, err error)
	GetByEmail(email string) (user *model.User, err error)
	GetByToken(token string) (user *model.User, err error)
	GetByTempToken(tempToken string) (user *model.User, err error)
	Create(user *model.User) (err error)
	Update(user *model.User, newdata map[string]interface{}) (err error)
	GetRedis(key string) (value string, err error)
	SetRedis(key string, value string, expiration time.Duration) (err error)
}

//DB連線資訊
type userRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return userRepository{
		DB:    db,
		Redis: redis,
	}
}

//userRepository結構體的函式
func (r userRepository) GetByID(ID int64) (user *model.User, err error) {
	user = new(model.User)
	err = r.DB.Where("id", ID).First(user).Error
	return
}

func (r userRepository) GetByAccount(account string) (user *model.User, err error) {
	user = new(model.User)
	err = r.DB.Where("account", account).First(user).Error
	return
}

func (r userRepository) GetByEmail(email string) (user *model.User, err error) {
	user = new(model.User)
	err = r.DB.Where("email", email).First(user).Error
	return
}

func (r userRepository) GetByToken(token string) (user *model.User, err error) {
	user = new(model.User)
	err = r.DB.Where("token", token).First(user).Error
	return
}

func (r userRepository) GetByTempToken(tempToken string) (user *model.User, err error) {
	user = new(model.User)
	err = r.DB.Where("temp_token", tempToken).First(user).Error
	return
}

func (r userRepository) Create(user *model.User) (err error) {
	err = r.DB.Create(user).Error
	if err != nil {
		err = errors.New("User Create Error!")
	}
	return
}

func (r userRepository) Update(user *model.User, newdata map[string]interface{}) (err error) {
	err = r.DB.Model(user).Updates(newdata).Error
	if err != nil {
		err = errors.New("User Update Error!")
	}
	return
}

func (r userRepository) GetRedis(key string) (value string, err error) {
	ctx := context.TODO()
	return r.Redis.Get(ctx, key).Result()
}

func (r userRepository) SetRedis(key string, value string, expiration time.Duration) (err error) {
	ctx := context.TODO()
	return r.Redis.Set(ctx, key, value, expiration).Err()
}
