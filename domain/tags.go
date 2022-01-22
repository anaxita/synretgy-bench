package domain

import "synergybench/entity"

func CheckTagsMap(allowedTags []entity.Tag, tags []int64) []int64 {
	if len(tags) == 0 {
		return tags
	}

	newTags := make([]int64, 0, len(tags))

	f := make(map[int64]struct{}, len(allowedTags))

	for _, tag := range allowedTags {
		f[tag.ID] = struct{}{}
	}

	for _, check := range tags {
		_, ok := f[check]
		if !ok {
			continue
		}

		newTags = append(newTags, check)
	}

	return newTags
}

func CheckTagsLoop(allowedTags []entity.Tag, tags []int64) []int64 {
	if len(tags) == 0 {
		return tags
	}

	newTags := make([]int64, 0, len(tags))

	for _, check := range tags {
		for _, v := range allowedTags {
			if v.ID == check {
				newTags = append(newTags, check)

				break
			}
		}
	}

	return newTags
}
