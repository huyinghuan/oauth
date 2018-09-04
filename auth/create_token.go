package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

var TokenSeparator = "."

type ResourceActions struct {
	Type    string   `json:"type"`
	Class   string   `json:"class,omitempty"`
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}
type ClaimSet struct {
	// Public claims
	Issuer     string `json:"iss"`
	Subject    string `json:"sub"`
	Audience   string `json:"aud"`
	Expiration int64  `json:"exp"`
	NotBefore  int64  `json:"nbf"`
	IssuedAt   int64  `json:"iat"`
	JWTID      string `json:"jti"`

	// Private claims
	Access []*ResourceActions `json:"access"`
}

// Copy-pasted from libtrust where it is private.
func joseBase64UrlEncode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

func CreateToken(form *UserPostForm) (string, error) {
	//conf := config.Get()
	now := time.Now().Unix()
	// header := token.Header{
	// 	Type:       "JWT",
	// 	SigningAlg: sigAlg,
	// 	KeyID:      tokenConfig.PublicKey.KeyID(),
	// }
	// headerJSON, err := json.Marshal(header)
	// if err != nil {
	// 	return "", err
	// }
	claims := ClaimSet{
		Issuer:     "Test", // TODO
		Subject:    form.Account,
		Audience:   form.Service,
		NotBefore:  now - 10,
		IssuedAt:   now,
		Expiration: now + 900, // TODO
		JWTID:      fmt.Sprintf("%d", rand.Int63()),
		Access:     []*ResourceActions{},
	}

	if len(form.Scopes) > 0 {
		for _, scope := range form.Scopes {
			ra := &ResourceActions{
				Type:    scope.Type,
				Name:    scope.Name,
				Actions: scope.Actions,
			}
			if ra.Actions == nil {
				ra.Actions = []string{}
			}
			sort.Strings(ra.Actions)
			claims.Access = append(claims.Access, ra)
		}
	}
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payload := fmt.Sprintf("%s%s%s", joseBase64UrlEncode([]byte{}), TokenSeparator, joseBase64UrlEncode(claimsJSON))
	// sig, sigAlg2, err := tokenConfig.PrivateKey.Sign(strings.NewReader(payload), 0)
	// if err != nil || sigAlg2 != sigAlg {
	// 	return "", fmt.Errorf("failed to sign token: %s", err)
	// }
	// return fmt.Sprintf("%s%s%s", payload, TokenSeparator, joseBase64UrlEncode(sig)), nil
	return payload, nil
}
