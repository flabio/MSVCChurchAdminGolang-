package utils

const (
	DB_DIFF_ID              = "id <>?"
	DB_EQUAL_ID             = "id=?"
	DB_NAME                 = "name =?"
	DB_USERNAME_EQUAL       = "username=?"
	DB_EMAIL_EQUAL          = "email=?"
	DB_IDENTIFICATION_EQUAL = "identification=?"
	DB_ROL_ID_EQUAL         = "rol_id=?"
)
const (
	ID      = "id"
	MESSAGE = "message"
	STATUS  = "status"
	DATA    = "data"
	GET     = "GET"
	SUCCESS = "success"
	FAIL    = "fail"
	CREATED = "was successfully created"
	UPDATED = "was updated successfully"
	REMOVED = "was successfully removed"
)

const (
	ID_NO_EXIST                  = "The id not exists"
	NAME_ALREADY_EXIST           = "name already exists"
	IDENTIFICATION_ALREADY_EXIST = "identification already exists"
	EMAIL_ALREADY_EXIST          = "email already exists"
	USERNAME_ALREADY_EXIST       = "username already exists"
)

const (
	ROL_MSVC_URL                       = "http://localhost:3001/api/rol/"
	CHURCH_MSVC_URL                    = "http://localhost:3002/api/church/"
	TEAM_MSVC_URL                      = "http://localhost:3003/api/team_pesca/"
	MINISTERIAL_USER_MSVC_URL          = "http://localhost:3004/api/ministerial/user_ministerial/"
	FUNCTION_MINISTERIAL_USER_MSVC_URL = "http://localhost:3005/api/function_ministerial/user"
)
const (
	ROL_ID    = "The id rol not exist"
	CHURCH_ID = "The id church not exist"
)

const (
	INTERNAL_SERVER_ERROR          = "Internal Server Error"
	UNAUTHORIZED                   = "Unauthorized"
	FORBIDDEN                      = "Forbidden"
	NOT_FOUND                      = "Not Found"
	METHOD_NOT_ALLOWED             = "Method Not Allowed"
	REQUEST_ENTITY_TOO_LARGE       = "Request Entity Too Large"
	TOO_MANY_REQUESTS              = "Too Many Requests"
	PRECONDITION_FAILED            = "Precondition Failed"
	REQUEST_TIMEOUT                = "Request Timeout"
	CONFLICT                       = "Conflict"
	UNSUPPORTED_MEDIA_TYPE         = "Unsupported Media Type"
	TOO_MANY_REQUESTS_ERROR        = "Too Many Requests. Please try again later."
	INTERNAL_SERVER_ERROR_ERROR    = "Internal Server Error. Please try again later."
	UNAUTHORIZED_ERROR             = "Unauthorized. Please authenticate."
	FORBIDDEN_ERROR                = "Forbidden. You are not authorized to access this resource."
	NOT_FOUND_ERROR                = "Not Found. The requested resource was not found."
	METHOD_NOT_ALLOWED_ERROR       = "Method Not Allowed. The requested method is not allowed for this resource."
	REQUEST_ENTITY_TOO_LARGE_ERROR = "Request Entity Too Large. The request entity is too large."
	PRECONDITION_FAILED_ERROR      = "Precondition Failed. The precondition given in the request headers is not met."
	REQUEST_TIMEOUT_ERROR          = "Request Timeout. The server did not receive a request within the time that it was prepared to."
	CONFLICT_ERROR                 = "Conflict. The request could not be completed due to a conflict with the current state of the resource."
	UNSUPPORTED_MEDIA_TYPE_ERROR   = "Unsupported Media Type. The server does not support the requested media type."
	EMAIL_NOT_FOUND                = "The email not found."
	EMAIL_ALREADY_EXISTS           = "The email already exists."
	PASSWORD_INCORRECT             = "The password is incorrect."
)
