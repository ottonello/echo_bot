package swagger

import (
)

type ModelError struct {
    Code  int32  `json:"code,omitempty"`
    Status  string  `json:"status,omitempty"`
    Message  string  `json:"message,omitempty"`
    
}
