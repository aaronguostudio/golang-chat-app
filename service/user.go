package service

import (
	"../model"
	"errors"
	"fmt"
	"math/rand"
	"../util"
	"time"
)

type UserService struct {

}

func (s *UserService) Register (
	mobile,
	plainpwd,
	nickname,
	avatar,
	sex string) (user model.User, err error) {

		tmp := model.User{}

		// check if mobile exists
		_, err = DbEngin.Where("mobile=? ", mobile).Get(&tmp)

		if err != nil { return tmp, err }

		if tmp.Id > 0 {
			return tmp, errors.New("this mobile number has been used")
		}

		tmp.Mobile = mobile
		tmp.Avatar = avatar
		tmp.Nickname = nickname
		tmp.Sex = sex
		tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
		tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
		tmp.Createat = time.Now()

		_, err = DbEngin.InsertOne(&tmp)

		return user, err
}


func (s *UserService) Login (
	mobile,
	plainpwd string) (user model.User, err error) {

	return user, err
}
