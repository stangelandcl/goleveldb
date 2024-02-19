package leveldb

import (
	"testing"

	"github.com/stangelandcl/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
