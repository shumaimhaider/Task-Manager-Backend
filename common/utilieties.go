package common

import (
	"errors"
	"fmt"
)

func HandleDBUtilError(err error, rowsAffected int) error {

	switch {
	case err != nil:
		return fmt.Errorf("error getting rows affetced: %w", err)
	case rowsAffected == 0:
		return errors.New("not found")
	}

	return nil
}
