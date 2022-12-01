package registry

import (
	"fmt"
	pb "github.com/pydio/cells/v4/common/proto/registry"
	"golang.org/x/exp/maps"
	"time"
)

var _ Registry = (*TransientWrapper)(nil)

type TransientWrapper struct {
	Registry

	opts Options
}

func NewTransientWrapper(reg Registry, opts ...Option) Registry {

	// mandatory
	o := Options{}
	for _, opt := range opts {
		opt(&o)
	}

	return &TransientWrapper{
		Registry: reg,
		opts:     o,
	}
}

func (m *TransientWrapper) Register(item Item, opts ...RegisterOption) error {
	items := m.opts.filterItems(item)

	if len(items) == 1 {
		targetItems := m.Registry.ListAdjacentItems(item, WithType(pb.ItemType_SERVER), WithType(pb.ItemType_NODE))

		status := item.Metadata()[MetaStatusKey]

		for _, targetItem := range targetItems {
			if ms, ok := targetItem.(MetaSetter); ok {
				targetMeta := targetItem.Metadata()
				if (status == string(StatusStarting) || status == string(StatusStopping)) && targetMeta[MetaStatusKey] != string(StatusTransient) {
					targetMetaClone := maps.Clone(targetMeta)
					targetMetaClone[MetaTimestampKey] = fmt.Sprintf("%d", time.Now().UnixNano())
					targetMetaClone[MetaStatusKey] = string(StatusTransient)
					ms.SetMetadata(targetMetaClone)

					if err := m.Registry.Register(targetItem, opts...); err != nil {
						return err
					}
				}
			}
		}
	}

	return m.Registry.Register(item, opts...)
}
