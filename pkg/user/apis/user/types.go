package user

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

// Auth system gets identity name and provider
// POST to UserIdentityMapping, get back error or a filled out UserIdentityMapping object

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type User struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	FullName string

	Identities []string

	Groups []string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []User
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Identity struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	// ProviderName is the source of identity information
	ProviderName string

	// ProviderUserName uniquely represents this identity in the scope of the provider
	ProviderUserName string

	// User is a reference to the user this identity is associated with
	// Both Name and UID must be set
	User core.ObjectReference

	Extra map[string]string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IdentityList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Identity
}

// +genclient
// +genclient:nonNamespaced
// +genclient:onlyVerbs=get,create,update,delete
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserIdentityMapping struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Identity core.ObjectReference
	User     core.ObjectReference
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Group represents a referenceable set of Users
type Group struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Users []string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GroupList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Group
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TODO: Docs
type IdentityMetadata struct {
	metav1.TypeMeta
	// Standard object's metadata.
	metav1.ObjectMeta

	// ExpiresIn is the seconds from CreationTime before this authentication expires
	// Wired to TTLFunc in ETCD
	ExpiresIn int64

	// UserName is the user name associated with this authentication
	UserName string

	// UserUID is the unique UID associated with this authentication
	UserUID string

	// IdentityProviderName is the name of the IDP associated with this authentication
	// Do we want this?
	IdentityProviderName string

	// IdentityProviderUserName uniquely identifies this particular user for this provider
	// Do we want this?
	IdentityProviderUserName string

	// IdentityProviderGroups is the names of the groups the user is a member of for the given IDP
	IdentityProviderGroups []string

	// IdentityProviderExtra contains any additional information that the IDP thought was interesting
	// Difference between this and normal extra which contains scopes?
	IdentityProviderExtra map[string]OptionalNames
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityMetadataList is a collection of IdentityMetadata
// TODO: Review protobuf best practices
type IdentityMetadataList struct {
	metav1.TypeMeta
	// Standard object's metadata.
	metav1.ListMeta
	// Items is the list of IdentityMetadata
	Items []IdentityMetadata
}
