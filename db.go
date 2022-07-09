package main
import "github.com/gomodule/redigo/redis"

type Account struct {
	userId string
	name string
	secret string
}

type AccountRepository interface {
	SetUpAccount(userId string, name string, secret string) error
	ListAccounts(userId string) (map[string]string, error)
}

type AccountRedisRepository struct {
	pool *redis.Pool
}

func (r *AccountRedisRepository) SetUpAccount(userId string, name string, secret string) error {
	conn := r.pool.Get()

	_, err := conn.Do("HMSET", userId, "name", name, "secret", secret)

	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRedisRepository) ListAccounts(userId string) (map[string]string, error) {
	conn := r.pool.Get()

	accounts, err := redis.StringMap(conn.Do("HGETALL", userId))

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func NewAccountRepository(pool *redis.Pool) AccountRepository {
	return &AccountRedisRepository{pool: pool}
}