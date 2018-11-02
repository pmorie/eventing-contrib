// Package genomics provides access to the Genomics API.
//
// See https://cloud.google.com/genomics
//
// Usage example:
//
//   import "google.golang.org/api/genomics/v2alpha1"
//   ...
//   genomicsService, err := genomics.New(oauthHttpClient)
package genomics // import "google.golang.org/api/genomics/v2alpha1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "genomics:v2alpha1"
const apiName = "genomics"
const apiVersion = "v2alpha1"
const basePath = "https://genomics.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage Genomics data
	GenomicsScope = "https://www.googleapis.com/auth/genomics"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Pipelines = NewPipelinesService(s)
	s.Projects = NewProjectsService(s)
	s.Workers = NewWorkersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Pipelines *PipelinesService

	Projects *ProjectsService

	Workers *WorkersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewPipelinesService(s *Service) *PipelinesService {
	rs := &PipelinesService{s: s}
	return rs
}

type PipelinesService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Operations = NewProjectsOperationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Operations *ProjectsOperationsService
}

func NewProjectsOperationsService(s *Service) *ProjectsOperationsService {
	rs := &ProjectsOperationsService{s: s}
	return rs
}

type ProjectsOperationsService struct {
	s *Service
}

func NewWorkersService(s *Service) *WorkersService {
	rs := &WorkersService{s: s}
	return rs
}

type WorkersService struct {
	s *Service
}

// Accelerator: Carries information about an accelerator that can be
// attached to a VM.
type Accelerator struct {
	// Count: How many accelerators of this type to attach.
	Count int64 `json:"count,omitempty,string"`

	// Type: The accelerator type string (for example,
	// "nvidia-tesla-k80").
	//
	// Only NVIDIA GPU accelerators are currently supported. If an NVIDIA
	// GPU is
	// attached, the required runtime libraries will be made available to
	// all
	// containers under `/usr/local/nvidia`. The driver version to install
	// must
	// be specified using the NVIDIA driver version parameter on the
	// virtual
	// machine specification. Note that attaching a GPU increases the worker
	// VM
	// startup time by a few minutes.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Accelerator) MarshalJSON() ([]byte, error) {
	type NoMethod Accelerator
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Action: Specifies a single action that runs a Docker container.
type Action struct {
	// Commands: If specified, overrides the `CMD` specified in the
	// container. If the
	// container also has an `ENTRYPOINT` the values are used as
	// entrypoint
	// arguments. Otherwise, they are used as a command and arguments to
	// run
	// inside the container.
	Commands []string `json:"commands,omitempty"`

	// Credentials: If the specified image is hosted on a private registry
	// other than Google
	// Container Registry, the credentials required to pull the image must
	// be
	// specified here as an encrypted secret.
	//
	// The secret must decrypt to a JSON-encoded dictionary containing
	// both
	// `username` and `password` keys.
	Credentials *Secret `json:"credentials,omitempty"`

	// Entrypoint: If specified, overrides the `ENTRYPOINT` specified in the
	// container.
	Entrypoint string `json:"entrypoint,omitempty"`

	// Environment: The environment to pass into the container. This
	// environment is merged
	// with any values specified in the `Pipeline` message. These
	// values
	// overwrite any in the `Pipeline` message.
	//
	// In addition to the values passed here, a few other values
	// are
	// automatically injected into the environment. These cannot be hidden
	// or
	// overwritten.
	//
	// `GOOGLE_PIPELINE_FAILED` will be set to "1" if the pipeline
	// failed
	// because an action has exited with a non-zero status (and did not have
	// the
	// `IGNORE_EXIT_STATUS` flag set). This can be used to determine if
	// additional
	// debug or logging actions should execute.
	//
	// `GOOGLE_LAST_EXIT_STATUS` will be set to the exit status of the
	// last
	// non-background action that executed. This can be used by workflow
	// engine
	// authors to determine whether an individual action has succeeded or
	// failed.
	Environment map[string]string `json:"environment,omitempty"`

	// Flags: The set of flags to apply to this action.
	//
	// Possible values:
	//   "FLAG_UNSPECIFIED" - Unspecified flag.
	//   "IGNORE_EXIT_STATUS" - Normally, a non-zero exit status causes the
	// pipeline to fail. This flag
	// allows execution of other actions to continue instead.
	//   "RUN_IN_BACKGROUND" - This flag allows an action to continue
	// running in the background while
	// executing subsequent actions. This is useful to provide services
	// to
	// other actions (or to provide debugging support tools like SSH
	// servers).
	//   "ALWAYS_RUN" - By default, after an action fails, no further
	// actions are run. This flag
	// indicates that this action must be run even if the pipeline has
	// already
	// failed. This is useful for actions that copy output files off of the
	// VM
	// or for debugging.
	//   "ENABLE_FUSE" - Enable access to the FUSE device for this action.
	// Filesystems can then
	// be mounted into disks shared with other actions. The other actions
	// do
	// not need the `ENABLE_FUSE` flag to access the mounted
	// filesystem.
	//
	// This has the effect of causing the container to be executed
	// with
	// `CAP_SYS_ADMIN` and exposes `/dev/fuse` to the container, so use it
	// only
	// for containers you trust.
	//   "PUBLISH_EXPOSED_PORTS" - Exposes all ports specified by `EXPOSE`
	// statements in the container. To
	// discover the host side port numbers, consult the `ACTION_STARTED`
	// event
	// in the operation metadata.
	//   "DISABLE_IMAGE_PREFETCH" - All container images are typically
	// downloaded before any actions are
	// executed. This helps prevent typos in URIs or issues like lack of
	// disk
	// space from wasting large amounts of compute resources.
	//
	// If set, this flag prevents the worker from downloading the image
	// until
	// just before the action is executed.
	//   "DISABLE_STANDARD_ERROR_CAPTURE" - A small portion of the
	// container's standard error stream is typically
	// captured and returned inside the `ContainerStoppedEvent`. Setting
	// this
	// flag disables this functionality.
	Flags []string `json:"flags,omitempty"`

	// ImageUri: The URI to pull the container image from. Note that all
	// images referenced
	// by actions in the pipeline are pulled before the first action runs.
	// If
	// multiple actions reference the same image, it is only pulled
	// once,
	// ensuring that the same image is used for all actions in a single
	// pipeline.
	ImageUri string `json:"imageUri,omitempty"`

	// Labels: Labels to associate with the action. This field is provided
	// to assist
	// workflow engine authors in identifying actions (for example, to
	// indicate
	// what sort of action they perform, such as localization or
	// debugging).
	// They are returned in the operation metadata, but are otherwise
	// ignored.
	Labels map[string]string `json:"labels,omitempty"`

	// Mounts: A list of mounts to make available to the action.
	//
	// In addition to the values specified here, every action has a
	// special
	// virtual disk mounted under `/google` that contains log files and
	// other
	// operational components.
	//
	// <ul>
	//   <li><code>/google/logs</code> All logs written during the pipeline
	//   execution.</li>
	//   <li><code>/google/logs/output</code> The combined standard output
	// and
	//   standard error of all actions run as part of the pipeline
	//   execution.</li>
	//   <li><code>/google/logs/action/*/stdout</code> The complete contents
	// of
	//   each individual action's standard output.</li>
	//   <li><code>/google/logs/action/*/stderr</code> The complete contents
	// of
	//   each individual action's standard error output.</li>
	// </ul>
	Mounts []*Mount `json:"mounts,omitempty"`

	// Name: An optional name for the container. The container hostname will
	// be set to
	// this name, making it useful for inter-container communication. The
	// name
	// must contain only upper and lowercase alphanumeric characters and
	// hypens
	// and cannot start with a hypen.
	Name string `json:"name,omitempty"`

	// PidNamespace: The PID namespace to run the action inside. If
	// unspecified, a separate
	// isolated namespace is used.
	PidNamespace string `json:"pidNamespace,omitempty"`

	// PortMappings: A map of containers to host port mappings for this
	// container. If the
	// container already specifies exposed ports, use
	// the
	// `PUBLISH_EXPOSED_PORTS` flag instead.
	//
	// The host port number must be less than 65536. If it is zero, an
	// unused
	// random port is assigned. To determine the resulting port number,
	// consult
	// the `ContainerStartedEvent` in the operation metadata.
	PortMappings map[string]int64 `json:"portMappings,omitempty"`

	// Timeout: The maximum amount of time to give the action to complete.
	// If the action
	// fails to complete before the timeout, it will be terminated and the
	// exit
	// status will be non-zero. The pipeline will continue or terminate
	// based
	// on the rules defined by the `ALWAYS_RUN` and `IGNORE_EXIT_STATUS`
	// flags.
	Timeout string `json:"timeout,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Commands") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Commands") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Action) MarshalJSON() ([]byte, error) {
	type NoMethod Action
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CancelOperationRequest: The request message for
// Operations.CancelOperation.
type CancelOperationRequest struct {
}

// CheckInRequest: The parameters to the CheckIn method.
type CheckInRequest struct {
	// DeadlineExpired: The deadline has expired and the worker needs more
	// time.
	DeadlineExpired *Empty `json:"deadlineExpired,omitempty"`

	// Event: A workflow specific event occurred.
	Event googleapi.RawMessage `json:"event,omitempty"`

	// Result: The operation has finished with the given result.
	Result *Status `json:"result,omitempty"`

	// WorkerStatus: Data about the status of the worker VM.
	WorkerStatus *WorkerStatus `json:"workerStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeadlineExpired") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeadlineExpired") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CheckInRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CheckInRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CheckInResponse: The response to the CheckIn method.
type CheckInResponse struct {
	// Deadline: The deadline by which the worker must request an extension.
	//  The backend
	// will allow for network transmission time and other delays, but the
	// worker
	// must attempt to transmit the extension request no later than the
	// deadline.
	Deadline string `json:"deadline,omitempty"`

	// Metadata: The metadata that describes the operation assigned to the
	// worker.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Deadline") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Deadline") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CheckInResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CheckInResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ComputeEngine: Describes a Compute Engine resource that is being
// managed by a running
// pipeline.
type ComputeEngine struct {
	// DiskNames: The names of the disks that were created for this
	// pipeline.
	DiskNames []string `json:"diskNames,omitempty"`

	// InstanceName: The instance on which the operation is running.
	InstanceName string `json:"instanceName,omitempty"`

	// MachineType: The machine type of the instance.
	MachineType string `json:"machineType,omitempty"`

	// Zone: The availability zone in which the instance resides.
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DiskNames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DiskNames") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ComputeEngine) MarshalJSON() ([]byte, error) {
	type NoMethod ComputeEngine
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ContainerKilledEvent: An event generated when a container is forcibly
// terminated by the
// worker. Currently, this only occurs when the container outlives
// the
// timeout specified by the user.
type ContainerKilledEvent struct {
	// ActionId: The numeric ID of the action that started the container.
	ActionId int64 `json:"actionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ContainerKilledEvent) MarshalJSON() ([]byte, error) {
	type NoMethod ContainerKilledEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ContainerStartedEvent: An event generated when a container starts.
type ContainerStartedEvent struct {
	// ActionId: The numeric ID of the action that started this container.
	ActionId int64 `json:"actionId,omitempty"`

	// IpAddress: The public IP address that can be used to connect to the
	// container. This
	// field is only populated when at least one port mapping is present. If
	// the
	// instance was created with a private address, this field will be empty
	// even
	// if port mappings exist.
	IpAddress string `json:"ipAddress,omitempty"`

	// PortMappings: The container-to-host port mappings installed for this
	// container. This
	// set will contain any ports exposed using the `PUBLISH_EXPOSED_PORTS`
	// flag
	// as well as any specified in the `Action` definition.
	PortMappings map[string]int64 `json:"portMappings,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ContainerStartedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod ContainerStartedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ContainerStoppedEvent: An event generated when a container exits.
type ContainerStoppedEvent struct {
	// ActionId: The numeric ID of the action that started this container.
	ActionId int64 `json:"actionId,omitempty"`

	// ExitStatus: The exit status of the container.
	ExitStatus int64 `json:"exitStatus,omitempty"`

	// Stderr: The tail end of any content written to standard error by the
	// container.
	// If the content emits large amounts of debugging noise or
	// contains
	// sensitive information, you can prevent the content from being printed
	// by
	// setting the `DISABLE_STANDARD_ERROR_CAPTURE` flag.
	//
	// Note that only a small amount of the end of the stream is captured
	// here.
	// The entire stream is stored in the `/google/logs` directory mounted
	// into
	// each action, and can be copied off the machine as described
	// elsewhere.
	Stderr string `json:"stderr,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ContainerStoppedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod ContainerStoppedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DelayedEvent: An event generated whenever a resource limitation or
// transient error
// delays execution of a pipeline that was otherwise ready to run.
type DelayedEvent struct {
	// Cause: A textual description of the cause of the delay. The string
	// can change
	// without notice because it is often generated by another service (such
	// as
	// Compute Engine).
	Cause string `json:"cause,omitempty"`

	// Metrics: If the delay was caused by a resource shortage, this field
	// lists the
	// Compute Engine metrics that are preventing this operation from
	// running
	// (for example, `CPUS` or `INSTANCES`). If the particular metric is
	// not
	// known, a single `UNKNOWN` metric will be present.
	Metrics []string `json:"metrics,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cause") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Cause") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DelayedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod DelayedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Disk: Carries information about a disk that can be attached to a
// VM.
//
// See https://cloud.google.com/compute/docs/disks/performance for
// more
// information about disk type, size, and performance considerations.
type Disk struct {
	// Name: A user-supplied name for the disk. Used when mounting the disk
	// into
	// actions. The name must contain only upper and lowercase
	// alphanumeric
	// characters and hypens and cannot start with a hypen.
	Name string `json:"name,omitempty"`

	// SizeGb: The size, in GB, of the disk to attach. If the size is
	// not
	// specified, a default is chosen to ensure reasonable I/O
	// performance.
	//
	// If the disk type is specified as `local-ssd`, multiple local drives
	// are
	// automatically combined to provide the requested size. Note, however,
	// that
	// each physical SSD is 375GB in size, and no more than 8 drives can
	// be
	// attached to a single instance.
	SizeGb int64 `json:"sizeGb,omitempty"`

	// SourceImage: An optional image to put on the disk before attaching it
	// to the VM.
	SourceImage string `json:"sourceImage,omitempty"`

	// Type: The Compute Engine disk type. If unspecified, `pd-standard` is
	// used.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Disk) MarshalJSON() ([]byte, error) {
	type NoMethod Disk
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DiskStatus: The status of a disk on a VM.
type DiskStatus struct {
	// FreeSpaceBytes: Free disk space.
	FreeSpaceBytes uint64 `json:"freeSpaceBytes,omitempty,string"`

	// TotalSpaceBytes: Total disk space.
	TotalSpaceBytes uint64 `json:"totalSpaceBytes,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "FreeSpaceBytes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FreeSpaceBytes") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DiskStatus) MarshalJSON() ([]byte, error) {
	type NoMethod DiskStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Event: Carries information about events that occur during pipeline
// execution.
type Event struct {
	// Description: A human-readable description of the event. Note that
	// these strings can
	// change at any time without notice. Any application logic must use
	// the
	// information in the `details` field.
	Description string `json:"description,omitempty"`

	// Details: Machine-readable details about the event.
	Details googleapi.RawMessage `json:"details,omitempty"`

	// Timestamp: The time at which the event occurred.
	Timestamp string `json:"timestamp,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Event) MarshalJSON() ([]byte, error) {
	type NoMethod Event
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FailedEvent: An event generated when the execution of a pipeline has
// failed. Note
// that other events can continue to occur after this event.
type FailedEvent struct {
	// Cause: The human-readable description of the cause of the failure.
	Cause string `json:"cause,omitempty"`

	// Code: The Google standard error code that best describes this
	// failure.
	//
	// Possible values:
	//   "OK" - Not an error; returned on success
	//
	// HTTP Mapping: 200 OK
	//   "CANCELLED" - The operation was cancelled, typically by the
	// caller.
	//
	// HTTP Mapping: 499 Client Closed Request
	//   "UNKNOWN" - Unknown error.  For example, this error may be returned
	// when
	// a `Status` value received from another address space belongs to
	// an error space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	//   "INVALID_ARGUMENT" - The client specified an invalid argument.
	// Note that this differs
	// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates
	// arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	//   "DEADLINE_EXCEEDED" - The deadline expired before the operation
	// could complete. For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example,
	// a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	//   "NOT_FOUND" - Some requested entity (e.g., file or directory) was
	// not found.
	//
	// Note to server developers: if a request is denied for an entire
	// class
	// of users, such as gradual feature rollout or undocumented
	// whitelist,
	// `NOT_FOUND` may be used. If a request is denied for some users
	// within
	// a class of users, such as user-based access control,
	// `PERMISSION_DENIED`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	//   "ALREADY_EXISTS" - The entity that a client attempted to create
	// (e.g., file or directory)
	// already exists.
	//
	// HTTP Mapping: 409 Conflict
	//   "PERMISSION_DENIED" - The caller does not have permission to
	// execute the specified
	// operation. `PERMISSION_DENIED` must not be used for rejections
	// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
	// instead for those errors). `PERMISSION_DENIED` must not be
	// used if the caller can not be identified (use
	// `UNAUTHENTICATED`
	// instead for those errors). This error code does not imply the
	// request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	//   "UNAUTHENTICATED" - The request does not have valid authentication
	// credentials for the
	// operation.
	//
	// HTTP Mapping: 401 Unauthorized
	//   "RESOURCE_EXHAUSTED" - Some resource has been exhausted, perhaps a
	// per-user quota, or
	// perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	//   "FAILED_PRECONDITION" - The operation was rejected because the
	// system is not in a state
	// required for the operation's execution.  For example, the
	// directory
	// to be deleted is non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to
	// decide
	// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
	//  (a) Use `UNAVAILABLE` if the client can retry just the failing
	// call.
	//  (b) Use `ABORTED` if the client should retry at a higher level
	//      (e.g., when a client-specified test-and-set fails, indicating
	// the
	//      client should restart a read-modify-write sequence).
	//  (c) Use `FAILED_PRECONDITION` if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an
	// "rmdir"
	//      fails because the directory is non-empty, `FAILED_PRECONDITION`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	//   "ABORTED" - The operation was aborted, typically due to a
	// concurrency issue such as
	// a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between
	// `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 409 Conflict
	//   "OUT_OF_RANGE" - The operation was attempted past the valid range.
	// E.g., seeking or
	// reading past end-of-file.
	//
	// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit
	// file
	// system will generate `INVALID_ARGUMENT` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will
	// generate
	// `OUT_OF_RANGE` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `FAILED_PRECONDITION`
	// and
	// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more
	// specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `OUT_OF_RANGE` error to detect
	// when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	//   "UNIMPLEMENTED" - The operation is not implemented or is not
	// supported/enabled in this
	// service.
	//
	// HTTP Mapping: 501 Not Implemented
	//   "INTERNAL" - Internal errors.  This means that some invariants
	// expected by the
	// underlying system have been broken.  This error code is reserved
	// for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	//   "UNAVAILABLE" - The service is currently unavailable.  This is most
	// likely a
	// transient condition, which can be corrected by retrying with
	// a backoff.
	//
	// See the guidelines above for deciding between
	// `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 503 Service Unavailable
	//   "DATA_LOSS" - Unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code string `json:"code,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cause") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Cause") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FailedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod FailedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ImportReadGroupSetsResponse: The read group set import response.
type ImportReadGroupSetsResponse struct {
	// ReadGroupSetIds: IDs of the read group sets that were created.
	ReadGroupSetIds []string `json:"readGroupSetIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReadGroupSetIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ReadGroupSetIds") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ImportReadGroupSetsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ImportReadGroupSetsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ImportVariantsResponse: The variant data import response.
type ImportVariantsResponse struct {
	// CallSetIds: IDs of the call sets created during the import.
	CallSetIds []string `json:"callSetIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CallSetIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CallSetIds") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ImportVariantsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ImportVariantsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOperationsResponse: The response message for
// Operations.ListOperations.
type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListOperationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListOperationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metadata: Carries information about the pipeline execution that is
// returned
// in the long running operation's metadata field.
type Metadata struct {
	// CreateTime: The time at which the operation was created by the API.
	CreateTime string `json:"createTime,omitempty"`

	// EndTime: The time at which execution was completed and resources were
	// cleaned up.
	EndTime string `json:"endTime,omitempty"`

	// Events: The list of events that have happened so far during the
	// execution of this
	// operation.
	Events []*Event `json:"events,omitempty"`

	// Labels: The user-defined labels associated with this operation.
	Labels map[string]string `json:"labels,omitempty"`

	// Pipeline: The pipeline this operation represents.
	Pipeline *Pipeline `json:"pipeline,omitempty"`

	// StartTime: The first time at which resources were allocated to
	// execute the pipeline.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Metadata) MarshalJSON() ([]byte, error) {
	type NoMethod Metadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Mount: Carries information about a particular disk mount inside a
// container.
type Mount struct {
	// Disk: The name of the disk to mount, as specified in the resources
	// section.
	Disk string `json:"disk,omitempty"`

	// Path: The path to mount the disk inside the container.
	Path string `json:"path,omitempty"`

	// ReadOnly: If true, the disk is mounted read-only inside the
	// container.
	ReadOnly bool `json:"readOnly,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Disk") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Disk") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Mount) MarshalJSON() ([]byte, error) {
	type NoMethod Mount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Network: VM networking options.
type Network struct {
	// Name: The network name to attach the VM's network interface to. The
	// value will
	// be prefixed with `global/networks/` unless it contains a `/`, in
	// which
	// case it is assumed to be a fully specified network resource URL.
	//
	// If unspecified, the global default network is used.
	Name string `json:"name,omitempty"`

	// Subnetwork: If the specified network is configured for custom subnet
	// creation, the
	// name of the subnetwork to attach the instance to must be specified
	// here.
	//
	// The value is prefixed with `regions/*/subnetworks/` unless it
	// contains a
	// `/`, in which case it is assumed to be a fully specified
	// subnetwork
	// resource URL.
	//
	// If the `*` character appears in the value, it is replaced with the
	// region
	// that the virtual machine has been allocated in.
	Subnetwork string `json:"subnetwork,omitempty"`

	// UsePrivateAddress: If set to true, do not attach a public IP address
	// to the VM. Note that
	// without a public IP address, additional configuration is required
	// to
	// allow the VM to access Google services.
	//
	// See
	// https://cloud.google.com/vpc/docs/configure-private-google-access
	// for more information.
	UsePrivateAddress bool `json:"usePrivateAddress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Network) MarshalJSON() ([]byte, error) {
	type NoMethod Network
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: An OperationMetadata or Metadata object. This will always
	// be returned with the Operation.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that originally returns it. For example&#58;
	// `operations/CJHU7Oi_ChDrveSpBRjfuL-qzoWAgEw`
	Name string `json:"name,omitempty"`

	// Response: If importing ReadGroupSets, an ImportReadGroupSetsResponse
	// is returned. If importing Variants, an ImportVariantsResponse is
	// returned. For pipelines and exports, an Empty response is returned.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OperationEvent: An event that occurred during an Operation.
type OperationEvent struct {
	// Description: Required description of event.
	Description string `json:"description,omitempty"`

	// EndTime: Optional time of when event finished. An event can have a
	// start time and no
	// finish time. If an event has a finish time, there must be a start
	// time.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Optional time of when event started.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationEvent) MarshalJSON() ([]byte, error) {
	type NoMethod OperationEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OperationMetadata: Metadata describing an Operation.
type OperationMetadata struct {
	// ClientId: This field is deprecated. Use `labels` instead. Optionally
	// provided by the
	// caller when submitting the request that creates the operation.
	ClientId string `json:"clientId,omitempty"`

	// CreateTime: The time at which the job was submitted to the Genomics
	// service.
	CreateTime string `json:"createTime,omitempty"`

	// EndTime: The time at which the job stopped running.
	EndTime string `json:"endTime,omitempty"`

	// Events: Optional event messages that were generated during the job's
	// execution.
	// This also contains any warnings that were generated during import
	// or export.
	Events []*OperationEvent `json:"events,omitempty"`

	// Labels: Optionally provided by the caller when submitting the request
	// that creates
	// the operation.
	Labels map[string]string `json:"labels,omitempty"`

	// ProjectId: The Google Cloud Project in which the job is scoped.
	ProjectId string `json:"projectId,omitempty"`

	// Request: The original request that started the operation. Note that
	// this will be in
	// current version of the API. If the operation was started with v1beta2
	// API
	// and a GetOperation is performed on v1 API, a v1 request will be
	// returned.
	Request googleapi.RawMessage `json:"request,omitempty"`

	// RuntimeMetadata: Runtime metadata on this Operation.
	RuntimeMetadata googleapi.RawMessage `json:"runtimeMetadata,omitempty"`

	// StartTime: The time at which the job began to run.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClientId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod OperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Pipeline: Specifies a series of actions to execute, expressed as
// Docker containers.
type Pipeline struct {
	// Actions: The list of actions to execute, in the order they are
	// specified.
	Actions []*Action `json:"actions,omitempty"`

	// Environment: The environment to pass into every action. Each action
	// can also specify
	// additional environment variables but cannot delete an entry from this
	// map
	// (though they can overwrite it with a different value).
	Environment map[string]string `json:"environment,omitempty"`

	// Resources: The resources required for execution.
	Resources *Resources `json:"resources,omitempty"`

	// Timeout: The maximum amount of time to give the pipeline to complete.
	//  This includes
	// the time spent waiting for a worker to be allocated.  If the pipeline
	// fails
	// to complete before the timeout, it will be cancelled and the error
	// code
	// will be set to DEADLINE_EXCEEDED.
	//
	// If unspecified, it will default to 7 days.
	Timeout string `json:"timeout,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Actions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Actions") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Pipeline) MarshalJSON() ([]byte, error) {
	type NoMethod Pipeline
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PullStartedEvent: An event generated when the worker starts pulling
// an image.
type PullStartedEvent struct {
	// ImageUri: The URI of the image that was pulled.
	ImageUri string `json:"imageUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PullStartedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod PullStartedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PullStoppedEvent: An event generated when the worker stops pulling an
// image.
type PullStoppedEvent struct {
	// ImageUri: The URI of the image that was pulled.
	ImageUri string `json:"imageUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PullStoppedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod PullStoppedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Resources: The system resources for the pipeline run.
//
// At least one zone or region must be specified or the pipeline run
// will fail.
type Resources struct {
	// ProjectId: The project ID to allocate resources in.
	ProjectId string `json:"projectId,omitempty"`

	// Regions: The list of regions allowed for VM allocation. If set, the
	// `zones` field
	// must not be set.
	Regions []string `json:"regions,omitempty"`

	// VirtualMachine: The virtual machine specification.
	VirtualMachine *VirtualMachine `json:"virtualMachine,omitempty"`

	// Zones: The list of zones allowed for VM allocation. If set, the
	// `regions` field
	// must not be set.
	Zones []string `json:"zones,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ProjectId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ProjectId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Resources) MarshalJSON() ([]byte, error) {
	type NoMethod Resources
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RunPipelineRequest: The arguments to the `RunPipeline` method. The
// requesting user must have
// the `iam.serviceAccounts.actAs` permission for the Cloud Genomics
// service
// account or the request will fail.
type RunPipelineRequest struct {
	// Labels: User-defined labels to associate with the returned operation.
	// These
	// labels are not propagated to any Google Cloud Platform resources used
	// by
	// the operation, and can be modified at any time.
	//
	// To associate labels with resources created while executing the
	// operation,
	// see the appropriate resource message (for example, `VirtualMachine`).
	Labels map[string]string `json:"labels,omitempty"`

	// Pipeline: The description of the pipeline to run.
	Pipeline *Pipeline `json:"pipeline,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Labels") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RunPipelineRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RunPipelineRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RunPipelineResponse: The response to the RunPipeline method, returned
// in the operation's result
// field on success.
type RunPipelineResponse struct {
}

// RuntimeMetadata: Runtime metadata that will be populated in
// the
// runtimeMetadata
// field of the Operation associated with a RunPipeline execution.
type RuntimeMetadata struct {
	// ComputeEngine: Execution information specific to Google Compute
	// Engine.
	ComputeEngine *ComputeEngine `json:"computeEngine,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ComputeEngine") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ComputeEngine") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RuntimeMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod RuntimeMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Secret: Holds encrypted information that is only decrypted and stored
// in RAM
// by the worker VM when running the pipeline.
type Secret struct {
	// CipherText: The value of the cipherText response from the `encrypt`
	// method. This field
	// is intentionally unaudited.
	CipherText string `json:"cipherText,omitempty"`

	// KeyName: The name of the Cloud KMS key that will be used to decrypt
	// the secret
	// value. The VM service account must have the required permissions
	// and
	// authentication scopes to invoke the `decrypt` method on the specified
	// key.
	KeyName string `json:"keyName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CipherText") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CipherText") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Secret) MarshalJSON() ([]byte, error) {
	type NoMethod Secret
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ServiceAccount: Carries information about a Google Cloud service
// account.
type ServiceAccount struct {
	// Email: Email address of the service account. If not specified, the
	// default
	// Compute Engine service account for the project will be used.
	Email string `json:"email,omitempty"`

	// Scopes: List of scopes to be enabled for this service account on the
	// VM, in
	// addition to the Cloud Genomics API scope.
	Scopes []string `json:"scopes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Email") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Email") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ServiceAccount) MarshalJSON() ([]byte, error) {
	type NoMethod ServiceAccount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UnexpectedExitStatusEvent: An event generated when the execution of a
// container results in a
// non-zero exit status that was not otherwise ignored. Execution
// will
// continue, but only actions that are flagged as `ALWAYS_RUN` will
// be
// executed. Other actions will be skipped.
type UnexpectedExitStatusEvent struct {
	// ActionId: The numeric ID of the action that started the container.
	ActionId int64 `json:"actionId,omitempty"`

	// ExitStatus: The exit status of the container.
	ExitStatus int64 `json:"exitStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UnexpectedExitStatusEvent) MarshalJSON() ([]byte, error) {
	type NoMethod UnexpectedExitStatusEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VirtualMachine: Carries information about a Compute Engine VM
// resource.
type VirtualMachine struct {
	// Accelerators: The list of accelerators to attach to the VM.
	Accelerators []*Accelerator `json:"accelerators,omitempty"`

	// BootDiskSizeGb: The size of the boot disk, in GB. The boot disk must
	// be large
	// enough to accommodate all of the Docker images from each action in
	// the
	// pipeline at the same time. If not specified, a small but
	// reasonable
	// default value is used.
	BootDiskSizeGb int64 `json:"bootDiskSizeGb,omitempty"`

	// BootImage: The host operating system image to use.
	//
	// Currently, only Container-Optimized OS images can be used.
	//
	// The default value is
	// `projects/cos-cloud/global/images/family/cos-stable`,
	// which selects the latest stable release of Container-Optimized
	// OS.
	//
	// This option is provided to allow testing against the beta release of
	// the
	// operating system to ensure that the new version does not
	// interact
	// negatively with production pipelines.
	//
	// To test a pipeline against the beta release of Container-Optimized
	// OS,
	// use the value `projects/cos-cloud/global/images/family/cos-beta`.
	BootImage string `json:"bootImage,omitempty"`

	// CpuPlatform: The CPU platform to request. An instance based on a
	// newer platform can be
	// allocated, but never one with fewer capabilities. The value of
	// this
	// parameter must be a valid Compute Engine CPU platform name (such as
	// "Intel
	// Skylake"). This parameter is only useful for carefully optimized
	// work
	// loads where the CPU platform has a significant impact.
	//
	// For more information about the effect of this parameter,
	// see
	// https://cloud.google.com/compute/docs/instances/specify-min-cpu-pl
	// atform.
	CpuPlatform string `json:"cpuPlatform,omitempty"`

	// Disks: The list of disks to create and attach to the VM.
	Disks []*Disk `json:"disks,omitempty"`

	// Labels: Optional set of labels to apply to the VM and any attached
	// disk resources.
	// These labels must adhere to the name and value restrictions on VM
	// labels
	// imposed by Compute Engine.
	//
	// Labels applied at creation time to the VM. Applied on a best-effort
	// basis
	// to attached disk resources shortly after VM creation.
	Labels map[string]string `json:"labels,omitempty"`

	// MachineType: The machine type of the virtual machine to create. Must
	// be the short name
	// of a standard machine type (such as "n1-standard-1") or a custom
	// machine
	// type (such as "custom-1-4096").
	MachineType string `json:"machineType,omitempty"`

	// Network: The VM network configuration.
	Network *Network `json:"network,omitempty"`

	// NvidiaDriverVersion: The NVIDIA driver version to use when attaching
	// an NVIDIA GPU accelerator.
	// The version specified here must be compatible with the GPU
	// libraries
	// contained in the container being executed, and must be one of the
	// drivers
	// hosted in the `nvidia-drivers-us-public` bucket on Google Cloud
	// Storage.
	NvidiaDriverVersion string `json:"nvidiaDriverVersion,omitempty"`

	// Preemptible: If true, allocate a preemptible VM.
	Preemptible bool `json:"preemptible,omitempty"`

	// ServiceAccount: The service account to install on the VM. This
	// account does not need
	// any permissions other than those required by the pipeline.
	ServiceAccount *ServiceAccount `json:"serviceAccount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Accelerators") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Accelerators") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VirtualMachine) MarshalJSON() ([]byte, error) {
	type NoMethod VirtualMachine
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkerAssignedEvent: An event generated after a worker VM has been
// assigned to run the
// pipeline.
type WorkerAssignedEvent struct {
	// Instance: The worker's instance name.
	Instance string `json:"instance,omitempty"`

	// Zone: The zone the worker is running in.
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Instance") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Instance") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WorkerAssignedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod WorkerAssignedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkerReleasedEvent: An event generated when the worker VM that was
// assigned to the pipeline
// has been released (deleted).
type WorkerReleasedEvent struct {
	// Instance: The worker's instance name.
	Instance string `json:"instance,omitempty"`

	// Zone: The zone the worker was running in.
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Instance") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Instance") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WorkerReleasedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod WorkerReleasedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WorkerStatus: The status of the worker VM.
type WorkerStatus struct {
	// AttachedDisks: Status of attached disks.
	AttachedDisks map[string]DiskStatus `json:"attachedDisks,omitempty"`

	// BootDisk: Status of the boot disk.
	BootDisk *DiskStatus `json:"bootDisk,omitempty"`

	// FreeRamBytes: Free RAM.
	FreeRamBytes uint64 `json:"freeRamBytes,omitempty,string"`

	// TotalRamBytes: Total RAM.
	TotalRamBytes uint64 `json:"totalRamBytes,omitempty,string"`

	// UptimeSeconds: System uptime.
	UptimeSeconds int64 `json:"uptimeSeconds,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "AttachedDisks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AttachedDisks") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WorkerStatus) MarshalJSON() ([]byte, error) {
	type NoMethod WorkerStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "genomics.pipelines.run":

type PipelinesRunCall struct {
	s                  *Service
	runpipelinerequest *RunPipelineRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Run: Runs a pipeline.
//
// **Note:** Before you can use this method, the Genomics Service
// Agent
// must have access to your project. This is done automatically when
// the
// Cloud Genomics API is first enabled, but if you delete this
// permission,
// or if you enabled the Cloud Genomics API before the v2alpha1
// API
// launch, you must disable and re-enable the API to grant the
// Genomics
// Service Agent the required permissions.
// Authorization requires the following
// [Google
// IAM](https://cloud.google.com/iam/) permission:
//
// * `genomics.operations.create`
//
// [1]: /genomics/gsa
func (r *PipelinesService) Run(runpipelinerequest *RunPipelineRequest) *PipelinesRunCall {
	c := &PipelinesRunCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.runpipelinerequest = runpipelinerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PipelinesRunCall) Fields(s ...googleapi.Field) *PipelinesRunCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PipelinesRunCall) Context(ctx context.Context) *PipelinesRunCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PipelinesRunCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PipelinesRunCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runpipelinerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2alpha1/pipelines:run")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "genomics.pipelines.run" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *PipelinesRunCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs a pipeline.\n\n**Note:** Before you can use this method, the Genomics Service Agent\nmust have access to your project. This is done automatically when the\nCloud Genomics API is first enabled, but if you delete this permission,\nor if you enabled the Cloud Genomics API before the v2alpha1 API\nlaunch, you must disable and re-enable the API to grant the Genomics\nService Agent the required permissions.\nAuthorization requires the following [Google\nIAM](https://cloud.google.com/iam/) permission:\n\n* `genomics.operations.create`\n\n[1]: /genomics/gsa",
	//   "flatPath": "v2alpha1/pipelines:run",
	//   "httpMethod": "POST",
	//   "id": "genomics.pipelines.run",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2alpha1/pipelines:run",
	//   "request": {
	//     "$ref": "RunPipelineRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.projects.operations.cancel":

type ProjectsOperationsCancelCall struct {
	s                      *Service
	name                   string
	canceloperationrequest *CancelOperationRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Cancel: Starts asynchronous cancellation on a long-running
// operation.
// The server makes a best effort to cancel the operation, but success
// is not
// guaranteed. Clients may use Operations.GetOperation
// or Operations.ListOperations
// to check whether the cancellation succeeded or the operation
// completed
// despite cancellation.
// Authorization requires the following [Google
// IAM](https://cloud.google.com/iam) permission&#58;
//
// * `genomics.operations.cancel`
func (r *ProjectsOperationsService) Cancel(name string, canceloperationrequest *CancelOperationRequest) *ProjectsOperationsCancelCall {
	c := &ProjectsOperationsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.canceloperationrequest = canceloperationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsCancelCall) Fields(s ...googleapi.Field) *ProjectsOperationsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsCancelCall) Context(ctx context.Context) *ProjectsOperationsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.canceloperationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2alpha1/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "genomics.projects.operations.cancel" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsOperationsCancelCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts asynchronous cancellation on a long-running operation.\nThe server makes a best effort to cancel the operation, but success is not\nguaranteed. Clients may use Operations.GetOperation\nor Operations.ListOperations\nto check whether the cancellation succeeded or the operation completed\ndespite cancellation.\nAuthorization requires the following [Google IAM](https://cloud.google.com/iam) permission\u0026#58;\n\n* `genomics.operations.cancel`",
	//   "flatPath": "v2alpha1/projects/{projectsId}/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "genomics.projects.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2alpha1/{+name}:cancel",
	//   "request": {
	//     "$ref": "CancelOperationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.projects.operations.get":

type ProjectsOperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.
// Clients can use this method to poll the operation result at intervals
// as
// recommended by the API service.
// Authorization requires the following [Google
// IAM](https://cloud.google.com/iam) permission&#58;
//
// * `genomics.operations.get`
func (r *ProjectsOperationsService) Get(name string) *ProjectsOperationsGetCall {
	c := &ProjectsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsGetCall) Context(ctx context.Context) *ProjectsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2alpha1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "genomics.projects.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsOperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.\nClients can use this method to poll the operation result at intervals as\nrecommended by the API service.\nAuthorization requires the following [Google IAM](https://cloud.google.com/iam) permission\u0026#58;\n\n* `genomics.operations.get`",
	//   "flatPath": "v2alpha1/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "genomics.projects.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2alpha1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.projects.operations.list":

type ProjectsOperationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists operations that match the specified filter in the
// request.
// Authorization requires the following [Google
// IAM](https://cloud.google.com/iam) permission&#58;
//
// * `genomics.operations.list`
func (r *ProjectsOperationsService) List(name string) *ProjectsOperationsListCall {
	c := &ProjectsOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": A string for filtering
// Operations.
// In v2alpha1, the following filter fields are supported&#58;
//
// * createTime&#58; The time this job was created
// * events&#58; The set of event (names) that have occurred while
// running
//   the pipeline.  The &#58; operator can be used to determine if a
//   particular event has occurred.
// * error&#58; If the pipeline is running, this value is NULL.  Once
// the
//   pipeline finishes, the value is the standard Google error code.
// * labels.key or labels."key with space" where key is a label key.
// * done&#58; If the pipeline is running, this value is false. Once
// the
//   pipeline finishes, the value is true.
//
// In v1 and v1alpha2, the following filter fields are supported&#58;
//
// * projectId&#58; Required. Corresponds to
//   OperationMetadata.projectId.
// * createTime&#58; The time this job was created, in seconds from the
//   [epoch](http://en.wikipedia.org/wiki/Unix_time). Can use `>=`
// and/or `<=`
//   operators.
// * status&#58; Can be `RUNNING`, `SUCCESS`, `FAILURE`, or `CANCELED`.
// Only
//   one status may be specified.
// * labels.key where key is a label key.
//
// Examples&#58;
//
// * `projectId = my-project AND createTime >= 1432140000`
// * `projectId = my-project AND createTime >= 1432140000 AND createTime
// <= 1432150000 AND status = RUNNING`
// * `projectId = my-project AND labels.color = *`
// * `projectId = my-project AND labels.color = red`
func (c *ProjectsOperationsListCall) Filter(filter string) *ProjectsOperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return. If unspecified, defaults to
// 256. The maximum value is 2048.
func (c *ProjectsOperationsListCall) PageSize(pageSize int64) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsOperationsListCall) PageToken(pageToken string) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsListCall) IfNoneMatch(entityTag string) *ProjectsOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsListCall) Context(ctx context.Context) *ProjectsOperationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2alpha1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "genomics.projects.operations.list" call.
// Exactly one of *ListOperationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOperationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsListCall) Do(opts ...googleapi.CallOption) (*ListOperationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOperationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request.\nAuthorization requires the following [Google IAM](https://cloud.google.com/iam) permission\u0026#58;\n\n* `genomics.operations.list`",
	//   "flatPath": "v2alpha1/projects/{projectsId}/operations",
	//   "httpMethod": "GET",
	//   "id": "genomics.projects.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A string for filtering Operations.\nIn v2alpha1, the following filter fields are supported\u0026#58;\n\n* createTime\u0026#58; The time this job was created\n* events\u0026#58; The set of event (names) that have occurred while running\n  the pipeline.  The \u0026#58; operator can be used to determine if a\n  particular event has occurred.\n* error\u0026#58; If the pipeline is running, this value is NULL.  Once the\n  pipeline finishes, the value is the standard Google error code.\n* labels.key or labels.\"key with space\" where key is a label key.\n* done\u0026#58; If the pipeline is running, this value is false. Once the\n  pipeline finishes, the value is true.\n\nIn v1 and v1alpha2, the following filter fields are supported\u0026#58;\n\n* projectId\u0026#58; Required. Corresponds to\n  OperationMetadata.projectId.\n* createTime\u0026#58; The time this job was created, in seconds from the\n  [epoch](http://en.wikipedia.org/wiki/Unix_time). Can use `\u003e=` and/or `\u003c=`\n  operators.\n* status\u0026#58; Can be `RUNNING`, `SUCCESS`, `FAILURE`, or `CANCELED`. Only\n  one status may be specified.\n* labels.key where key is a label key.\n\nExamples\u0026#58;\n\n* `projectId = my-project AND createTime \u003e= 1432140000`\n* `projectId = my-project AND createTime \u003e= 1432140000 AND createTime \u003c= 1432150000 AND status = RUNNING`\n* `projectId = my-project AND labels.color = *`\n* `projectId = my-project AND labels.color = red`",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the operation's parent resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of results to return. If unspecified, defaults to\n256. The maximum value is 2048.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2alpha1/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsOperationsListCall) Pages(ctx context.Context, f func(*ListOperationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "genomics.workers.checkIn":

type WorkersCheckInCall struct {
	s              *Service
	id             string
	checkinrequest *CheckInRequest
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// CheckIn: The worker uses this method to retrieve the assigned
// operation and
// provide periodic status updates.
func (r *WorkersService) CheckIn(id string, checkinrequest *CheckInRequest) *WorkersCheckInCall {
	c := &WorkersCheckInCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.id = id
	c.checkinrequest = checkinrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WorkersCheckInCall) Fields(s ...googleapi.Field) *WorkersCheckInCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *WorkersCheckInCall) Context(ctx context.Context) *WorkersCheckInCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *WorkersCheckInCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *WorkersCheckInCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.checkinrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2alpha1/workers/{id}:checkIn")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "genomics.workers.checkIn" call.
// Exactly one of *CheckInResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CheckInResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *WorkersCheckInCall) Do(opts ...googleapi.CallOption) (*CheckInResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CheckInResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "The worker uses this method to retrieve the assigned operation and\nprovide periodic status updates.",
	//   "flatPath": "v2alpha1/workers/{id}:checkIn",
	//   "httpMethod": "POST",
	//   "id": "genomics.workers.checkIn",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The worker id, assigned when it was created.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2alpha1/workers/{id}:checkIn",
	//   "request": {
	//     "$ref": "CheckInRequest"
	//   },
	//   "response": {
	//     "$ref": "CheckInResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}
