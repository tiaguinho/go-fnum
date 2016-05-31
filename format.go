package fnum

import (
	"bytes"
	"strconv"
	"strings"
)

// FormatInt produces a string form of the given number in base 10 with
// sep after every three orders of magnitude.
func FormatInt(v int64, sep string) string {
	sign := ""
	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}

	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], sep)
}

// FormatFloat produces a string form of the given number in base 10 with
// thousant after every three orders of magnitude and decimal for the decimal peace
func FormatFloat(v float64, thousant, decimal string) string {
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	t := []byte(thousant)

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(t)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(t)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte(decimal))
		buf.WriteString(parts[1])
	}
	return buf.String()
}
