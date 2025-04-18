package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//nolint:tparallel // subtests use same mock controller
func TestCheck(t *testing.T) {
	t.Parallel()

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// inbox / outbox addr only matters for mocks, using devnet
	addrs, err := contracts.GetAddresses(context.Background(), netconf.Devnet)
	require.NoError(t, err)
	outbox := addrs.SolverNetOutbox

	for _, tt := range checkTestCases(t, solver) {
		t.Run(tt.name, func(t *testing.T) {
			backends, clients := testBackends(t)

			callAllower := func(_ uint64, _ common.Address, _ []byte) bool { return !tt.disallowCall }
			handler := handlerAdapter(newCheckHandler(newChecker(backends, callAllower, solver, outbox)))

			if tt.mock != nil {
				tt.mock(clients)

				destClient := clients.Client(t, tt.req.DestinationChainID)
				mockDidFill(t, destClient, outbox, false)
				mockFill(t, destClient, outbox, tt.res.RejectReason == types.RejectDestCallReverts.String())
				mockFillFee(t, destClient, outbox)
			}

			body, err := json.Marshal(tt.req)
			require.NoError(t, err)

			ctx := context.Background()
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointCheck, bytes.NewBuffer(body))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)

			respBody, err := io.ReadAll(rr.Body)
			require.NoError(t, err)

			var res types.CheckResponse
			err = json.Unmarshal(respBody, &res)
			require.NoError(t, err)

			t.Logf("resp_body: %s", respBody)

			require.Equal(t, tt.res.Rejected, res.Rejected)
			require.Equal(t, tt.res.RejectReason, res.RejectReason)
			require.Equal(t, tt.res.Accepted, res.Accepted)

			clients.Finish(t)
		})
	}
}
