package models

type UserIn struct {
	Login        string `db:"login" json:"login"`
	Name         string `db:"name" json:"name"`
	LastName     string `db:"last_name" json:"last_name"`
	Email        string `db:"email" json:"email"`
	Phone        string `db:"phone" json:"phone"`
	HashPassword string `db:"hash_password" json:"password"`
}

type SelectUserIn struct {
	Login string `db:"login" json:"login"`
}

type UserOut struct {
	UserID   string `db:"user_id" json:"user_id"`
	Login    string `db:"login" json:"login"`
	Name     string `db:"name" json:"name"`
	LastName string `db:"last_name" json:"last_name"`
	Email    string `db:"email" json:"email"`
	Phone    string `db:"phone" json:"phone"`
}

type SelectUserOut struct {
	Data   *UserOut `db:"data" json:"data"`
	Status int      `db:"status" json:"status"`
}
