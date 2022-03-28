package repository

import "golang-template/entity"

func (r *repository) ListUser() ([]entity.User, error) {
	results := []entity.User{}

	users, err := r.db.Query("select * from users")

	if err != nil {
		return results, err
	}

	defer users.Close()

	for users.Next() {
		var each entity.User

		err = users.Scan(&each.ID, &each.Name, &each.Email, each.Phone)

		if err != nil {
			return results, err
		}

		results = append(results, each)
	}

	return results, nil
}
