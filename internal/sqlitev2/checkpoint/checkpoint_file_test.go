package checkpoint

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkcom/statshouse/internal/sqlitev2/checkpoint/gen2/tlsqlite"
	"pgregory.net/rapid"
)

// TODO add more tests

func Test_encode(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		commitOffset := rapid.Int64().Draw(t, "commit_offset")
		m := tlsqlite.MetainfoBytes{Offset: commitOffset}
		metainfo, err := decode(encode(m, nil))
		require.NoError(t, err)
		require.Equal(t, commitOffset, metainfo.Offset)
	})
}
