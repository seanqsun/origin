package etcd

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"

	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	userapi "github.com/openshift/origin/pkg/user/apis/user/"
	"github.com/openshift/origin/pkg/user/registry/identitymetadata"
	"github.com/openshift/origin/pkg/util/restoptions"
)

type REST struct {
	*registry.Store
}

func NewREST(optsGetter restoptions.Getter) (*REST, error) {
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &userapi.IdentityMetadata{} },
		NewListFunc:              func() runtime.Object { return &userapi.IdentityMetadataList{} },
		DefaultQualifiedResource: userapi.Resource("identitymetadatas"),

		TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)},

		CreateStrategy: identitymetadata.Strategy,
		UpdateStrategy: identitymetadata.Strategy,
		DeleteStrategy: identitymetadata.Strategy,
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}

	return &REST{store}, nil
}
