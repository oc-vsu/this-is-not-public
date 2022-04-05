package model

type Custom struct {
	ID string `json:"id"`
}

func (c Custom) ToString() string {
	return c.ID
}
