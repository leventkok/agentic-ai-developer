package domain

import "strconv"

func ValidateCreateInput(in CreateBookmarkInput) (CreateBookmarkInput, error) {
	title, err := NewTitle(in.Title)
	if err != nil {
		return CreateBookmarkInput{}, err
	}
	u, err := NewBookmarkURL(in.URL)
	if err != nil {
		return CreateBookmarkInput{}, err
	}
	return CreateBookmarkInput{
		Title: title.String(),
		URL:   u.String(),
		Tags:  in.Tags,
	}, nil
}

func ValidateUpdateInput(in UpdateBookmarkInput) (UpdateBookmarkInput, error) {
	out := UpdateBookmarkInput{Tags: in.Tags}
	if in.Title != nil {
		title, err := NewTitle(*in.Title)
		if err != nil {
			return UpdateBookmarkInput{}, err
		}
		s := title.String()
		out.Title = &s
	}
	if in.URL != nil {
		u, err := NewBookmarkURL(*in.URL)
		if err != nil {
			return UpdateBookmarkInput{}, err
		}
		s := u.String()
		out.URL = &s
	}
	return out, nil
}

func ValidateBulkCreateInputs(inputs []CreateBookmarkInput) ([]CreateBookmarkInput, error) {
	if len(inputs) == 0 {
		return nil, ValidationError{Message: "bookmarks must not be empty"}
	}
	validated := make([]CreateBookmarkInput, len(inputs))
	for i, item := range inputs {
		v, err := ValidateCreateInput(item)
		if err != nil {
			return nil, ValidationError{Message: "bookmarks[" + strconv.Itoa(i) + "]: " + err.Error()}
		}
		validated[i] = v
	}
	return validated, nil
}
