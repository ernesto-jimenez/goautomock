package parseutil_test

import (
	"testing"

	"github.com/ernesto-jimenez/go-parse-utils"
	"github.com/stretchr/testify/require"
)

func TestPackageAST(t *testing.T) {
	pkg, err := parseutil.PackageAST(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil", pkg.Name)
}

func TestPackageTestAST(t *testing.T) {
	pkg, err := parseutil.PackageTestAST(project)
	require.Nil(t, err)
	require.Equal(t, "parseutil_test", pkg.Name)
}
