package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.starlark.net/starlark"
)

func TestNewDag(t *testing.T) {
	files := []string{
		"starlark/related.star",
		"starlark/home.star",
	}
	for _, file := range files {

		g, err := NewDag(file)
		require.NoError(t, err)
		fmt.Println(g)
	}
}

func TestEdgesToDependencyMap(t *testing.T) {
	edges := starlark.NewList([]starlark.Value{
		starlark.Tuple{
			starlark.String("a"),
			starlark.String("b"),
		},
		starlark.Tuple{
			starlark.String("b"),
			starlark.String("c"),
		},
		starlark.Tuple{
			starlark.String("a"),
			starlark.String("c"),
		},
	})
	dependencyMap := EdgesToDependencyMap(edges)
	expected := map[string][]string{
		"b": {"a"},
		"c": {"b", "a"},
	}
	require.Equal(t, expected, dependencyMap)
}
