package diary

import (
	"github.com/xrstf/stardew-diary/diary/differ"
	"github.com/xrstf/stardew-diary/sdv"
)

type Historian struct {
	differs []differ.Differ
}

func NewHistorian() *Historian {
	return &Historian{make([]differ.Differ, 0)}
}

func (d *Historian) AddWorker(worker differ.Differ) {
	d.differs = append(d.differs, worker)
}

func (d *Historian) AddAllWorkers() {
	d.differs = append(
		d.differs,
		&differ.FirstDay{},
		&differ.PetAdoption{},
		// &differ.MoneyBottomLine{},
		&differ.SkillLevel{},
		&differ.BridgeFixed{},
		&differ.LostBooks{},
		&differ.BambooPole{},
	)
}

func (d *Historian) History(previous, current, next *sdv.SaveGame) []string {
	changes := make([]string, 0)
	yield := func(change string) {
		changes = append(changes, change)
	}

	for _, differ := range d.differs {
		differ.Diff(previous, current, next, yield)
	}

	return changes
}
