package dial

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/log"
)

type ActiveL2RollupProvider struct {
	rollupEndpoints []string
	checkDuration   time.Duration
	networkTimeout  time.Duration
	log             log.Logger

	activeTimeout time.Time

	currentRollupClient *sources.RollupClient
	clientLock          *sync.Mutex
}

func NewActiveL2RollupProvider(
	rollupUrls []string,
	checkDuration time.Duration,
	networkTimeout time.Duration,
	logger log.Logger,
) (*ActiveL2RollupProvider, error) {
	if len(rollupUrls) == 0 {
		return nil, errors.New("empty rollup urls list")
	}

	return &ActiveL2RollupProvider{
		rollupEndpoints: rollupUrls,
		checkDuration:   checkDuration,
		networkTimeout:  networkTimeout,
		log:             logger,
	}, nil
}

func (p *ActiveL2RollupProvider) RollupClient(ctx context.Context) (*sources.RollupClient, error) {
	err := p.ensureActiveEndpoint(ctx)
	if err != nil {
		return nil, err
	}
	p.clientLock.Lock()
	defer p.clientLock.Unlock()
	return p.currentRollupClient, nil
}

func (p *ActiveL2RollupProvider) ensureActiveEndpoint(ctx context.Context) error {
	if !p.shouldCheck() {
		return nil
	}

	if err := p.findActiveEndpoints(ctx); err != nil {
		return err
	}
	p.activeTimeout = time.Now().Add(p.checkDuration)
	return nil
}

func (p *ActiveL2RollupProvider) shouldCheck() bool {
	return time.Now().After(p.activeTimeout)
}

func (p *ActiveL2RollupProvider) findActiveEndpoints(ctx context.Context) error {
	// If current is not active, dial new sequencers until finding an active one.
	ts := time.Now()
	for i := 0; ; i++ {
		active, err := p.checkCurrentSequencer(ctx)
		if err != nil {
			if ctx.Err() != nil {
				p.log.Warn("Error querying active sequencer, trying next.", "err", err, "try", i)
				return fmt.Errorf("querying active sequencer: %w", err)
			}
			p.log.Warn("Error querying active sequencer, trying next.", "err", err, "try", i)
		} else if active {
			p.log.Debug("Current sequencer active.", "try", i)
			return nil
		} else {
			p.log.Info("Current sequencer inactive, trying next.", "try", i)
		}

		// After iterating over all endpoints, sleep if all were just inactive,
		// to avoid spamming the sequencers in a loop.
		if (i+1)%p.NumEndpoints() == 0 {
			d := time.Until(ts.Add(p.checkDuration))
			time.Sleep(d) // accepts negative
			ts = time.Now()
		}

		if err := p.dialNextSequencer(ctx, i); err != nil {
			return fmt.Errorf("dialing next sequencer: %w", err)
		}
	}
}

func (p *ActiveL2RollupProvider) checkCurrentSequencer(ctx context.Context) (bool, error) {
	cctx, cancel := context.WithTimeout(ctx, p.networkTimeout)
	defer cancel()
	p.clientLock.Lock()
	defer p.clientLock.Unlock()
	return p.currentRollupClient.SequencerActive(cctx)
}

func (p *ActiveL2RollupProvider) dialNextSequencer(ctx context.Context, idx int) error {
	cctx, cancel := context.WithTimeout(ctx, p.networkTimeout)
	defer cancel()
	ep := p.rollupEndpoints[idx]

	rollupClient, err := DialRollupClientWithTimeout(cctx, p.networkTimeout, p.log, ep)
	if err != nil {
		return fmt.Errorf("dialing rollup client: %w", err)
	}
	p.clientLock.Lock()
	defer p.clientLock.Unlock()
	p.currentRollupClient = rollupClient
	return nil
}

func (p *ActiveL2RollupProvider) NumEndpoints() int {
	return len(p.rollupEndpoints)
}

func (p *ActiveL2RollupProvider) Close() {
	if p.currentRollupClient != nil {
		p.currentRollupClient.Close()
	}
}
