package goFiles

func invalidUsername() error {
	return errors.New("this username is invalid, please register or log in with a valid one")
}
