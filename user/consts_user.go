package user

const (
	InsertUser = `insert into public."user" (login, name, last_name, email, phone, hash_password)
values (@login, @name, @last_name, @email, @phone, @hash_password) returning (user_id)`
	SelectUser = `select user_id, login, name, last_name, email, phone from public."user" where login=@login`
)
