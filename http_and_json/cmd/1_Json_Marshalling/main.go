package main

import (
	"encoding/json"
	"fmt"
)

type Color string

const (
	Green  Color = "green"
	Yellow Color = "yellow"
	Orange Color = "orange"
)

var (
	validColor = map[Color]struct{}{
		Green:  {},
		Yellow: {},
		Orange: {},
	}
)

func (c Color) String() string {
	return string(c)
}

func (c Color) MashalJSON() ([]byte, error) {
	_, ok := validColor[c]
	if !ok {
		return nil, fmt.Errorf("Invalid Color : %s", c)
	}
	return []byte(`"` + c + `"`), nil
}

func (c *Color) UnmarshallingJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("failed to unmarshal : %w", err)
	}
	_, ok := validColor[Color(s)]
	if !ok {
		return fmt.Errorf("Invalid Color : %s", s)
	}
	*c = Color(s)

	return nil
}

func main() {
	c := Green
	b, err := json.Marshal(c)

	if err != nil {
		fmt.Println("Error marshalling color : ", err)
		panic(err)
	}

	fmt.Println(string(b))

	var unmarshalledColor Color
	err = json.Unmarshal([]byte(`"orange"`), &unmarshalledColor)
	if err != nil {
		fmt.Println("Error Unmarshalling color : ", err)
		panic(err)
	}

	fmt.Println(unmarshalledColor)
}
