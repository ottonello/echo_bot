package swagger

import (
)

type Extension struct {
    Id  string  `json:"id,omitempty"`
    Name  string  `json:"name,omitempty"`
    Picture  string  `json:"picture,omitempty"`
    Owner  string  `json:"owner,omitempty"`
    Webhooks  Webhook  `json:"webhooks,omitempty"`
    
}
