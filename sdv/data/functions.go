package data

func ItemByID(id int) *Item {
	for _, item := range Items {
		if item.ID == id {
			return &item
		}
	}

	return nil
}

func LocationByID(id string) *Location {
	for _, loc := range Locations {
		if loc.ID == id {
			return &loc
		}
	}

	return nil
}
