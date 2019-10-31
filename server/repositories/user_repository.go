package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/model"
)

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (this *userRepository) Get(db *gorm.DB, id int64) *model.User {
	ret := &model.User{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userRepository) Take(db *gorm.DB, where ...interface{}) *model.User {
	ret := &model.User{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *userRepository) QueryCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.User, err error) {
	err = cnd.Exec(db).Find(&list).Error
	return
}

func (this *userRepository) Query(db *gorm.DB, params *simple.QueryParams) (list []model.User, paging *simple.Paging) {
	params.StartQuery(db).Find(&list)
	params.StartCount(db).Model(&model.User{}).Count(&params.Paging.Total)
	paging = params.Paging
	return
}

func (this *userRepository) Create(db *gorm.DB, t *model.User) (err error) {
	err = db.Create(t).Error
	return
}

func (this *userRepository) Update(db *gorm.DB, t *model.User) (err error) {
	err = db.Save(t).Error
	return
}

func (this *userRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (this *userRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (this *userRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.User{}, "id = ?", id)
}

func (this *userRepository) GetByEmail(db *gorm.DB, email string) *model.User {
	return this.Take(db, "email = ?", email)
}

func (this *userRepository) GetByUsername(db *gorm.DB, username string) *model.User {
	return this.Take(db, "username = ?", username)
}
