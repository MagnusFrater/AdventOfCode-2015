package circuit

import "strconv"

func getUint16(s string) (uint16, error) {
	n, err := strconv.ParseUint(s, 0, 0)
	if err != nil {
		return 0, err
	}
	return uint16(n), nil
}
