package constants

const (
	// 1xx informational response
	STATUS_CONTINUE            = "Continue"
	STATUS_SWITCHING_PROTOCOLS = "Switching Protocols"
	STATUS_PROCESSING          = "Processing"
	STATUS_EARLY_HINTS         = "Early Hints"

	// 2xx success
	STATUS_OK                            = "OK"
	STATUS_CREATED                       = "Created"
	STATUS_ACCEPTED                      = "Accepted"
	STATUS_NON_AUTHORITATIVE_INFORMATION = "Non-Authoritative Information"
	STATUS_NO_CONTENT                    = "No Content"
	STATUS_RESET_CONTENT                 = "Reset Content"
	STATUS_PARTIAL_CONTENT               = "Partial Content"

	// 3xx redirection
	STATUS_MULTIPLE_CHOICES   = "Multiple Choices"
	STATUS_MOVED_PERMANENTLY  = "Moved Permanently"
	STATUS_FOUND              = "Found"
	STATUS_SEE_OTHER          = "See Other"
	STATUS_NOT_MODIFIED       = "Not Modified"
	STATUS_TEMPORARY_REDIRECT = "Temporary Redirect"
	STATUS_PERMANENT_REDIRECT = "Permanent Redirect"

	// 4xx client errors
	STATUS_BAD_REQUEST                     = "Bad Request"
	STATUS_UNAUTHORIZED                    = "Unauthorized"
	STATUS_PAYMENT_REQUIRED                = "Payment Required"
	STATUS_FORBIDDEN                       = "Forbidden"
	STATUS_NOT_FOUND                       = "Not Found"
	STATUS_METHOD_NOT_ALLOWED              = "Method Not Allowed"
	STATUS_NOT_ACCEPTABLE                  = "Not Acceptable"
	STATUS_PROXY_AUTHENTICATION_REQUIRED   = "Proxy Authentication Required"
	STATUS_REQUEST_TIMEOUT                 = "Request Timeout"
	STATUS_CONFLICT                        = "Conflict"
	STATUS_GONE                            = "Gone"
	STATUS_LENGTH_REQUIRED                 = "Length Required"
	STATUS_PRECONDITION_FAILED             = "Precondition Failed"
	STATUS_PAYLOAD_TOO_LARGE               = "Payload Too Large"
	STATUS_URI_TOO_LONG                    = "URI Too Long"
	STATUS_UNSUPPORTED_MEDIA_TYPE          = "Unsupported Media Type"
	STATUS_RANGE_NOT_SATISFIABLE           = "Range Not Satisfiable"
	STATUS_EXPECTATION_FAILED              = "Expectation Failed"
	STATUS_IM_A_TEAPOT                     = "I'm a teapot"
	STATUS_UNPROCESSABLE_ENTITY            = "Unprocessable Entity"
	STATUS_TOO_EARLY                       = "Too Early"
	STATUS_UPGRADE_REQUIRED                = "Upgrade Required"
	STATUS_PRECONDITION_REQUIRED           = "Precondition Required"
	STATUS_TOO_MANY_REQUESTS               = "Too Many Requests"
	STATUS_REQUEST_HEADER_FIELDS_TOO_LARGE = "Request Header Fields Too Large"
	STATUS_UNAVAILABLE_FOR_LEGAL_REASONS   = "Unavailable For Legal Reasons"

	// 5xx server errors
	STATUS_INTERNAL_SERVER_ERROR           = "Internal Server Error"
	STATUS_NOT_IMPLEMENTED                 = "Not Implemented"
	STATUS_BAD_GATEWAY                     = "Bad Gateway"
	STATUS_SERVICE_UNAVAILABLE             = "Service Unavailable"
	STATUS_GATEWAY_TIMEOUT                 = "Gateway Timeout"
	STATUS_HTTP_VERSION_NOT_SUPPORTED      = "HTTP Version Not Supported"
	STATUS_VARIANT_ALSO_NEGOTIATES         = "Variant Also Negotiates"
	STATUS_INSUFFICIENT_STORAGE            = "Insufficient Storage"
	STATUS_LOOP_DETECTED                   = "Loop Detected"
	STATUS_NOT_EXTENDED                    = "Not Extended"
	STATUS_NETWORK_AUTHENTICATION_REQUIRED = "Network Authentication Required"
)
