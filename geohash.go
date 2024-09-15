package geohash

const base32 = "0123456789bcdefghjkmnpqrstuvwxyz"

// TODO: fixme
func Encode(lat, lon float32, precision int) string {
	result := make([]byte, 0, precision)
	latRange := []float32{-90, 90}
	lonRange := []float32{-180, 180}
	symbol := 0
	bitIndex := 0

	for len(result) < precision {
		if bitIndex%2 == 0 {
			mid := (lonRange[0] + lonRange[1]) / 2
			if lon >= mid {
				symbol |= 1 << (4 - bitIndex%5)
				lonRange[0] = mid
			} else {
				lonRange[1] = mid
			}
		} else {
			mid := (latRange[0] + latRange[1]) / 2
			if lat >= mid {
				symbol |= 1 << (4 - bitIndex%5)
				latRange[0] = mid
			} else {
				latRange[1] = mid
			}
		}

		bitIndex++

		if bitIndex%5 == 0 {
			result = append(result, base32[symbol])
			symbol = 0
		}
	}
	return string(result)
}
