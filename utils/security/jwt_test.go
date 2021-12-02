package security

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	//a:=&auth{
	//	"ann",[]string{"user"},
	//}
	//token,err:=GenerateToken(a)
	//if err != nil {
	//	return
	//}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzUyMzg4NTUsImlzcyI6Im1hcmtldHMiLCJzdWIiOiJrdGVzdDEifQ.6K2WVgshSaF_dN3LLRhQHeoQLJc-Nyz92FhkYw_Y7gc"
	_, err := ExtractName(token)
	if err != nil {
		return
	}

}

type auth struct {
	Name        string
	Authorities []string
}

func (a *auth) GetName() string {
	return a.Name
}

func (a *auth) GetAuthorities() []string {
	return a.Authorities
}
