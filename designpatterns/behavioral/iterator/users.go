package iterator

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) HasNext() bool {
	return u.index < len(u.users)
}

func (u *UserIterator) Next() *User {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
