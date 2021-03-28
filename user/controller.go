package user

func ArchiveUser(ip string) (uint, error) {
	user, error := readUser(ip)
	if error != nil {
		newUser := &User{
			IP: ip,
		}
		newUser.ID, error = createUser(newUser)
		return newUser.ID, error
	}
	return user.ID, error
}
