package user

func Get(repo UserRepository, id string) (User, error) {
	return repo.Get(id)
}

func GetByEmail(repo UserRepository, email string) (User, error) {
	return repo.GetByEmail(email)
}
