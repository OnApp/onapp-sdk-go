package onappgo

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/digitalocean/godo"
)

// CloudBoot, Smart CloudBoot, Baremetal CloudBoot - Create, Edit
const cloudBootComputeResourcesBasePath string = "settings/assets/%s/hypervisors"
const cloudBootAvailableResourcesBasePath string = "settings/assets"

// CloudbootComputeResourcesService is an interface for interfacing with the CloudbootComputeResource
// endpoints of the OnApp API
// See: https://docs.onapp.com/apim/latest/compute-resources
type CloudbootComputeResourcesService interface {
	List(context.Context, *ListOptions) ([]CloudbootComputeResource, *Response, error)
	Get(context.Context, int) (*CloudbootComputeResource, *Response, error)
	Create(context.Context, *CloudbootComputeResourceCreateRequest) (*CloudbootComputeResource, *Response, error)
	Delete(context.Context, int, interface{}) (*Response, error)
	Edit(context.Context, int, *CloudbootComputeResourceEditRequest) (*Response, error)

	CloudbootAvailableResources(context.Context) ([]Asset, *Response, error)
}

// CloudbootComputeResourcesServiceOp handles communication with the CloudbootComputeResource related methods of the
// OnApp API.
type CloudbootComputeResourcesServiceOp struct {
	client *Client
}

var _ CloudbootComputeResourcesService = &CloudbootComputeResourcesServiceOp{}

type CloudbootComputeResource Hypervisor

type Asset struct {
	Mac string `json:"mac,omitempty"`
	IP  string `json:"ip,omitempty"`
}

type StorageDisk struct {
	Scsi     string `json:"scsi,omitempty"`
	Selected bool   `json:"selected,bool"`
}
type StorageNic struct {
	Mac  string `json:"mac,omitempty"`
	Type int    `json:"type,omitempty"`
}
type StorageCustomPci struct {
	Pci      string `json:"pci,omitempty"`
	Selected bool   `json:"selected,bool"`
}
type Storage struct {
	Disks      []StorageDisk      `json:"disks,omitempty"`
	Nics       []StorageNic       `json:"nics,omitempty"`
	CustomPcis []StorageCustomPci `json:"custom_pcis,omitempty"`
}

type CloudbootComputeResourceCreateRequest struct {
	Label                       string   `json:"label,omitempty"`
	PxeIPAddressID              int      `json:"pxe_ip_address_id,omitempty"`
	HypervisorType              string   `json:"hypervisor_type,omitempty"`
	HypervisorGroupID           int      `json:"hypervisor_group_id,omitempty"`
	SegregationOsType           string   `json:"segregation_os_type,omitempty"`
	ServerType                  string   `json:"server_type,omitempty"`
	Backup                      bool     `json:"backup,bool"`
	BackupIPAddress             string   `json:"backup_ip_address,omitempty"`
	Enabled                     bool     `json:"enabled,bool"`
	CollectStats                bool     `json:"collect_stats,bool"`
	DisableFailover             bool     `json:"disable_failover,bool"`
	FormatDisks                 bool     `json:"format_disks,bool"`
	PassthroughDisks            bool     `json:"passthrough_disks,bool"`
	Storage                     *Storage `json:"storage,omitempty"`
	StorageControllerMemorySize int      `json:"storage_controller_memory_size,omitempty"`
	StaticIntegratedStorage     bool     `json:"static_integrated_storage,bool"`
	DisksPerStorageController   int      `json:"disks_per_storage_controller,omitempty"`
	CloudBootOs                 string   `json:"cloud_boot_os,omitempty"`
	CustomConfig                string   `json:"custom_config,omitempty"`
	DefaultGateway              string   `json:"default_gateway,omitempty"`
	Vlan                        string   `json:"vlan,omitempty"`
	Mac                         string   `json:"mac,omitempty"` // Helper field
}

// CloudbootComputeResourceEditRequest represents a request to edit a CloudbootComputeResource
type CloudbootComputeResourceEditRequest struct {
	CollectStats                     bool     `json:"collect_stats,bool"`
	DisableFailover                  bool     `json:"disable_failover,bool"`
	PassthroughDisks                 bool     `json:"passthrough_disks,bool"`
	Storage                          *Storage `json:"storage,omitempty"`
	StorageControllerMemorySize      int      `json:"storage_controller_memory_size,omitempty"`
	StaticIntegratedStorage          bool     `json:"static_integrated_storage,bool"`
	DisksPerStorageController        int      `json:"disks_per_storage_controller,omitempty"`
	IntegratedStorageDisabled        bool     `json:"integrated_storage_disabled,omitempty"`
	CustomConfig                     string   `json:"custom_config,omitempty"`
	ApplyHypervisorGroupCustomConfig bool     `json:"apply_hypervisor_group_custom_config,bool"`
}

type cloudbootComputeResourceCreateRequestRoot struct {
	CloudbootComputeResourceCreateRequest *CloudbootComputeResourceCreateRequest `json:"hypervisor"`
}

type cloudbootComputeResourceRoot struct {
	CloudbootComputeResource *CloudbootComputeResource `json:"hypervisor"`
}

func (d CloudbootComputeResourceCreateRequest) String() string {
	return godo.Stringify(d)
}

// List all Cloudboot CloudbootComputeResources
func (s *CloudbootComputeResourcesServiceOp) List(ctx context.Context, opt *ListOptions) ([]CloudbootComputeResource, *Response, error) {
	path := hypervisorsBasePath + apiFormat
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var out []map[string]CloudbootComputeResource
	resp, err := s.client.Do(ctx, req, &out)
	if err != nil {
		return nil, resp, err
	}

	arr := make([]CloudbootComputeResource, len(out))
	for i := range arr {
		arr[i] = out[i]["hypervisor"]
	}

	return arr, resp, err
}

// Get individual Cloudboot CloudbootComputeResource
func (s *CloudbootComputeResourcesServiceOp) Get(ctx context.Context, id int) (*CloudbootComputeResource, *Response, error) {
	if id < 1 {
		return nil, nil, godo.NewArgError("id", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/%d%s", hypervisorsBasePath, id, apiFormat)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(cloudbootComputeResourceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.CloudbootComputeResource, resp, err
}

// Create Cloudboot CloudbootComputeResource
func (s *CloudbootComputeResourcesServiceOp) Create(ctx context.Context, createRequest *CloudbootComputeResourceCreateRequest) (*CloudbootComputeResource, *Response, error) {
	if createRequest == nil {
		return nil, nil, godo.NewArgError("CloudbootComputeResource createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(cloudBootComputeResourcesBasePath, createRequest.Mac) + apiFormat
	rootRequest := &cloudbootComputeResourceCreateRequestRoot{
		CloudbootComputeResourceCreateRequest: createRequest,
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, rootRequest)
	if err != nil {
		return nil, nil, err
	}
	log.Println("CloudbootComputeResource [Create] req: ", req)

	root := new(cloudbootComputeResourceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.CloudbootComputeResource, resp, err
}

// Delete Cloudboot CloudbootComputeResource
func (s *CloudbootComputeResourcesServiceOp) Delete(ctx context.Context, id int, meta interface{}) (*Response, error) {
	if id < 1 {
		return nil, godo.NewArgError("id", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/%d%s", hypervisorsBasePath, id, apiFormat)
	path, err := addOptions(path, meta)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	log.Println("CloudbootComputeResource [Delete] req: ", req)

	return s.client.Do(ctx, req, nil)
}

// Edit Cloudboot CloudbootComputeResource
func (s *CloudbootComputeResourcesServiceOp) Edit(ctx context.Context, id int, editRequest *CloudbootComputeResourceEditRequest) (*Response, error) {
	if editRequest == nil || id < 1 {
		return nil, godo.NewArgError("editRequest || id", "cannot be nil or less than 1")
	}

	path := fmt.Sprintf("%s/%d%s", hypervisorsBasePath, id, apiFormat)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, editRequest)
	if err != nil {
		return nil, err
	}
	log.Println("CloudbootComputeResource [Edit]  req: ", req)

	return s.client.Do(ctx, req, nil)
}


// CloudbootAvailableResources - List all Cloudboot available resources
func (s *CloudbootComputeResourcesServiceOp) CloudbootAvailableResources(ctx context.Context) ([]Asset, *Response, error) {
	path := cloudBootAvailableResourcesBasePath + apiFormat
	path, err := addOptions(path, nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var out []map[string]Asset
	resp, err := s.client.Do(ctx, req, &out)
	if err != nil {
		return nil, resp, err
	}

	arr := make([]Asset, len(out))
	for i := range arr {
		arr[i] = out[i]["asset"]
	}

	return arr, resp, err
}
