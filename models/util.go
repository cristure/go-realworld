package models

type FilterArticle struct {
	Tag       string
	Author    string
	Favorited string
	Limit     int
	Offset    int
}

func Filter[T any](ss []T, filter func(T) bool) (ret []T) {
	for _, s := range ss {
		if filter(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func FilterArticles(articles []Article, filter FilterArticle) (result []Article) {
	for _, a := range articles {

		//Tags
		if filter.Tag != "" {
			tags := a.Tags

			for _, t := range tags {
				if t.Name == filter.Tag {
					result = append(result, a)
				}
			}
		}

		//Author
		if filter.Author != "" {
			user, err := GetUserByID(a.ID)
			if err != nil {
				return nil
			}
			if user.Username == filter.Author {
				result = append(result, a)
			}
		}

		////Favorited
		//if filter.Favorited
	}

	return result
}
