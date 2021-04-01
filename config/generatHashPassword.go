package config

// "golang.org/x/crypto/bcrypt"

import (
	"fmt"

	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

func GeneratePasswordFunc(password string) (string, error) {

	c := sha512_crypt.New()
	hash, err := c.Generate([]byte(password), []byte("$6$usesomesillystringforsalt"))
	if err != nil {
		return "", err

	}

	fmt.Println(hash)
	return hash, nil

}
