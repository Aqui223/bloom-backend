package UserApp

func (u *UserApp) IsUserWithEmailExists(email string) (bool, error) {
	_, err := u.users.GetByEmail(email)
	if err != nil {
		return false, err
	}

	return true, nil
}
