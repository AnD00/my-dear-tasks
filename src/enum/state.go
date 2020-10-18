package enum

import (
	"fmt"
	"io"
)

// enumに使用する型宣言
type State int

// Stateに使用する定数を宣言
const (
	_ State = iota
	OPEN
	DOING
	CLOSE
)

// https://gqlgen.com/reference/scalars/
// UnmarshalGQL implements the graphql.Unmarshaler interface
func (state *State) UnmarshalGQL(v interface{}) error {
	enum, ok := v.(string)
	if !ok {
		return fmt.Errorf("state must be strings")
	}

	switch enum {
	case "OPEN":
		*state = OPEN
	case "DOING":
		*state = DOING
	case "CLOSE":
		*state = CLOSE
	}
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (state State) MarshalGQL(w io.Writer) {
	switch state {
	case OPEN:
		w.Write([]byte(`"OPEN"`))
	case DOING:
		w.Write([]byte(`"DOING"`))
	case CLOSE:
		w.Write([]byte(`"CLOSE"`))
	}
}

// Stringer interfaceを自前で実装する場合
//func (state State) String() string  {
//	switch state {
//	case OPEN:
//		return "OPEN"
//	case DOING:
//		return "DOING"
//	case CLOSE:
//		return "CLOSE"
//	}
//	panic("Unknown State value")
//}

// Stringer interfaceをツールで実装する場合
// $ go get golang.org/x/tools/cmd/stringer
// $ stringer -type State state.go
// enum/state_string.go が生成される
