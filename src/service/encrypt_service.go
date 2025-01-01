package service

import "golang.org/x/crypto/bcrypt"

// EncryptServiceはパスワード関連の機能を提供します
type encryptService struct{}

// NewEncryptServiceは新しいEncryptServiceを作成します
func NewEncryptService() EncryptService {
	return &encryptService{}
}

// ComparePasswordは、ハッシュ化されたパスワードと平文パスワードを比較します
func (p *encryptService) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// HashPasswordは平文パスワードをハッシュ化します
func (p *encryptService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
