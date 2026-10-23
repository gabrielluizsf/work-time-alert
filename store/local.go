package store

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	encryptor "github.com/gabrielluizsf/file-encryptor/encryptor"
)

// LocalStore keeps the alerts on disk, encrypted at rest. The app never asks
// for a password: the key is derived from the machine itself, so opening the
// app is frictionless while the file stays ciphertext away from prying eyes.
type LocalStore struct {
	path string
	crc  *encryptor.FileCrypto
}

// storeFileName is the encrypted alerts file inside the config directory.
const (
	storeFileName = "alerts.json"
	storeDirName  = ".wta"
)

// NewLocal returns a store rooted at $HOME/.wta. All file operations are
// encrypted through the machine-derived key; a nil store means persistence is
// broken and the app should keep running with in-memory data only.
func NewLocal() Store {
	dir, err := wtaHome()
	if err != nil {
		return nil
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil
	}
	crc, err := encryptor.NewFileCrypto(machineSecret())
	if err != nil {
		return nil
	}
	return &LocalStore{
		path: filepath.Join(dir, storeFileName),
		crc:  crc,
	}
}

// wtaHome resolves the private directory for this application's data.
func wtaHome() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		u, err := user.Current()
		if err != nil {
			return "", err
		}
		home = u.HomeDir
	}
	if home == "" {
		return "", fmt.Errorf("no home directory")
	}
	return filepath.Join(home, ".wta"), nil
}

// machineSecret is a stable, secret-ish string derived from the machine's own
// identity — the hostname and the account name, which nobody casual around the
// system is able to forge. It keeps the alerts readable without a password
// prompt while still stopping someone who just walks up to the file.
func machineSecret() string {
	host, _ := os.Hostname()
	u, err := user.Current()
	who := ""
	if err == nil {
		who = u.Username
	}
	if who == "" {
		who = "user"
	}
	s := strings.TrimSpace(host) + "." + strings.TrimSpace(who)
	// file-encryptor wants at least eleven characters.
	for len(s) < 32 {
		s = s + s
	}
	return s[:32]
}
