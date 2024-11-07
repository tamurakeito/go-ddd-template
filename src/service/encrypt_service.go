package service

import "golang.org/x/crypto/bcrypt"

// PasswordServiceはパスワード関連の機能を提供します
type encryptService struct{}

// NewPasswordServiceは新しいPasswordServiceを作成します
func NewEncryptService() EncryptService {
	return &encryptService{}
}

// ComparePasswordは、ハッシュ化されたパスワードと平文パスワードを比較します
func (p *encryptService) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
