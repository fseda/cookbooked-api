package globalerrors

import "errors"

var (
	UserEmailExists     = errors.New("email already exists")
	UserUsernameExists  = errors.New("username already exists")
	UserInvalidUsername = errors.New("invalid username")
	UserCouldNotQuery   = errors.New("could not query")

	AuthInvalidToken            = errors.New("invalid token")
	AuthInvalidCredentials      = errors.New("invalid credentials")
	AuthMissingAuthHeader       = errors.New("missing authorization header")
	AuthInvalidAuthHeaderFormat = errors.New("Invalid authorization header Format")

	RecipeTitleOfUserExists   = errors.New("recipe title already exists for user")
	RecipeDuplicateIngredient = errors.New("recipe contains duplicate ingredients")
	RecipeDuplicateTag        = errors.New("recipe contains duplicate tags")
	RecipeInvalidIngredient   = errors.New("recipe contains invalid ingredient(s)")
	RecipeInvalidUnit         = errors.New("recipe contains invalid unit")
	RecipeInvalidQuantity     = errors.New("recipe contains invalid quantity")
	RecipeInvalidTag          = errors.New("recipe contains invalid tag")
	RecipeInvalidTitle        = errors.New("recipe contains invalid title")

	GlobalInternalServerError = errors.New("internal server error")
	GlobalUnableToParseBody   = errors.New("unable to parse body")
)
