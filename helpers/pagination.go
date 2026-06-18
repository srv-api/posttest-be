package helpers

import (
	"strconv"
	"strings"

	"posttest-be/dto"

	"github.com/labstack/echo/v4"
)

func GeneratePaginationRequest(context echo.Context) *dto.Pagination {
	// default limit, page & sort parameter
	limit := 10
	page := 1
	sort := "created_at desc"

	var searchs []dto.Search

	query := context.QueryParams()

	for key, values := range query {
		queryValue := values[len(values)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}

		// check if query parameter key contains dot
		if strings.Contains(key, ".") {
			// split query parameter key by dot
			searchKeys := strings.Split(key, ".")

			// create search object
			search := dto.Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

			// add search object to searchs array
			searchs = append(searchs, search)
		}
	}

	return &dto.Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}
}

func TruncateString(input string, length int) string {
	if len(input) > length {
		return input[:length] + "..."
	}
	return input
}
