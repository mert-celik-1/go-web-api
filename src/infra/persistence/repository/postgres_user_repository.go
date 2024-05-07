package repository

const userFilterExp string = "username = ?"
const countFilterExp string = "count(*) > 0"

type PostgresUserRepository struct {
	//*BaseRepository
}
