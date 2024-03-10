package user

const (
	InsertUser = `insert into public."user" (login, name, last_name, email, phone, hash_password)
values (@login, @name, @last_name, @email, @phone, @hash_password) returning (user_id)`
	SelectUser  = `select user_id, login, name, last_name, email, phone from public."user" where login=@login`
	SelectUsers = `select user_id, login, name, last_name, email, phone from public."user"`
	UpdateUser  = `update public."user" set name = @name, last_name = @last_name, email = @email,
                         phone = @phone, hash_password = @hash_password  where login=@login
                         returning (user_id)`
)
