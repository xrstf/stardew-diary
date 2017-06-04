package differ

import (
	"fmt"
	"strings"

	"github.com/xrstf/stardew-diary/sdv"
)

type PetAdoption struct{}

func (d *PetAdoption) Diff(previous, current, next *sdv.SaveGame, out Yielder) {
	if pet := current.Pet(); pet != nil {
		out(fmt.Sprintf("I've adopted a %s and named it %s!", strings.ToLower(pet.Type), pet.Name))
	}
}
