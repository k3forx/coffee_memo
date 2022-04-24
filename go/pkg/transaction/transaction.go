package transaction

import (
	"fmt"

	"github.com/k3forx/coffee_memo/pkg/ent"
)

func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
