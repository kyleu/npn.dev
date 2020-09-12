package request

type Request struct {
	Key         string     `json:"-"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Prototype   *Prototype `json:"prototype"`
}

func NewRequest() *Request {
	return &Request{Prototype: NewPrototype()}
}

func (r *Request) TitleWithFallback() string {
	if len(r.Title) == 0 {
		return r.Key
	}
	return r.Title
}

type Requests []*Request

func (r *Request) Normalize(key string) *Request {
	if r == nil {
		return nil
	}
	r.Key = key
	if r.Prototype == nil {
		r.Prototype = NewPrototype()
	}
	r.Prototype = r.Prototype.Normalize()
	return r
}
