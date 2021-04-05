package logoutput

import (
	"testing"
)

func TestHasError(t *testing.T) {
	boolReturn := HasError("... error ...")
	if boolReturn != true {
		t.Fatalf(`HasError("... error ...") want true, error`)
	}

	boolReturn = HasError("... stacktrace ...")
	if boolReturn != true {
		t.Fatalf(`HasError("... error ...") want true, error`)
	}

	boolReturn = HasError("... something else ...")
	if boolReturn {
		t.Fatalf(`HasError("... error ...") want false, error`)
	}
}

func TestParseMessageByLevel(t *testing.T) {
	i := 1
	pass := &i
	type args struct {
		i        int
		log      string
		loglevel *int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "a", args: args{1, "test", pass}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
