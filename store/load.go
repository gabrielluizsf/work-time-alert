package store

import (
	"encoding/json"
	"os"
)

// Load reads, decrypts and parses the alert file. A missing file is not an
// error: the first run simply starts empty. A corrupted ciphertext or a file
// that was tampered with surfaces as an error so the user is not left staring
// at a silently reset list.
func (s *LocalStore) Load() ([]Alert, error) {
	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		return nil, nil
	}

	temp, err := os.CreateTemp("", "wta-alerts-*.json")
	if err != nil {
		return nil, err
	}
	plain := temp.Name()
	temp.Close()
	defer os.Remove(plain)

	if err := s.crc.Decrypt(s.path, plain); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(plain)
	if err != nil {
		return nil, err
	}

	alerts := []Alert{}
	if err := json.Unmarshal(data, &alerts); err != nil {
		return nil, err
	}
	return alerts, nil
}
