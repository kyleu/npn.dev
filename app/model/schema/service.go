package schema

import "emperror.dev/errors"

type Service struct {
	Key     string  `json:"key"`
	Methods Methods `json:"methods,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

func (s *Service) AddMethod(m *Method) error {
	if s.Methods.Get(m.Key) != nil {
		return errors.New("method [" + m.Key + "] already exists")
	}
	s.Methods = append(s.Methods, m)
	return nil
}

type Services []*Service

func (s Services) Get(key string) *Service {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}
