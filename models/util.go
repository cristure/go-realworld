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

func FilterArticles(articles []*Article, filter FilterArticle) (result []*Article) {
	i := 0
	n := len(articles)

	for i < n {

		//Tags
		if filter.Tag != "" {
			tags := articles[i].Tags

			ok := false
			for _, t := range tags {
				if filter.Tag == t.Name {
					ok = true
					break
				}
			}

			if !ok {
				articles = append(articles[:i], articles[i+1:]...)
				n--
				continue
			}
		}

		if filter.Author != "" {
			user, err := GetUserByID(articles[i].ID)
			if err != nil {
				return nil
			}
			if user.Username != filter.Author {
				articles = append(articles[:i], articles[i+1:]...)
				n--
				continue
			}
		}

		//Favorited
		if filter.Favorited != "" {
			user, err := GetUserByName(filter.Favorited)
			if err != nil {
				return nil
			}

			ok := false
			for _, aa := range user.FavoriteArticles {
				if articles[i].ID == aa.ID {
					ok = true
					break
				}
			}
			if !ok {
				articles = append(articles[:i], articles[i+1:]...)
				n--
				continue
			}
		}

		i++
	}

	if filter.Limit != 0 {
		if len(articles) > filter.Limit {
			articles = articles[0:filter.Limit]
		}
	}

	if filter.Offset != 0 {
		if len(articles) < filter.Offset {
			return
		}
		articles = articles[filter.Offset:]
	}

	return articles
}
