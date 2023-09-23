package postgre

const (
	// TODO: replace * with fields
	queryGetUserByEmail = `select * from users where email = $1;`
)
