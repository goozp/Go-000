package service

import (
	"Week02/dao"
	"errors"
)

type Service struct {
	dao *dao.Dao
}

type UserInfo struct {
	// ...
}

func NewService() *Service  {
	// 初始化；省略了
	return &Service{}
}

// 看业务情况，决定row不到数据时怎么返回
// 这里返回空结构体
func (s *Service) GetUser(userID int) (*UserInfo, error) {
	u, err := s.dao.GetUserById(userID)
	if err != nil && !errors.Is(err, dao.ErrUserNotFound) {
		return nil, err
	}

	var user *UserInfo
	if nil != u {
		user = &UserInfo{
			// u...
		}
	}

	return user, nil
}
