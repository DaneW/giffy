package util

import (
	"os"
	"strconv"
	"testing"
)

func TestEnv(t *testing.T) {
	in, out := "baz", "bar"

	os.Setenv("ENV_STRING", out)

	if got := Env("ENV_STRING", in); got != out {
		t.Errorf(`String("ENV_STRING", "%v") = %v, want %v`, in, got, out)
	}
}

func TestEnvDefault(t *testing.T) {
	in, out := "baz", "baz"

	if got := Env("ENV_STRING_DEFAULT", in); got != out {
		t.Errorf(`String("ENV_STRING_DEFAULT", "%v") = %v, want %v`, in, got, out)
	}
}

func TestEnvInt(t *testing.T) {
	in, out := 1, 2

	os.Setenv("ENV_INT", strconv.Itoa(out))

	if got := EnvInt("ENV_INT", in); got != out {
		t.Errorf(`Int("ENV_INT", %v) = %v, want %v`, in, got, out)
	}
}

func TestEnvIntDefault(t *testing.T) {
	in, out := 3, 3

	if got := EnvInt("ENV_INT_DEFAULT", in); got != out {
		t.Errorf(`Int("ENV_INT_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}
