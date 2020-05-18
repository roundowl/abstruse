package worker

import (
	"github.com/jkuri/abstruse/internal/worker/etcd"
	"github.com/jkuri/abstruse/internal/worker/grpc"
	"github.com/spf13/viper"
)

// Options is global config for worker app.
type Options struct {
	Etcd *etcd.Options
	GRPC *grpc.Options
}

// NewOptions returns worker app config.
func NewOptions(v *viper.Viper) (*Options, error) {
	opts := &Options{}
	err := v.Unmarshal(opts)
	return opts, err
}
