package identitymetadata

import (
	"context"
	"crypto/sha256"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	userapi "github.com/openshift/origin/pkg/user/apis/user"
	"github.com/openshift/origin/pkg/user/apis/user/validation"
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

//
var Strategy = strategy{legacyscheme.Scheme, names.SimpleNameGenerator}

func (strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {

}

func (strategy) NamespaceScoped() bool {
	return false
}

// PrepareForCreate sorts and dedupes the groups
// Name of the object is the hash of all the fields
func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	im := obj.(*userapi.IdentityMetadata)
	im.IdentityProviderGroups = strings.New(im.IdentityProviderGroups...).List()
	if len(im.Name) == 0 && len(im.GeneratedName) == 0 {
		copied := *im
		copied.TypeMeta = metav1.TypeMeta{}
		copied.ObjectMeta = metav1.ObjectMeta{}
		copied.ExpiresIn = 0

		hasher := sha256.New()
		hashutil.DeepHashObect(hasher, copied)

		im.Name := hex.EncodeToString(hasher.sum32())
	}

}

func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return validation.ValidateIdentityMetadata(obj.(*userapi.IdentityMetadata))
}

func (strategy) AllowCreateOnUpdate() bool {
	return false
}

func (strategy) AllowUnconditionalUpdate() bool {
	return false
}

func (strategy) Canonicalize(obj runtime.Object) {}

func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return validation.ValidateIdentitytMetadataUpdate(obj.(*userapi.IdentityMetadata), old.(*userapi.IdentityMetadata))
}
