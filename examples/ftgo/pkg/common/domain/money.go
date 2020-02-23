package domain

type Money int

func (m Money) Equals(value interface{}) bool {
	switch v := value.(type) {
	case Money:
		return m == v
	case int:
		return m == Money(v)
	default:
		return false
	}
}
