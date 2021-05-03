package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tdarci/prj-999/config"
)

func TestAdd(t *testing.T) {

	type Args struct {
		A int
		B int
	}
	type Expected struct {
		Result int
	}
	tests := []struct {
		Message  string
		Args     Args
		Expected Expected
	}{
		{
			Message: "zeros make zeros",
			Args: Args{
				A: 0,
				B: 0,
			},
			Expected: Expected{
				Result: 0,
			},
		},
		{
			Message: "positive integers work",
			Args: Args{
				A: 1,
				B: 2,
			},
			Expected: Expected{
				Result: 3,
			},
		},
		{
			Message: "negative integers work",
			Args: Args{
				A: 1,
				B: -10,
			},
			Expected: Expected{
				Result: -9,
			},
		},
	}

	cfg, err := config.NewLocal()
	require.NoError(t, err, "config must be created")
	eng := NewEngine(cfg)

	for _, tt := range tests {
		t.Run(tt.Message, func(t *testing.T) {
			assert.Equal(t, tt.Expected.Result, eng.Add(tt.Args.A, tt.Args.B), "should get expected result")
		})
	}

}
