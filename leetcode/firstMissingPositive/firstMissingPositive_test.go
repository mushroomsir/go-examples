package firstMissingPositive


import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestFirstMissingPositive(t *testing.T){
	require:=require.New(t)
	
	require.Equal(2,firstMissingPositive([]int{3,4,-1,1}))
	require.Equal(3,firstMissingPositive([]int{1,2,0}))
}