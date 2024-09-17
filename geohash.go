package geohash

import (
	"errors"
)

const base32 = "0123456789bcdefghjkmnpqrstuvwxyz"

// Encode converts latitude and longitude to a geohash string with the given precision.
// The precision is the number of characters in the resulting geohash.
// The higher the precision, the more accurate the geohash.
// The precision must be greater than 0.
// The latitude must be between -90 and 90.
// The longitude must be between -180 and 180.
func Encode(lat, lon float64, precision int) (string, error) {
	if err := validateInputs(lat, lon, precision); err != nil {
		return "", err
	}

	result := make([]byte, 0, precision)
	latRange := [2]float64{-90, 90}
	lonRange := [2]float64{-180, 180}
	symbol := uint8(0)
	shifts := int8(4)

	totalBits := precision * 5

	for i := 0; i < totalBits; i++ {
		if i%2 == 0 {
			mid := (lonRange[0] + lonRange[1]) / 2
			if lon > mid {
				symbol |= 1 << shifts
				lonRange[0] = mid
			} else {
				lonRange[1] = mid
			}
		} else {
			mid := (latRange[0] + latRange[1]) / 2
			if lat > mid {
				symbol |= 1 << shifts
				latRange[0] = mid
			} else {
				latRange[1] = mid
			}
		}

		shifts -= 1

		if shifts < 0 {
			result = append(result, base32[symbol])
			symbol = 0
			shifts = 4
		}
	}
	return string(result), nil
}

func validateInputs(lat, lon float64, precision int) error {
	if precision <= 0 {
		return errors.New("precision must be greater than 0")
	}
	if lat < -90 || lat > 90 {
		return errors.New("latitude must be between -90 and 90")
	}
	if lon < -180 || lon > 180 {
		return errors.New("longitude must be between -180 and 180")
	}
	return nil
}
