package api

import (
	"path"
	"regexp"

	"github.com/xigxog/kubefox/utils"
)

const (
	DefaultLogFormat             = "json"
	DefaultLogLevel              = "info"
	DefaultMaxEventSizeBytes     = 5242880 // 5 MiB
	DefaultReleaseLimitCount     = 100
	DefaultReleaseTimeoutSeconds = 300 // 5 mins
	DefaultTimeoutSeconds        = 30

	MaximumMaxEventSizeBytes = 16777216 // 16 MiB
)

// Kubernetes Labels
const (
	LabelK8sAppBranch           string = "kubefox.xigxog.io/app-branch"
	LabelK8sAppCommit           string = "kubefox.xigxog.io/app-commit"
	LabelK8sAppDeployment       string = "kubefox.xigxog.io/app-deployment"
	LabelK8sAppName             string = "app.kubernetes.io/name"
	LabelK8sAppTag              string = "kubefox.xigxog.io/app-tag"
	LabelK8sAppVersion          string = "kubefox.xigxog.io/app-version"
	LabelK8sComponent           string = "app.kubernetes.io/component"
	LabelK8sComponentCommit     string = "kubefox.xigxog.io/component-commit"
	LabelK8sEnvironment         string = "kubefox.xigxog.io/environment"
	LabelK8sInstance            string = "app.kubernetes.io/instance"
	LabelK8sPlatform            string = "kubefox.xigxog.io/platform"
	LabelK8sReleaseStatus       string = "kubefox.xigxog.io/release-status"
	LabelK8sResolvedEnvironment string = "kubefox.xigxog.io/resolved-environment"
	LabelK8sRuntimeVersion      string = "app.kubernetes.io/runtime-version"
)

// Kubernetes Annotations
const (
	AnnotationTemplateData     string = "kubefox.xigxog.io/template-data"
	AnnotationTemplateDataHash string = "kubefox.xigxog.io/template-data-hash"
	AnnotationUpdateTime       string = "kubefox.xigxog.io/update-time"
)

// Container Labels
const (
	LabelOCIApp       string = "com.xigxog.kubefox.app"
	LabelOCIComponent string = "com.xigxog.kubefox.component"
	LabelOCICreated   string = "org.opencontainers.image.created"
	LabelOCIRevision  string = "org.opencontainers.image.revision"
	LabelOCISource    string = "org.opencontainers.image.source"
)

const (
	FinalizerReleaseProtection string = "kubefox.xigxog.io/release-protection"
)

const (
	EnvNodeName = "KUBEFOX_NODE"
	EnvPodIP    = "KUBEFOX_POD_IP"
	EnvPodName  = "KUBEFOX_POD"
)

type EnvVarType string

const (
	EnvVarTypeArray   EnvVarType = "array"
	EnvVarTypeBoolean EnvVarType = "boolean"
	EnvVarTypeNumber  EnvVarType = "number"
	EnvVarTypeString  EnvVarType = "string"
)

type ComponentType string

const (
	ComponentTypeDatabase ComponentType = "db"
	ComponentTypeGenesis  ComponentType = "genesis"
	ComponentTypeHTTP     ComponentType = "http"
	ComponentTypeKubeFox  ComponentType = "kubefox"
)

type FollowRedirects string

const (
	FollowRedirectsAlways   FollowRedirects = "Always"
	FollowRedirectsNever    FollowRedirects = "Never"
	FollowRedirectsSameHost FollowRedirects = "SameHost"
)

type EventType string

// Component event types
const (
	EventTypeCron       EventType = "io.kubefox.cron"
	EventTypeDapr       EventType = "io.kubefox.dapr"
	EventTypeHTTP       EventType = "io.kubefox.http"
	EventTypeKubeFox    EventType = "io.kubefox.kubefox"
	EventTypeKubernetes EventType = "io.kubefox.kubernetes"
)

// Platform event types
const (
	EventTypeAck       EventType = "io.kubefox.ack"
	EventTypeBootstrap EventType = "io.kubefox.bootstrap"
	EventTypeError     EventType = "io.kubefox.error"
	EventTypeHealth    EventType = "io.kubefox.health"
	EventTypeMetrics   EventType = "io.kubefox.metrics"
	EventTypeNack      EventType = "io.kubefox.nack"
	EventTypeRegister  EventType = "io.kubefox.register"
	EventTypeRejected  EventType = "io.kubefox.rejected"
	EventTypeTelemetry EventType = "io.kubefox.telemetry"
	EventTypeUnknown   EventType = "io.kubefox.unknown"
)

type ReleaseType string

const (
	ReleaseTypePromotion ReleaseType = "Promotion"
	ReleaseTypeRelease   ReleaseType = "Release"
	ReleaseTypeRollback  ReleaseType = "Rollback"
)

type ReleaseStatus string

const (
	ReleaseStatusFailed     ReleaseStatus = "Failed"
	ReleaseStatusPending    ReleaseStatus = "Pending"
	ReleaseStatusReleased   ReleaseStatus = "Released"
	ReleaseStatusRolledBack ReleaseStatus = "RolledBack"
	ReleaseStatusSuperseded ReleaseStatus = "Superseded"
)

// Keys for well known values.
const (
	ValKeyHeader     = "header"
	ValKeyHost       = "host"
	ValKeyMethod     = "method"
	ValKeyPath       = "path"
	ValKeyQuery      = "queryParam"
	ValKeySpanId     = "spanId"
	ValKeyStatus     = "status"
	ValKeyStatusCode = "statusCode"
	ValKeyTraceFlags = "traceFlags"
	ValKeyTraceId    = "traceId"
	ValKeyURL        = "url"
)

// Headers and query params.
const (
	HeaderAbbrvDep       = "kf-dep"
	HeaderAbbrvEnv       = "kf-env"
	HeaderAbbrvEventType = "kf-type"
	HeaderAdapter        = "kubefox-adapter"
	HeaderContentLength  = "Content-Length"
	HeaderDep            = "kubefox-deployment"
	HeaderEnv            = "kubefox-environment"
	HeaderEventType      = "kubefox-type"
	HeaderHost           = "Host"
	HeaderShortDep       = "kfd"
	HeaderShortEnv       = "kfe"
	HeaderShortEventType = "kft"
	HeaderTraceId        = "kubefox-trace-id"
)

const (
	CharSetUTF8 = "charset=UTF-8"

	DataSchemaEvent = "kubefox.proto.v1.Event"

	ContentTypeHTML     = "text/html"
	ContentTypeJSON     = "application/json"
	ContentTypePlain    = "text/plain"
	ContentTypeProtobuf = "application/protobuf"
)

var (
	RegexpCommit = regexp.MustCompile(`^[0-9a-f]{40}$`)
	RegexpGitRef = regexp.MustCompile(`^[a-z0-9][a-z0-9-\\.]{0,28}[a-z0-9]$`)
	RegexpImage  = regexp.MustCompile(`^.*:[a-z0-9-]{40}$`)
	RegexpName   = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,28}[a-z0-9]$`)
	RegexpUUID   = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

var (
	KubeFoxHome = utils.EnvDef("KUBEFOX_HOME", path.Join("/", "tmp", "kubefox"))

	FileCACert  = "ca.crt"
	FileTLSCert = "tls.crt"
	FileTLSKey  = "tls.key"

	PathCACert      = path.Join(KubeFoxHome, FileCACert)
	PathSvcAccToken = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	PathTLSCert     = path.Join(KubeFoxHome, FileTLSCert)
	PathTLSKey      = path.Join(KubeFoxHome, FileTLSKey)
)

const (
	DefaultRouteId = -1
)
