package hash

import (
	"crypto/sha512"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
	"net/smtp"
	"validation/config"
	"validation/internal/usecase"

	emailAdapter "github.com/jordan-wright/email"
)

type HashService struct {
	config *config.Config
	repo   usecase.SendRepo
}

func NewHashService(cfg *config.Config, repo usecase.SendRepo) *HashService {
	return &HashService{config: cfg, repo: repo}
}

func (s HashService) VerifyEmail(hashToCompare string) (string, error) {
	email, err := s.repo.GetEmailByHash(hashToCompare)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (s HashService) SaveEmailHash(email string) error {
	hash := getHashString(email)
	if err := s.repo.SaveEmailHash(email, hash); err != nil {
		return err
	}
	log.Printf("Хэш для email %v сохранен\n", email)

	return sendVerifyEmail(
		hash,
		s.config.Login,
		email,
		s.config.SMTPHost,
		s.config.SMTPPort,
		s.config.Login,
		s.config.Password,
	)
}

func getHashString(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

func sendVerifyEmail(hash, from, to, smtpHost, smtpPort, smtpLogin, smtpPass string) error {
	e := emailAdapter.Email{
		To:      []string{to},
		From:    fmt.Sprintf("Maksim Dubakin <%s@yandex.ru>", from),
		Subject: "Подтверждение почты",
		HTML:    []byte(fmt.Sprintf("<a href='http://localhost:8088/verify/%s'>Подтвердить почту</a>", hash)),
	}

	tlsConfig := &tls.Config{
		ServerName:         smtpHost,
		InsecureSkipVerify: true,
	}

	if err := e.SendWithTLS(
		smtpHost+":"+smtpPort,
		smtp.PlainAuth("", smtpLogin, smtpPass, smtpHost),
		tlsConfig,
	); err != nil {
		return fmt.Errorf("ошибка отправки письма: %w", err)
	}

	return nil
}
