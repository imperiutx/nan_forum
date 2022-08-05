package common

import (
	"fmt"
	"strings"
)

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(err, "internal error", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyExists", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		"You have no permission",
		"ErrNoPermission",
	)
}

func ErrCannotGenerateToken(tokenType string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot generate %s token", tokenType),
		"ErrCannotGenerateToken",
	)
}

func ErrCannotSignInByEmail(email string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot sign %s in, check your email again", email),
		"ErrCannotSignInByEmail",
	)
}

func ErrCannotList(err error) *AppError {
	return NewCustomError(
		err,
		"You can ONLY choose ONE ID",
		"ErrCannotList",
	)
}

func ErrCannotConvertGivenData(data interface{}, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot convert %v, check your data again", data),
		"ErrCannotConvertGivenData",
	)
}
