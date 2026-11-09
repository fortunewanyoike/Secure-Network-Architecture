// Package proxy provides secure HTTP/2 header parsing with strict bounds validation
package proxy

import (
	"errors"
	"fmt"
)

var ErrHeaderBlockTooLarge = errors.New("HTTP/2 header block fragment exceeds safety bounds")

// ValidateHeaderFrame checks incoming HPACK compressed header length and frame size.
func ValidateHeaderFrame(payloadSize int, maxAllowed int) error {
	if payloadSize <= 0 {
		return errors.New("invalid header frame payload size")
	}
	if payloadSize > maxAllowed {
		return fmt.Errorf("%w: size %d exceeds limit %d", ErrHeaderBlockTooLarge, payloadSize, maxAllowed)
	}
	return nil
}
