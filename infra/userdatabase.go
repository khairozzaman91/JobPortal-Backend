package infra

import "github.com/khairozzaman91/JobPortal-Backend/domain"

var userList []domain.User

func UserStore(u domain.User) domain.User {
	u.ID = uint(len(userList) + 1)
	userList = append(userList, u)
	return u
}

func UserList() []domain.User {
	return userList
}

func UserGet(uId int) *domain.User {
	// TODO: Refactor to return original slice element using index.
	// Current implementation returns pointer to range variable.
	for _, user := range userList {
		if user.ID == uint(uId) {
			return &user
		}
	}
	return nil
}

func UserUpdate(user domain.User) {
	for idx, j := range userList {
		if j.ID == uint(user.ID) {
			userList[idx] = user
			return
		}
	}
}

func UserDelete(uId uint) {
	var tempList []domain.User
	for _, usr := range userList {
		if usr.ID != uId {
			tempList = append(tempList, usr)
		}
	}
	userList = tempList
}


func UserDeleteAll() {
	userList = nil
}
