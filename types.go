package helloworldchaos

import (
	"context"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/utils"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

type Impl struct {
	client.Client
	Log logr.Logger

	decoder *utils.ContainerRecordDecoder
}

// This corresponds to the Apply phase of HelloWorldChaos. The execution of HelloWorldChaos will be triggered.
func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("Hello world!")
	return v1alpha1.Injected, nil
}

// This corresponds to the Recover phase of HelloWorldChaos. The reconciler will be triggered to recover the chaos action.
func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("Goodbye world!")
	return v1alpha1.NotInjected, nil
}

// NewImpl returns a new HelloWorldChaos implementation instance.
func NewImpl(c client.Client, log logr.Logger, decoder *utils.ContainerRecordDecoder) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "helloworldchaos",
		Object: &v1alpha1.HelloWorldChaos{},
		Impl: &Impl{
			Client:  c,
			Log:     log.WithName("helloworldchaos"),
			decoder: decoder,
		},
		ObjectList: &v1alpha1.HelloWorldChaosList{},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
