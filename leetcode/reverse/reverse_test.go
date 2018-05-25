package reverse

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T){
	require:=require.New(t)
	require.Equal(97531,reverse(13579))
	require.Equal(-97531,reverse(-13579))
}