package nonil

import (
	"errors"
	"testing"
)

func ThirdParty(return_err bool) (int, error) {
	if return_err {
		return 0, errors.New("example error")
	}
	return 1, nil
}

func TestToResult(t *testing.T) {
	if ToResult(ThirdParty(false)).IsError() {
		t.Fatal("valid returned error")
	}
	if ToResult(ThirdParty(true)).IsValid() {
		t.Fatal("error returned valid")
	}
}
