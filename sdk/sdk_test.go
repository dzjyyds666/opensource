package sdk

import "testing"

func TestJwt(t *testing.T) {
	data := "爱你有"
	secretKey := "secretKey"
	token, err := CreateJwtToken(secretKey, []byte(data))
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	claims, err := ParseJwtToken(secretKey, token)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims)

}
