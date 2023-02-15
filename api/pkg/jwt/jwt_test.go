package jwt

import (
	"fmt"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"testing"
)

func TestGnerateToken(t *testing.T) {
	token, err := GnerateToken("111")
	if err != nil {
		log.Info(err)
	}
	log.Info(token)
}

func TestParseToken(t *testing.T) {
	token, err := GnerateToken("111")
	if err != nil {
		log.Info("generate fail...")
	}
	fmt.Println(token, err)
	claim, err := ParseToken(token)
	if err != nil {
		log.Warn("parse fail...")
	}
	fmt.Println(claim["Id"], err)
}
