package jwt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JwtTestSuite struct {
	suite.Suite
}

func (suite *JwtTestSuite) SetupSuite() {
	fmt.Println(">>> Starting JWT test suite")
}

func (suite *JwtTestSuite) TearDownSuite() {
	fmt.Println("<<< Finish JWT test suite")
}

func (suite *JwtTestSuite) TestCreateJwtToken() {
	infos := make(map[string]string, 0)
	infos["username"] = "jack"
	tokenString, err := GenerateToken(infos, 3)
	assert.Nil(suite.T(), err)

	headers := "Bearer " + tokenString

	token, err := GetTokenFromHeader(headers)
	assert.Nil(suite.T(), err)
	customeInfos, _ := ParseCustomInfosFromToken(token)
	assert.Equal(suite.T(), "jack", customeInfos.Infos["username"])
}

func TestJwtSuite(t *testing.T) {
	suite.Run(t, new(JwtTestSuite))
}
