package onappgo

import (
  "fmt"
  "time"
)

// Infoboxes - 
type Infoboxes struct {
  DisplayInfoboxes bool     `json:"display_infoboxes,bool"`
  HiddenInfoboxes  []string `json:"hidden_infoboxes,omitempty"`
}

// Permission - 
type Permission struct {
  CreatedAt  time.Time `json:"created_at,omitempty"`
  ID         int       `json:"id,omitempty"`
  Identifier string    `json:"identifier,omitempty"`
  UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

// Permissions - 
type Permissions struct {
  Permission Permission `json:"permission,omitempty"`
}

// Role - 
type Role struct {
  CreatedAt   time.Time     `json:"created_at,omitempty"`
  ID          int           `json:"id,omitempty"`
  Identifier  string        `json:"identifier,omitempty"`
  Label       string        `json:"label,omitempty"`
  System      bool          `json:"system,bool"`
  UpdatedAt   time.Time     `json:"updated_at,omitempty"`
  UsersCount  int           `json:"users_count,omitempty"`
  Permissions []Permissions `json:"permissions,omitempty"`
}

// Roles - 
type Roles struct {
  Role Role `json:"role,omitempty"`
}

// User - 
type User struct {
  ActivatedAt             time.Time          `json:"activated_at,omitempty"`
  Avatar                  interface{}        `json:"avatar,omitempty"`
  BillingPlanID           int                `json:"billing_plan_id,omitempty"`
  CdnAccountStatus        string             `json:"cdn_account_status,omitempty"`
  CdnStatus               string             `json:"cdn_status,omitempty"`
  CreatedAt               time.Time          `json:"created_at,omitempty"`
  DeletedAt               time.Time          `json:"deleted_at,omitempty"`
  Email                   string             `json:"email,omitempty"`
  FirewallID              int                `json:"firewall_id,omitempty"`
  FirstName               string             `json:"first_name,omitempty"`
  GroupID                 int                `json:"group_id,omitempty"`
  ID                      int                `json:"id,omitempty"`
  Identifier              string             `json:"identifier,omitempty"`
  ImageTemplateGroupID    int                `json:"image_template_group_id,omitempty"`
  Infoboxes               Infoboxes          `json:"infoboxes,omitempty"`
  LastName                string             `json:"last_name,omitempty"`
  Locale                  string             `json:"locale,omitempty"`
  Login                   string             `json:"login,omitempty"`
  PasswordChangedAt       time.Time          `json:"password_changed_at,omitempty"`
  RegisteredYubikey       bool               `json:"registered_yubikey,bool"`
  Status                  string             `json:"status,omitempty"`
  Supplied                bool               `json:"supplied,bool"`
  SuspendAt               time.Time          `json:"suspend_at,omitempty"`
  SystemTheme             string             `json:"system_theme,omitempty"`
  TimeZone                string             `json:"time_zone,omitempty"`
  UpdatedAt               time.Time          `json:"updated_at,omitempty"`
  UseGravatar             bool               `json:"use_gravatar,bool"`
  UserGroupID             int                `json:"user_group_id,omitempty"`
  BucketID                int                `json:"bucket_id,omitempty"`
  UsedCpus                int                `json:"used_cpus,omitempty"`
  UsedMemory              int                `json:"used_memory,omitempty"`
  UsedCPUShares           int                `json:"used_cpu_shares,omitempty"`
  UsedDiskSize            int                `json:"used_disk_size,omitempty"`
  MemoryAvailable         int                `json:"memory_available,omitempty"`
  DiskSpaceAvailable      int                `json:"disk_space_available,omitempty"`
  Roles                   []Roles            `json:"roles,omitempty"`
  MonthlyPrice            int                `json:"monthly_price,omitempty"`
  PaymentAmount           int                `json:"payment_amount,omitempty"`
  OutstandingAmount       int                `json:"outstanding_amount,omitempty"`
  TotalAmount             int                `json:"total_amount,omitempty"`
  DiscountDueToFree       int                `json:"discount_due_to_free,omitempty"`
  TotalAmountWithDiscount int                `json:"total_amount_with_discount,omitempty"`
  AdditionalFields        []AdditionalFields `json:"additional_fields,omitempty"`
  UsedIPAddresses         []IPAddress        `json:"used_ip_addresses,omitempty"`
}

// UserCreateRequest - 
type UserCreateRequest struct {
  Login            string             `json:"login,omitempty"`
  Email            string             `json:"email,omitempty"`
  FirstName        string             `json:"first_name,omitempty"`
  LastName         string             `json:"last_name,omitempty"`
  Password         string             `json:"password,omitempty"`
  UserGroupID      string             `json:"user_group_id,omitempty"`
  BillingPlanID    string             `json:"billing_plan_id,omitempty"`
  RoleIds          []string           `json:"role_ids,omitempty"`
  AdditionalFields []AdditionalFields `json:"additional_fields,omitempty"`
}

// Debug - print formatted User structure
func (obj User) Debug() {
  fmt.Printf("           ID: %d\n", obj.ID)
  fmt.Printf("    FirstName: %s\n", obj.FirstName)
  fmt.Printf("     LastName: %s\n", obj.LastName)
  fmt.Printf("        Email: %s\n", obj.Email)
  fmt.Printf("        Login: %s\n", obj.Login)
  fmt.Printf("   Identifier: %s\n", obj.Identifier)
  fmt.Printf("    CreatedAt: %s\n", obj.CreatedAt)
  fmt.Printf("     UsedCpus: %d\n", obj.UsedCpus)
  fmt.Printf("   UsedMemory: %d\n", obj.UsedMemory)
  fmt.Printf("UsedCPUShares: %d\n", obj.UsedCPUShares)
  fmt.Printf(" UsedDiskSize: %d\n", obj.UsedDiskSize)
}
