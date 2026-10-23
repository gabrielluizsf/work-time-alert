package store

import (
	"encoding/json"
	"os"
)

// Save writes alerts to a temporary file, encrypts it in place of the real
// file, and only then swaps the encrypted file into place. The temp file
// carries the final name so the swap is a plain rename — atomic on the same
// filesystem, so a crash mid-write never leaves a half-written store.
func (s *LocalStore) Save(alerts ...Alert) error {
	data, err := json.MarshalIndent(alerts, "", "\t")
	if err != nil {
		return err
	}

	temp, err := os.CreateTemp("", "wta-alerts-*.json")
	if err != nil {
		return err
	}
	plain := temp.Name()
	if _, err := temp.Write(data); err != nil {
		temp.Close()
		os.Remove(plain)
		return err
	}
	if err := temp.Close(); err != nil {
		os.Remove(plain)
		return err
	}
	defer os.Remove(plain)

	if err := s.crc.Encrypt(plain, s.path); err != nil {
		return err
	}
	_ = os.Chmod(s.path, 0o600)
	return nil
}
