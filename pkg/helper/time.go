package helper

import "time"

func ParseTime(t string) (*time.Time, error) {
	v, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
