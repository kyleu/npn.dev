package types

type Type struct {
	Key string
}

type Types []Type

var all = Types{TypeString}

func TypeFromKey(key string) Type {
	for _, t := range all {
		if t.Key == key {
			return t
		}
	}
	return TypeString
}
