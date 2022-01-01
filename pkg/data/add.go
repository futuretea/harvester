package data

import (
	"context"

	"github.com/harvester/harvester/pkg/config"
)

func Init(ctx context.Context, mgmtCtx *config.Management, options config.Options) error {
	__traceStack()

	if err := createCRDs(ctx, mgmtCtx.RestConfig); err != nil {
		return err
	}

	if err := addPublicNamespace(mgmtCtx.Apply); err != nil {
		return err
	}
	if err := addAPIService(mgmtCtx.Apply, options.Namespace); err != nil {
		return err
	}
	if err := addAuthenticatedRoles(mgmtCtx.Apply); err != nil {
		return err
	}

	return createTemplates(mgmtCtx, publicNamespace)
}
