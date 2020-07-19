package schema

type Option struct {
	T        string    `json:"t"`
	K        string    `json:"k"`
	V        string    `json:"v"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

type Options []*Option

func (s Options) Filter(t string) Options {
	ret := Options{}
	for _, x := range s {
		if x.T == t {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s Options) Get(t string, k string) *Option {
	for _, x := range s {
		if x.T == t && x.K == k {
			return x
		}
	}
	return nil
}
