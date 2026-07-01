package operation

import (
	"fmt"
)

func ApplyInsert(document string, position int, text string) (string, error) {
	if position < 0 || position > len(document) {
		return "", fmt.Errorf("Invalid position: %d", position)
	}

	result := document[:position] + text + document[position:]

	return result, nil
}

func ApplyDelete(document string, position int, length int) (string, error) {
	if position < 0 || position >= len(document) {
		return "", fmt.Errorf("Invalid position: %d", position)
	}

	if length < 0 || position+length > len(document) {
		return "", fmt.Errorf("Invalid length: %d", length)
	}

	result := document[:position] + document[position+length:]

	return result, nil
}

/*
Old Document
      ↓
GenerateOperation
      ↓
Operation
      ↓
ApplyOperation
      ↓
New Document
*/
func ApplyOperation(document string, op *Operation) (string, error) {
	switch op.Type {
	case Insert:
		return ApplyInsert(document,op.Position,op.Text)
	case Delete:
		return ApplyDelete(document,op.Position,op.Length)
	default:
		return "", fmt.Errorf("unknown operation type")
	}
}