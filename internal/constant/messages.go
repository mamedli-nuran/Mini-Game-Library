package constant

const (
	ErrEnvLoading          = "Error loading .env file, check root directory"
	ErrEmailRequired       = "Email is required, please enter your email"
	ErrInvalidEmail        = "Invalid email, make sure that you enter valid email"
	ErrValidationFailed    = "Validation failed, please try again with valid data"
	ErrInvalidBody         = "Invalid body"
	ErrUserAlreadyExists   = "User already exists or could not be created"
	ErrInternalServerError = "Internal Server Error"

	ErrUsernameRequired = "Username is required, please enter your username"
	ErrUsernameLength   = "Your username field should be less than 20 and greater than 4"
	ErrPasswordLength   = "Password must be between 12 and 25 characters."

	ErrUnauthorized       = "User unauthorized"
	ErrInvalidGenre       = "Invalid genre, please use real genre"
	ErrInvalidReleaseYear = "Invalid release year, write only digits and valid year"
	ErrMissingId          = "Missing id parameter"
	ErrParseId            = "Invalid id parameter, can not parse id"
)
