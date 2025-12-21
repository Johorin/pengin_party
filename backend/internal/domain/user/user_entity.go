package user

import (
	user "pengin_party/internal/domain/user/valueobjects"
	// "pengin_party/internal/infrastructure/dbmodel"
)

type UserEntity struct {
	id    *uint
	name  string
	email user.Email
	uid   string
	uuid  string
}

// Getters
func (u *UserEntity) ID()    *uint      { return u.id }
func (u *UserEntity) Name()  string     { return u.name }
func (u *UserEntity) Email() user.Email { return u.email }
func (u *UserEntity) UID()   string     { return u.uid }
func (u *UserEntity) UUID()  string     { return u.uuid }

func NewUserEntity(
	name, email, uid, uuid string,
) (*UserEntity, error) {
	emailVO, err := user.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &UserEntity{
		id: nil,
		name: name,
		email: emailVO,
		uid: uid,
		uuid: uuid,
	}, nil
}

// // DBモデルからユーザーエンティティを作成
// func FromDBModel(u *dbmodel.User) (*UserEntity, error) {
// 	//
// }

// // ユーザーエンティティを再構築
// func (u *UserEntity) Reconstruct(data ReconstructData) (err error) {
// 	//
// }

// // ユーザーエンティティを再構成
// func ReconstituteUserEntity(
// 	id uint,
// 	name string,
// 	email string,
// 	uid string,
// 	uuid string,
// ) { 
// 	//
// }