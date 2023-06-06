package opslevel

import "github.com/rs/zerolog/log"

func (c *Client) SyncOpsLevelYml(service string, repository string) error {
	log.Info().Msgf("Syncing %s %s", service, repository)
	return nil
}
