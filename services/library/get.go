package library

import (
	"fmt"

	dto "posttest-be/dto"

	"github.com/labstack/echo/v4"
)

func (s *libraryService) Get(context echo.Context, req *dto.Pagination) dto.Response {
	if req.Page < 1 {
		req.Page = 1
	}

	operationResult, totalPages := s.Repo.Get(req)

	// Set page 1-based untuk pagination link
	urlPath := context.Request().URL.Path
	searchQueryParams := ""

	for _, search := range req.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// URL dengan base 1 untuk first page, previous page, next page, dan last page
	data := operationResult.Result.(*dto.Pagination)
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=1&sort=%s", urlPath, req.Limit, req.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, req.Limit, totalPages, req.Sort) + searchQueryParams

	if req.Page > 1 {
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, req.Limit, req.Page-1, req.Sort) + searchQueryParams
	}
	if req.Page < totalPages {
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, req.Limit, req.Page+1, req.Sort) + searchQueryParams
	}

	return dto.Response{Success: true, Data: data}
}
