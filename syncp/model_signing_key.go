package syncp

import (
	"encoding/json"
	"fmt"
	"sort"
)

var NoLimit int64 = -1

type UserScope struct {
	Kind     SigningKeyScopeType  `json:"kind"`
	Key      string               `json:"key"`
	Role     string               `json:"role"`
	Template UserPermissionLimits `json:"template"`
}

func NewUserScope() *UserScope {
	var s UserScope
	s.Kind = SIGNINGKEYSCOPETYPE_USER_SCOPE
	s.Template.NatsLimits = NatsLimits{&NoLimit, &NoLimit, &NoLimit}
	return &s
}

func (us UserScope) SigningKey() string {
	return us.Key
}

// SigningKeys is a map keyed by a public account key
type SigningKeys map[string]*UserScope

// MarshalJSON serializes the scoped signing keys as an array
func (sk *SigningKeys) MarshalJSON() ([]byte, error) {
	if sk == nil {
		return nil, nil
	}

	keys := sk.Keys()
	sort.Strings(keys)

	var a []interface{}
	for _, k := range keys {
		if (*sk)[k] != nil {
			a = append(a, (*sk)[k])
		} else {
			a = append(a, k)
		}
	}
	return json.Marshal(a)
}

func (sk *SigningKeys) UnmarshalJSON(data []byte) error {
	if *sk == nil {
		*sk = make(SigningKeys)
	}
	// read an array - we can have a string or a map
	var a []interface{}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	for _, i := range a {
		switch v := i.(type) {
		case string:
			(*sk)[v] = nil
		case map[string]interface{}:
			d, err := json.Marshal(v)
			if err != nil {
				return err
			}
			switch v["kind"] {
			case string(SIGNINGKEYSCOPETYPE_USER_SCOPE):
				us := NewUserScope()
				if err := json.Unmarshal(d, &us); err != nil {
					return err
				}
				(*sk)[us.Key] = us
			default:
				return fmt.Errorf("unknown signing key scope %q", v["type"])
			}
		}
	}
	return nil
}

func (sk SigningKeys) Keys() []string {
	var keys []string
	for k := range sk {
		keys = append(keys, k)
	}
	return keys
}

// GetScope returns nil if the key is not associated
func (sk SigningKeys) GetScope(k string) (*UserScope, bool) {
	v, ok := sk[k]
	if !ok {
		return nil, false
	}
	return v, true
}

func (sk SigningKeys) Contains(k string) bool {
	_, ok := sk[k]
	return ok
}

func (sk SigningKeys) Add(keys ...string) {
	for _, k := range keys {
		sk[k] = nil
	}
}

func (sk SigningKeys) AddScopedSigner(s *UserScope) {
	sk[s.SigningKey()] = s
}

func (sk SigningKeys) Remove(keys ...string) {
	for _, k := range keys {
		delete(sk, k)
	}
}
