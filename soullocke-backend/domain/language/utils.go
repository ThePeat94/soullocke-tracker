package language

func GetLanguageIds(languages []Language) []uint16 {
	ids := make([]uint16, 0, len(languages))
	for _, l := range languages {
		ids = append(ids, l.ID)
	}
	return ids
}
