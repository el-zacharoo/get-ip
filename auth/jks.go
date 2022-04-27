package auth

import (
	"io/ioutil"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
)

func JKS(path string) (jwk.Set, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return jwk.Parse(byt)
}
