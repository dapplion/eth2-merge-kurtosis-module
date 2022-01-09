package el

import (
	"github.com/kurtosis-tech/eth2-merge-kurtosis-module/kurtosis-module/impl/participant_network/log_levels"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/services"
)

type ELClientLauncher interface {
	Launch(
		enclaveCtx *enclaves.EnclaveContext,
		serviceId services.ServiceID,
		loglevel log_levels.ParticipantLogLevel,
		networkId string,
		// If nil, then the node will be launched as a bootnode
		bootnodeContext *ELClientContext,
	) (
		resultClientCtx *ELClientContext,
		resultErr error,
	)
}
