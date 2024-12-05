package constants

import (
	"time"
)

const (
	DDMMYY = "02-01-2006"
	DMY    = "02/01/2006"

	// Configurations Keys
	ETCD             = "etcd"
	Services         = "services"
	Storages         = "storages"
	Brokers          = "brokers"
	ServiceDiscovery = "serviceDiscovery"
	Encryption       = "encryption"
	Databases        = "databases"
	Monitoring       = "monitoring"
	OAUTH            = "oauth"
	TLS              = "tls"

	// Configurations
	AppTLS     = "app"
	ConsulTLS  = "consul"
	KafkaTLS   = "kafka"
	MariadbTLS = "mariadb"
	PasetoTLS  = "paseto"
	RedisTLS   = "redis"

	Validate        = "validate"
	FieldValidation = "field validation"
	RequiredHeaders = "required header"
	Base64          = "base64"
	Unmarshal       = "unmarshal"
	Uuid            = "uuid"
	Cookie          = "cookie"
	Token           = "token"
	Bcrypt          = "bcrypt"
	SQLState        = "sqlstate"
	EnvType         = "env"
	Yaml            = "yaml"
	EnvName         = "app"

	// uuidErr
	InvalidUUIDLength = "uuid length"
	InvalidUUIDFormat = "uuid format"

	// http-headers
	Accept              = "Accept"
	ContentType         = "Content-Type"
	ApplicationJSON     = "application/json; charset=utf8"
	XForwardedFor       = "X-Forwarded-For"
	UserAgent           = "user-agent"
	XPlatform           = "x-pt-e3f264be"
	XAPIKey             = "x-api-key"
	XIdempotencyKey     = "x-idempotency-key"
	WithCredentials     = "withCredentials"
	WithCredentialsTrue = "true"
	Authorization       = "Authorization"
	Bearer              = "Bearer"

	// string concatenation
	TrailingSlash = "/"
	Ampersand     = "&"
	Equal         = "="
	QuestionMark  = "?"
	Delete        = "delete"
	Publish       = "publish"

	// http-form
	FileForm = "file"

	// http-cookies
	AccessToken    = "_s_04250fa5_session_id"
	RefreshToken   = "_s_65r60230_session_id"
	SameSite       = "strict"
	Path           = "/"
	APICallTimeout = 15 * time.Second
	Production     = "production"

	// http-params / query-params
	UserID     = "userId"
	SecretCode = "secretCode"
	SerialCode = "serialCode"
	ID         = "id"
	UID        = "uid"
	NIK        = "nik"
	Search     = "search"
	Q          = "q"
	PageSize   = "page_size"
	PageID     = "page_id"

	// mime types
	ImageWEBP = "image/webp"
	ImageJPEG = "image/jpeg"
	ImageJPG  = "image/jpg"
	ImagePNG  = "image/png"

	// Databases
	ElasticSearch = "ElasticSearch"
	MongoDB       = "MongoDB"
	Redis         = "redis"
	PostgreSQL    = "PostgreSQL"
	MariaDB       = "mariadb"
	MySQL         = "mysql"

	// Logger
	GRPC       = "gRPC"
	Size       = "size"
	URI        = "URI"
	Status     = "status"
	StatusCode = "status_code"
	StatusText = "status_text"
	HTTP       = "HTTP"
	GraphQL    = "GraphQL"
	Error      = "ERROR"
	Protocol   = "protocol"
	Duration   = "duration"
	Method     = "method"
	MetaData   = "metadata"
	Request    = "request"
	Reply      = "reply"
	Time       = "time"
	Took       = "took"

	// TokenType
	AccessType  = "access"
	RefreshType = "refresh"

	// CleanArchitecture Layers
	Resolver       = "resolver"
	Handler        = "handler"
	Usecase        = "usecase"
	Repository     = "repository"
	ProducerWorker = "producerWorker"
	Config         = "config"
	Worker         = "worker"

	PaymentService = "payment-service"
)
