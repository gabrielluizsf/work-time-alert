package store

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func newTestStore(t *testing.T) (Store, string) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	st := NewLocal()
	assert.NotNil(t, st)
	return st, filepath.Join(home, storeDirName, storeFileName)
}

func TestLoadOnAFirstRunIsEmpty(t *testing.T) {
	st, _ := newTestStore(t)
	alerts, err := st.Load()
	assert.NoError(t, err)
	assert.Equal(t, len(alerts), 0)
}

func TestSaveThenLoadRoundTrips(t *testing.T) {
	st, path := newTestStore(t)
	in := []Alert{{ID: 1, Name: "Foco", Minutes: 25}, {ID: 2, Name: "Pausa", Minutes: 5, Seconds: 30}}
	assert.NoError(t, st.Save(in...))

	out, err := st.Load()
	assert.NoError(t, err)
	assert.Equal(t, out, in)

	data, err := os.ReadFile(path)
	assert.NoError(t, err)
	assert.False(t, bytes.Contains(data, []byte("Foco")), "the file must be encrypted")
	info, err := os.Stat(path)
	assert.NoError(t, err)
	assert.Equal(t, info.Mode().Perm(), os.FileMode(0o600))
}

func TestLoadReportsATamperedFile(t *testing.T) {
	st, path := newTestStore(t)
	assert.NoError(t, st.Save(Alert{ID: 1, Name: "Foco", Minutes: 25}))
	assert.NoError(t, os.WriteFile(path, []byte("garbage that is not ciphertext"), 0o600))

	_, err := st.Load()
	assert.Error(t, err)
}
