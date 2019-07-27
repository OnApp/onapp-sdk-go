package onappgo

import (
  "context"
  "net/http"
  "fmt"

  "github.com/digitalocean/godo"
)

const instancePackagesBasePath = "instance_packages"

// InstancePackagesService is an interface for interfacing with the Instance Packages
// endpoints of the OnApp API
// https://docs.onapp.com/apim/latest/instance-packages
type InstancePackagesService interface {
  List(context.Context, *ListOptions) ([]InstancePackage, *Response, error)
  Get(context.Context, int) (*InstancePackage, *Response, error)
  Create(context.Context, *InstancePackageCreateRequest) (*InstancePackage, *Response, error)
  // Delete(context.Context, int) (*Response, error)
  Delete(context.Context, int, interface{}) (*Transaction, *Response, error)
  // Edit(context.Context, int, *ListOptions) ([]InstancePackage, *Response, error)
}

// InstancePackagesServiceOp handles communication with the Instance Package related methods of the
// OnApp API.
type InstancePackagesServiceOp struct {
  client *Client
}

var _ InstancePackagesService = &InstancePackagesServiceOp{}

// InstancePackage represents a InstancePackage
type InstancePackage struct {
  Bandwidth      int         `json:"bandwidth,omitempty"`
  BillingPlanIds []int       `json:"billing_plan_ids,omitempty"`
  BucketsIds     []int       `json:"buckets_ids,omitempty"`
  Cpus           int         `json:"cpus,omitempty"`
  CreatedAt      string      `json:"created_at,omitempty"`
  DiskSize       int         `json:"disk_size,omitempty"`
  ID             int         `json:"id,omitempty"`
  Label          string      `json:"label,omitempty"`
  Memory         int         `json:"memory,omitempty"`
  OpenstackID    int         `json:"openstack_id,omitempty"`
  UpdatedAt      string      `json:"updated_at,omitempty"`
}

// InstancePackageCreateRequest represents a request to create a InstancePackage
type InstancePackageCreateRequest struct {
  Label     string `json:"label,omitempty"`
  Cpus      string `json:"cpus,omitempty"`
  Memory    string `json:"memory,omitempty"`
  DiskSize  string `json:"disk_size,omitempty"`
  Bandwidth string `json:"bandwidth,omitempty"`
}

type instancePackageCreateRequestRoot struct {
  InstancePackageCreateRequest  *InstancePackageCreateRequest  `json:"instance_package"`
}

type instancePackageRoot struct {
  InstancePackage  *InstancePackage  `json:"instance_package"`
}

func (d InstancePackageCreateRequest) String() string {
  return godo.Stringify(d)
}

// List all DataStoreGroups.
func (s *InstancePackagesServiceOp) List(ctx context.Context, opt *ListOptions) ([]InstancePackage, *Response, error) {
  path := instancePackagesBasePath + apiFormat
  path, err := addOptions(path, opt)
  if err != nil {
    return nil, nil, err
  }

  req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
  if err != nil {
    return nil, nil, err
  }

  var out []map[string]InstancePackage
  resp, err := s.client.Do(ctx, req, &out)

  if err != nil {
    return nil, resp, err
  }

  arr := make([]InstancePackage, len(out))
  for i := range arr {
    arr[i] = out[i]["instance_package"]
  }

  return arr, resp, err
}

// Get individual InstancePackage.
func (s *InstancePackagesServiceOp) Get(ctx context.Context, id int) (*InstancePackage, *Response, error) {
  if id < 1 {
    return nil, nil, godo.NewArgError("id", "cannot be less than 1")
  }

  path := fmt.Sprintf("%s/%d%s", instancePackagesBasePath, id, apiFormat)
  req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
  if err != nil {
    return nil, nil, err
  }

  root := new(instancePackageRoot)
  resp, err := s.client.Do(ctx, req, root)
  if err != nil {
    return nil, resp, err
  }

  return root.InstancePackage, resp, err
}

// Create InstancePackage.
func (s *InstancePackagesServiceOp) Create(ctx context.Context, createRequest *InstancePackageCreateRequest) (*InstancePackage, *Response, error) {
  if createRequest == nil {
    return nil, nil, godo.NewArgError("InstancePackage createRequest", "cannot be nil")
  }

  path := instancePackagesBasePath + apiFormat
  rootRequest := &instancePackageCreateRequestRoot{
    InstancePackageCreateRequest: createRequest,
  }

  req, err := s.client.NewRequest(ctx, http.MethodPost, path, rootRequest)
  if err != nil {
    return nil, nil, err
  }

  fmt.Println("\nInstancePackage [Create] req: ", req)

  root := new(instancePackageRoot)
  resp, err := s.client.Do(ctx, req, root)
  if err != nil {
    return nil, nil, err
  }

  return root.InstancePackage, resp, err
}

// Delete InstancePackage.
func (s *InstancePackagesServiceOp) Delete(ctx context.Context, id int, meta interface{}) (*Transaction, *Response, error) {
  if id < 1 {
    return nil, nil, godo.NewArgError("id", "cannot be less than 1")
  }

  path := fmt.Sprintf("%s/%d%s", instancePackagesBasePath, id, apiFormat)
  path, err := addOptions(path, meta)
  if err != nil {
    return nil, nil, err
  }

  req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
  if err != nil {
    return nil, nil, err
  }

  resp, err := s.client.Do(ctx, req, nil)
  if err != nil {
    return nil, resp, err
  }

  filter := struct{
    ParentID    int
    ParentType  string
  }{
    ParentID    : id,
    ParentType  : "InstancePackage",
  }

  return lastTransaction(ctx, s.client, filter)
  // return lastTransaction(ctx, s.client, id, "InstancePackage")
}

// Debug - print formatted InstancePackage structure
func (obj InstancePackage) Debug() {
  fmt.Printf("             ID: %d\n", obj.ID)
  fmt.Printf("          Label: %s\n", obj.Label)
  fmt.Printf("      Bandwidth: %d\n", obj.Bandwidth)
  fmt.Printf("           Cpus: %d\n", obj.Cpus)
  fmt.Printf("       DiskSize: %d\n", obj.DiskSize)
  fmt.Printf("         Memory: %d\n", obj.Memory)
}