package swagger

import (
)

type Message struct {
    ExtensionId  string  `json:"extensionId,omitempty"`
    To  []string  `json:"to,omitempty"`
    Type_  string  `json:"type,omitempty"`
    Body  string  `json:"body,omitempty"`
    
}
