package user

func GetList(repo UserRepository) ([]User, error) {
	return repo.GetList()
}
