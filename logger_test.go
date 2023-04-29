package hsocks

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		defer func() {
			err := os.Remove("testdata/test.log")
			require.NoError(t, err)
		}()

		lg, err := newLogger(os.Stdout, "testdata/test.log")
		require.NoError(t, err)

		lg.Info("info log")
		lg.Infof("%s", "infof log")

		lg.Warning("warning log")
		lg.Warningf("%s", "warningf log")

		lg.Error("error log")
		lg.Errorf("%s", "errorf log")

		lg.Fatal("test func", "fatal log")
		lg.Fatalf("test func", "%s", "fatalf log")

		err = lg.Close()
		require.NoError(t, err)
	})

	t.Run("no file path", func(t *testing.T) {
		lg, err := newLogger(os.Stdout, "")
		require.NoError(t, err)

		lg.Info("info log")
		lg.Infof("%s", "infof log")

		lg.Warning("warning log")
		lg.Warningf("%s", "warningf log")

		lg.Error("error log")
		lg.Errorf("%s", "errorf log")

		lg.Fatal("test func", "fatal log")
		lg.Fatalf("test func", "%s", "fatalf log")

		err = lg.Close()
		require.NoError(t, err)
	})

	t.Run("failed to make directory", func(t *testing.T) {
		gomonkey.ApplyFuncReturn(os.MkdirAll).Reset()
	})

	t.Run("failed to open file", func(t *testing.T) {

	})

	t.Run("failed to close file", func(t *testing.T) {

	})
}
