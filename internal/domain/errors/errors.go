package globalerrors

import "errors"

var (
	UserEmailExists     = errors.New("email already exists")
	UserUsernameExists  = errors.New("username already exists")
	UserInvalidUsername = errors.New("invalid username")
	UserInvalidEmail    = errors.New("invalid email")
	UserNotFound        = errors.New("user not found")

	AuthInvalidToken            = errors.New("invalid token")
	AuthInvalidCredentials      = errors.New("invalid credentials")
	AuthMissingAuthHeader       = errors.New("missing authorization header")
	AuthInvalidAuthHeaderFormat = errors.New("Invalid authorization header Format")

	RecipeTitleOfUserExists       = errors.New("recipe title already exists for user")
	RecipeNotFound                = errors.New("recipe not found")
	RecipeDuplicateIngredient     = errors.New("recipe contains duplicate ingredients")
	RecipeDuplicateTag            = errors.New("recipe contains duplicate tags")
	RecipeInvalidIngredient       = errors.New("recipe contains invalid ingredient(s)")
	RecipeInvalidUnit             = errors.New("recipe contains invalid unit(s)")
	RecipeInvalidQuantity         = errors.New("recipe contains invalid quantity(ies)")
	RecipeInvalidTag              = errors.New("recipe contains invalid tag(s)")
	RecipeInvalidTitle            = errors.New("recipe contains invalid title")
	RecipeIngredientsMustBeUnique = errors.New("recipe ingredients must be unique")
	RecipeIngredientNotFound      = errors.New("recipe ingredient not found")

	GlobalInternalServerError = errors.New("internal server error")
	GlobalUnableToParseBody   = errors.New("unable to parse body")
	GlobalInvalidID           = errors.New("invalid id. must be positive integer")
)
