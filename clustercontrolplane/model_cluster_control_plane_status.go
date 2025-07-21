package clustercontrolplane

// ClusterStatus represents the status of a cluster control plane
type ClusterStatus struct {
	// Controller status information
	ControllerStatus ControllerStatus `json:"controller_status,omitempty"`

	// Management plane adapter status
	MpAdapterStatus AdapterStatus `json:"mp_adapter_status,omitempty"`

	// CCP adapter status
	CcpAdapterStatus AdapterStatus `json:"ccp_adapter_status,omitempty"`

	// Agent information
	AgentInfo AgentInfo `json:"agent_info,omitempty"`

	// Overall status
	Status string `json:"status,omitempty"`

	// Resource type
	ResourceType string `json:"resource_type,omitempty"`

	// ID of the cluster
	Id string `json:"id,omitempty"`

	// Display name
	DisplayName string `json:"display_name,omitempty"`

	// Path
	Path string `json:"path,omitempty"`

	// Relative path
	RelativePath string `json:"relative_path,omitempty"`

	// Parent path
	ParentPath string `json:"parent_path,omitempty"`

	// Remote path
	RemotePath string `json:"remote_path,omitempty"`

	// Unique ID
	UniqueId string `json:"unique_id,omitempty"`

	// Realization ID
	RealizationId string `json:"realization_id,omitempty"`

	// Owner ID
	OwnerId string `json:"owner_id,omitempty"`

	// Marked for delete
	MarkedForDelete bool `json:"marked_for_delete,omitempty"`

	// Overridden
	Overridden bool `json:"overridden,omitempty"`

	// Create time
	CreateTime int64 `json:"_create_time,omitempty"`

	// Create user
	CreateUser string `json:"_create_user,omitempty"`

	// Last modified time
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`

	// Last modified user
	LastModifiedUser string `json:"_last_modified_user,omitempty"`

	// System owned
	SystemOwned bool `json:"_system_owned,omitempty"`

	// Protection
	Protection string `json:"_protection,omitempty"`

	// Revision
	Revision int64 `json:"_revision,omitempty"`
}

// ControllerStatus represents the controller status
type ControllerStatus struct {
	Status            string      `json:"status,omitempty"`
	ConnectedAgentNum int32       `json:"connected_agent_num,omitempty"`
	Conditions        []Condition `json:"conditions,omitempty"`
}

// AdapterStatus represents adapter status
type AdapterStatus struct {
	Status     string      `json:"status,omitempty"`
	Conditions []Condition `json:"conditions,omitempty"`
}

// AgentInfo represents agent information
type AgentInfo struct {
	HealthyAgentNum  int32 `json:"healthy_agent_num,omitempty"`
	DegradedAgentNum int32 `json:"degraded_agent_num,omitempty"`
	FailedAgentNum   int32 `json:"failed_agent_num,omitempty"`
}

// Condition represents a status condition
type Condition struct {
	ConditionType     string `json:"condition_type,omitempty"`
	LastHeartbeatTime int64  `json:"last_heartbeat_time,omitempty"`
	Status            string `json:"status,omitempty"`
	Reason            string `json:"reason,omitempty"`
	Message           string `json:"message,omitempty"`
}
