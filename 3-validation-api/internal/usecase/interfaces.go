package usecase

type SendRepo interface {
	GetEmailByHash(hash string) (string, error)
	SaveEmailHash(email, hash string) error
}
