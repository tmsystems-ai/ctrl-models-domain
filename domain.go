package models

import (
	"time"

	"gorm.io/datatypes"
)

type Role struct {
	Id                    uint                   `json:"id" gorm:"primaryKey"`
	Name                  string                 `json:"name" gorm:"size:255"`
	PermissionAllocations []PermissionAllocation `json:"permission_allocations" gorm:"foreignKey:RoleId;references:Id"`
}

type PermissionAllocation struct {
	Id           uint       `json:"id" gorm:"primaryKey"`
	RoleId       uint       `json:"role_id" gorm:"index"`
	Role         Role       `json:"role"`
	PermissionId uint       `json:"permission_id" gorm:"index"`
	Permission   Permission `json:"permission"`
}

type Permission struct {
	Id                    uint                   `json:"id" gorm:"primaryKey"`
	Name                  string                 `json:"name" gorm:"size:255"`
	PermissionAllocations []PermissionAllocation `json:"permission_allocations" gorm:"foreignKey:PermissionId;references:Id"`
}

type PriorityLevel struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:255;unique"`
}

type PriorityLevelRule struct {
	Id              uint          `json:"id" gorm:"primaryKey"`
	AgencyName      string        `json:"agency_name" gorm:"size:255;index"`
	EmailDomain     string        `json:"email_domain" gorm:"size:255;index"`
	PriorityLevelId uint          `json:"priority_level_id"`
	PriorityLevel   PriorityLevel `json:"priority_level" gorm:"foreignKey:PriorityLevelId;references:Id"`
}

type Ticket struct {
	Id              uint          `json:"id" gorm:"primaryKey"`
	CreatedDate     time.Time     `json:"created_date"`
	ReceivedDate    time.Time     `json:"received_date"`
	EmailThreadId   uint          `json:"email_thread_id" gorm:"index"`
	EmailThread     EmailThread   `json:"email_thread"`
	CreatedById     *uint         `json:"created_by_id"`
	CreatedBy       *User         `json:"created_by" gorm:"foreignKey:CreatedById;references:Id"`
	TicketStatusId  uint          `json:"ticket_status_id"`
	TicketStatus    TicketStatus  `json:"ticket_status" gorm:"foreignKey:TicketStatusId;references:Id"`
	AssignedToId    *uint         `json:"assigned_to_id"`
	AssignedTo      User          `json:"assigned_to" gorm:"foreignKey:AssignedToId;references:Id"`
	IsDeleted       bool          `json:"is_deleted"`
	DeletedDate     *time.Time    `json:"deleted_date,omitempty"`
	DeletedById     *uint         `json:"deleted_by_id,omitempty"`
	DeletedBy       *User         `json:"deleted_by,omitempty" gorm:"foreignKey:DeletedById;references:Id"`
	IsArchived      bool          `json:"is_archived"`
	ArchivedDate    *time.Time    `json:"archived_date,omitempty"`
	PriorityLevelId uint          `json:"priority_level_id"`
	PriorityLevel   PriorityLevel `json:"priority_level" gorm:"foreignKey:PriorityLevelId;references:Id"`
	CompletedDate   *time.Time    `json:"completed_date,omitempty"`
	CustomerId      *uint         `json:"customer_id"`
	Customer        Customer      `json:"customer" gorm:"foreignKey:CustomerId;references:Id"`
}

type ArchivedTicket struct {
	ID                             uint64         `json:"id" gorm:"primaryKey"`
	OriginalID                     uint64         `json:"original_id"`
	CreatedDate                    time.Time      `json:"created_date"`
	ReceivedDate                   time.Time      `json:"received_date"`
	EmailThreadID                  uint64         `json:"email_thread_id"`
	EmailThreadSubject             string         `json:"email_thread_subject"`
	EmailThreadReceivedFrom        string         `json:"email_thread_received_from"`
	EmailThreadProviderThreadID    string         `json:"email_thread_provider_thread_id"`
	EmailThreadProviderThreadTopic string         `json:"email_thread_provider_thread_topic"`
	CreatedByID                    *uint64        `json:"created_by_id"`
	TicketStatusID                 uint64         `json:"ticket_status_id"`
	AssignedToID                   *uint64        `json:"assigned_to_id"`
	PriorityLevelID                uint64         `json:"priority_level_id"`
	CompletedDate                  *time.Time     `json:"completed_date"`
	CustomerID                     *uint64        `json:"customer_id"`
	ArchivedDate                   time.Time      `json:"archived_date"`
	Emails                         datatypes.JSON `json:"emails"`
	EmailAttachments               datatypes.JSON `json:"email_attachments"`
	SharedInboxGroupID             uint64         `json:"shared_inbox_group_id"`
	InboxGroupName                 string         `json:"inbox_group_name"`
	MailboxEmailAddress            string         `json:"mailbox_email_address"`
}

type ArchivedEmail struct {
	ID                    uint64                    `json:"id" gorm:"primaryKey"`
	SubjectLine           string                    `json:"subject_line"`
	Body                  string                    `json:"body"`
	To                    string                    `json:"to"`
	From                  string                    `json:"from"`
	URI                   string                    `json:"uri"`
	ReceivedDate          time.Time                 `json:"received_date"`
	MailProviderMessageID string                    `json:"mail_provider_message_id"`
	Attachments           []ArchivedEmailAttachment `json:"attachments" gorm:"foreignKey:EmailID"`
	GCSFolder             string                    `json:"gcs_folder"`
}

type ArchivedEmailAttachment struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	EmailID     uint64 `json:"email_id" gorm:"index"`
	URI         string `json:"uri"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Inline      bool   `json:"inline"`
}
type ArchivedEmailThread struct {
	ID                  uint      `gorm:"column:id;primaryKey;autoIncrement"`
	ReceivedFrom        string    `gorm:"column:received_from;type:varchar(255)"`
	SubjectLine         string    `gorm:"column:subject_line;type:varchar(255)"`
	SharedInboxGroupID  uint      `gorm:"column:shared_inbox_group_id"`
	ReceivedDate        time.Time `gorm:"column:received_date"`
	ProviderThreadID    string    `gorm:"column:provider_thread_id;type:varchar(550)"`
	ProviderThreadTopic string    `gorm:"column:provider_thread_topic;type:varchar(550)"`
	OriginalID          uint      `gorm:"column:original_id"`
	ArchivedDate        time.Time `gorm:"column:archived_date;default:CURRENT_TIMESTAMP(3)"`
}
type Email struct {
	Id                    uint              `json:"id" gorm:"primaryKey"`
	SubjectLine           string            `json:"subject_line" gorm:"type:text"`
	Body                  string            `json:"body" gorm:"type:mediumtext"`
	To                    string            `json:"to" gorm:"type:text"`
	From                  string            `json:"from" gorm:"type:text"`
	URI                   string            `json:"uri" gorm:"size:2048"`
	EmailAttachments      []EmailAttachment `json:"email_attachments" gorm:"foreignKey:EmailId;references:Id"`
	EmailThreadId         uint              `json:"email_thread_id" gorm:"index"`
	EmailThread           EmailThread       `json:"email_thread"`
	ReceivedDate          time.Time         `json:"received_date"`
	MailProviderMessageId string            `json:"mail_provider_message_id" gorm:"size:550;index"`
	SadHash               string            `json:"sad_hash" gorm:"size:64;index"`
	IsDeleted             bool              `json:"is_deleted" gorm:"default:false"`
	DeletedDate           *time.Time        `json:"deleted_date,omitempty" gorm:"default:null"`
	AttachmentsProcessed  bool              `json:"attachments_processed" gorm:"default:false"`
	AttachmentsProcessing bool              `json:"attachments_processing" gorm:"default:false"`
	GCSFolder             string            `json:"gcs_folder" gorm:"size:550"`
	EmailDownloadFailed   bool              `json:"email_download_failed" gorm:"default:false"`
	RetryCounter          int               `json:"retry_counter" gorm:"default:0"`
}

type EmailTemplate struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	Title         string    `json:"title" gorm:"size:255"`
	LastUpdatedOn time.Time `json:"last_updated_on"`
	Subject       string    `json:"subject" gorm:"size:255"`
	Body          string    `json:"body" gorm:"type:text"`
}

type EmailThread struct {
	Id                  uint             `json:"id" gorm:"primaryKey"`
	ReceivedFrom        string           `json:"received_from" gorm:"size:255"`
	Emails              []Email          `json:"emails" gorm:"foreignKey:EmailThreadId;references:Id"`
	SubjectLine         string           `json:"subject_line" gorm:"size:255"`
	SharedInboxGroupId  uint             `json:"shared_inbox_group_id"`
	SharedInboxGroup    SharedInboxGroup `json:"shared_inbox_group" gorm:"foreignKey:SharedInboxGroupId;references:Id"`
	ReceivedDate        time.Time        `json:"received_date"`
	ProviderThreadId    string           `json:"provider_thread_id" gorm:"size:550;index"`
	ProviderThreadTopic string           `json:"provider_thread_topic" gorm:"size:550"`
}

type EmailAttachment struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	URI         string `json:"uri" gorm:"size:4096"`
	EmailId     uint   `json:"email_id"`
	Filename    string `json:"filename" gorm:"size:4096"`
	ContentType string `json:"content_type" gorm:"size:255"`
	Inline      bool   `json:"inline" gorm:"default:false"`
}

type EmailTemplateField struct {
	Id             uint   `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"size:255"`
	TicketProperty string `json:"ticket_property" gorm:"size:255"`
}

type MailProvider struct {
	Id                  uint   `json:"id" gorm:"primaryKey"`
	Name                string `json:"name" gorm:"size:255"`
	GetURI              string `json:"get_uri" gorm:"size:2048"`
	SendURI             string `json:"send_uri" gorm:"size:2048"`
	SynchronizeURI      string `json:"synchronize_uri" gorm:"size:2048"`
	MoveURI             string `json:"move_uri" gorm:"size:2048"`
	GetAttachmentURI    string `json:"get_attachment_uri" gorm:"size:2048"`
	CreateAttachmentURI string `json:"create_attachment_uri" gorm:"size:2048"`
}

type MailProviderCredential struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	TenantId     string `json:"tenant_id" gorm:"size:255"`
	ClientId     string `json:"client_id" gorm:"size:255"`
	ClientSecret string `json:"client_secret" gorm:"size:255"`
	RedirectURI  string `json:"redirect_uri" gorm:"size:2048"`
	AccessToken  string `json:"access_token" gorm:"size:5200"`
	TokenExpiry  *int64 `json:"token_expiry"`
}

type SharedInboxGroup struct {
	Id                       uint                     `json:"id" gorm:"primaryKey"`
	InboxGroupName           string                   `json:"inbox_group_name" gorm:"size:255"`
	EmailAddress             string                   `json:"email_address" gorm:"size:255"`
	IsAutomatic              bool                     `json:"is_automatic"`
	IsEnabled                bool                     `json:"is_enabled"`
	TicketStatusAllocations  []TicketStatusAllocation `json:"ticket_status_allocations" gorm:"foreignKey:SharedInboxGroupId;references:Id"`
	MailProviderCredentialId uint                     `json:"mail_provider_credential_id"`
	MailProviderCredential   MailProviderCredential   `json:"mail_provider_credential" gorm:"foreignKey:MailProviderCredentialId;references:Id"`
	MailProviderId           uint                     `json:"mail_provider_id"`
	MailProvider             MailProvider             `json:"mail_provider" gorm:"foreignKey:MailProviderId;references:Id"`
	EmailThreads             []EmailThread            `json:"email_threads" gorm:"foreignKey:SharedInboxGroupId;references:Id"`
	LastSyncTime             *time.Time               `json:"last_sync_time" gorm:"default:null"`
	DisabledTime             *time.Time               `json:"disabled_time" gorm:"default:null"`
	IsPolling                bool                     `json:"is_polling" gorm:"default:false"`
	EmailSignature           string                   `json:"email_signature" gorm:"type:mediumtext"`
	IsInitial                bool                     `json:"is_initial" gorm:"default:false"`
	IsInitialActive          bool                     `json:"is_initial_active" gorm:"default:false"`
}

type SigUserAllocation struct {
	Id                 uint             `json:"id" gorm:"primaryKey"`
	SharedInboxGroupId uint             `json:"shared_inbox_group_id"`
	SharedInboxGroup   SharedInboxGroup `json:"shared_inbox_group" gorm:"foreignKey:SharedInboxGroupId;references:Id"`
	UserId             uint             `json:"user_id"`
	User               User             `json:"user" gorm:"foreignKey:UserId;references:Id"`
	IsEnabled          bool             `json:"is_enabled"`
}

type User struct {
	Id                   uint                `json:"id" gorm:"primaryKey"`
	FirstName            string              `json:"first_name" gorm:"size:255"`
	LastName             string              `json:"last_name" gorm:"size:255"`
	EmployeeId           uint                `json:"employee_id"`
	RoleId               uint                `json:"role_id"`
	Role                 Role                `json:"role" gorm:"foreignKey:RoleId;references:Id"`
	Password             string              `json:"password" gorm:"size:255"`
	Email                string              `json:"email" gorm:"size:255"`
	PhoneNumber          string              `json:"phone_number" gorm:"size:255"`
	Tickets              []Ticket            `json:"tickets" gorm:"foreignKey:AssignedToId;references:Id"`
	SigUserAllocations   []SigUserAllocation `json:"sig_user_allocations" gorm:"foreignKey:UserId;references:Id"`
	ProfilePictureURI    string              `json:"profile_picture_uri" gorm:"size:2048"`
	Country              string              `json:"country" gorm:"size:255"`
	RequirePasswordReset bool                `json:"require_password_reset"`
	UserPreferences      []UserPreference    `json:"user_preferences" gorm:"foreignKey:UserId;references:Id"`
}

type UserPreference struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	UserId *uint  `json:"user_id" gorm:"index"`
	Type   string `json:"type" gorm:"size:255;index"`
	Value  string `json:"value" gorm:"size:255"`
	Key    string `json:"key" gorm:"size:255;index"`
}

type TicketStatus struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"size:255"`
	IsEditable  bool   `json:"is_editable"`
	IsOrderable bool   `json:"is_orderable"`
}

type TicketStatusAllocation struct {
	Id                 uint             `json:"id" gorm:"primaryKey"`
	TicketStatusId     uint             `json:"ticket_status_id"`
	TicketStatus       TicketStatus     `json:"ticket_status" gorm:"foreignKey:TicketStatusId;references:Id"`
	SharedInboxGroupId uint             `json:"shared_inbox_group_id"`
	SharedInboxGroup   SharedInboxGroup `json:"shared_inbox_group" gorm:"foreignKey:SharedInboxGroupId;references:Id"`
	Order              int              `json:"order"`
	IsEnabled          bool             `json:"is_enabled"`
}

type Customer struct {
	Id          uint     `json:"id" gorm:"primaryKey"`
	AgencyName  string   `json:"agency_name" gorm:"size:255"`
	Email       string   `json:"email" gorm:"size:255"`
	EmailDomain string   `json:"email_domain" gorm:"size:255;index"`
	IsActive    bool     `json:"is_active"`
	Tickets     []Ticket `json:"tickets" gorm:"foreignKey:CustomerId;references:Id"`
	IsDeleted   bool     `json:"is_deleted"`
}

type Notification struct {
	Id               uint      `json:"id" gorm:"primaryKey"`
	UserId           uint      `json:"user_id"`
	User             User      `json:"user" gorm:"foreignKey:UserId;references:Id"`
	TicketId         uint      `json:"ticket_id"`
	Ticket           Ticket    `json:"ticket" gorm:"foreignKey:TicketId;references:Id"`
	NotificationTime time.Time `json:"notification_time"`
}

type SystemSetting struct {
	Key   string `gorm:"column:key;primaryKey"`
	Value string
}

func (SystemSetting) TableName() string {
	return "system_settings"
}
func (ArchivedEmail) TableName() string {
	return "archived_email"
}

func (ArchivedEmailThread) TableName() string {
	return "archived_email_thread"
}

type AuthToken struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `gorm:"precision:1" json:"expires_at"`
	CreatedAt time.Time `gorm:"precision:1" json:"created_at"`
}
