package infra

import "github.com/khairozzaman91/JobPortal-Backend/domain"



func (r *UserRepository) Store(u domain.User) (domain.User, error) {
	u.ID = uint(len(r.userList) + 1)
	r.userList = append(r.userList, u)

	return u, nil
}
func (r *UserRepository) List() ([]domain.User, error) {
	return r.userList, nil
}

func (r *UserRepository) Get(userID int) (*domain.User, error) {
	for _, user := range r.userList {
		if user.ID == uint(userID) {
			return &user, nil
		}
	}

	return nil, nil
}

func (r *UserRepository) Update(user domain.User) (domain.User, error) {
	for idx, u := range r.userList {
		if u.ID == uint(user.ID) {
			r.userList[idx] = user
			return user, nil
		}
	}

	return domain.User{}, nil
}


func (r *UserRepository) Delete(userID uint) error {
	var tempList []domain.User

	for _, user := range r.userList {
		if user.ID != userID {
			tempList = append(tempList, user)
		}
	}

	r.userList = tempList

	return nil
}

func (r *UserRepository) DeleteAll() error {
	r.userList = nil

	return nil
}