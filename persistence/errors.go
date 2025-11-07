package persistence

import "fmt"

type SignatureDeviceNotFoundError struct{ DeviceID string }

func (e *SignatureDeviceNotFoundError) Error() string {
    return fmt.Sprintf("signature device %s not found", e.DeviceID)
}


