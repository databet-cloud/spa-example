package common

type Decimal string

func (d Decimal) String() string {
	return string(d)
}
