package utils

import "errors"

var ErrPageExceedsBounds = errors.New("page value exceeds bounds")

func GetIndexRange(page, total, perPage int) (int, int, error) {
	if page < 1 {
		return -1, -1, ErrPageExceedsBounds
	}
	start := (page - 1) * perPage
	if start >= total {
		return -1, -1, ErrPageExceedsBounds
	}

	end := start + perPage
	if end >= total {
		return start, total, nil
	}
	return start, end, nil
}
