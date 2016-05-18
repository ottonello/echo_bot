package swagger

import (
)

type NewExtension struct {
    Name  string  `json:"name,omitempty"`
    Picture  string  `json:"picture,omitempty"`
    Owner  string  `json:"owner,omitempty"`
    Webhooks  Webhook  `json:"webhooks,omitempty"`
    
}
