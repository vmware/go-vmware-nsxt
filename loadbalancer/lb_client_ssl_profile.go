/*
 * NSX API
 *
 * VMware NSX REST API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type LbClientSslProfile struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	// Schema for this resource
	Schema string `json:"_schema,omitempty"`

	// Link to this resource
	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`

	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`

	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`

	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`

	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`

	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity.
	Protection string `json:"_protection,omitempty"`

	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// Unique identifier of this resource
	Id string `json:"id,omitempty"`

	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// supported SSL cipher list to client side
	Ciphers []string `json:"ciphers,omitempty"`

	// This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure.
	IsSecure bool `json:"is_secure,omitempty"`

	// During SSL handshake as part of the SSL client Hello client sends an ordered list of ciphers that it can support (or prefers) and typically server selects the first one from the top of that list it can also support. For Perfect Forward Secrecy(PFS), server could override the client's preference.
	PreferServerCiphers bool `json:"prefer_server_ciphers,omitempty"`

	// SSL versions TLS1.1 and TLS1.2 are supported and enabled by default. SSLv2, SSLv3, and TLS1.0 are supported, but disabled by default.
	Protocols []string `json:"protocols,omitempty"`

	// SSL session caching allows SSL client and server to reuse previously negotiated security parameters avoiding the expensive public key operation during handshake.
	SessionCacheEnabled bool `json:"session_cache_enabled,omitempty"`

	// Session cache timeout specifies how long the SSL session parameters are held on to and can be reused.
	SessionCacheTimeout int64 `json:"session_cache_timeout,omitempty"`
}
