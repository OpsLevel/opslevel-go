package opslevel

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type AliasCacher struct {
	mutex        sync.Mutex
	Tiers        map[string]Tier
	Lifecycles   map[string]Lifecycle
	Teams        map[string]Team
	Categories   map[string]Category
	Levels       map[string]Level
	Filters      map[string]Filter
	Integrations map[string]Integration
	Repositories map[string]Repository
}

func (c *AliasCacher) TryGetTier(alias string) (*Tier, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Tiers[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) TryGetLifecycle(alias string) (*Lifecycle, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Lifecycles[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) TryGetTeam(alias string) (*Team, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Teams[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) TryGetCategory(alias string) (*Category, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Categories[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) TryGetLevel(alias string) (*Level, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Levels[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) TryGetFilter(alias string) (*Filter, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Filters[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) TryGetIntegration(alias string) (*Integration, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Integrations[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) TryGetRepository(alias string) (*Repository, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if v, ok := c.Repositories[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (c *AliasCacher) doCacheTiers(client *Client) {
	log.Info().Msg("Caching 'Tier' lookup table from API ...")

	data, dataErr := client.ListTiers()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Tier' from API - REASON: %s", dataErr.Error())
	}
	for _, item := range data {
		c.Tiers[string(item.Alias)] = item
	}
}

func (c *AliasCacher) doCacheLifecycles(client *Client) {
	log.Info().Msg("Caching 'Lifecycle' lookup table from API ...")

	data, dataErr := client.ListLifecycles()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Lifecycle' from API - REASON: %s", dataErr.Error())
	}
	for _, item := range data {
		c.Lifecycles[string(item.Alias)] = item
	}
}

func (c *AliasCacher) doCacheTeams(client *Client) {
	log.Info().Msg("Caching 'Team' lookup table from API ...")

	data, dataErr := client.ListTeams()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Team' from API - REASON: %s", dataErr.Error())
	}

	for _, item := range data {
		c.Teams[string(item.Alias)] = item
	}
}

func (c *AliasCacher) doCacheCategories(client *Client) {
	log.Info().Msg("Caching 'Category' lookup table from API ...")

	data, dataErr := client.ListCategories()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Category' from API - REASON: %s", dataErr.Error())
	}

	for _, item := range data {
		c.Categories[item.Alias()] = item
	}
}

func (c *AliasCacher) doCacheLevels(client *Client) {
	log.Info().Msg("Caching 'Level' lookup table from API ...")

	data, dataErr := client.ListLevels()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Level' from API - REASON: %s", dataErr.Error())
	}

	for _, item := range data {
		c.Levels[string(item.Alias)] = item
	}
}

func (c *AliasCacher) doCacheFilters(client *Client) {
	log.Info().Msg("Caching 'Filter' lookup table from API ...")

	data, dataErr := client.ListFilters()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Filter' from API - REASON: %s", dataErr.Error())
	}

	for _, item := range data {
		c.Filters[item.Alias()] = item
	}
}

func (c *AliasCacher) doCacheIntegrations(client *Client) {
	log.Info().Msg("Caching 'Integration' lookup table from API ...")

	data, dataErr := client.ListIntegrations()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Integration' from API - REASON: %s", dataErr.Error())
	}

	for _, item := range data {
		c.Integrations[item.Alias()] = item
	}
}

func (c *AliasCacher) doCacheRepositories(client *Client) {
	log.Info().Msg("Caching 'Repository' lookup table from API ...")

	data, dataErr := client.ListRepositories()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Repository' from API - REASON: %s", dataErr.Error())
	}

	for _, item := range data {
		c.Repositories[item.DefaultAlias] = item
	}
}

func (c *AliasCacher) CacheTiers(client *Client) {
	c.mutex.Lock()
	c.doCacheTiers(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheLifecycles(client *Client) {
	c.mutex.Lock()
	c.doCacheLifecycles(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheTeams(client *Client) {
	c.mutex.Lock()
	c.doCacheTeams(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheCategories(client *Client) {
	c.mutex.Lock()
	c.doCacheCategories(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheLevels(client *Client) {
	c.mutex.Lock()
	c.doCacheLevels(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheFilters(client *Client) {
	c.mutex.Lock()
	c.doCacheFilters(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheIntegrations(client *Client) {
	c.mutex.Lock()
	c.doCacheIntegrations(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheRepositories(client *Client) {
	c.mutex.Lock()
	c.doCacheRepositories(client)
	c.mutex.Unlock()
}

func (c *AliasCacher) CacheAll(client *Client) {
	c.mutex.Lock()
	c.doCacheTiers(client)
	c.doCacheLifecycles(client)
	c.doCacheTeams(client)
	c.doCacheCategories(client)
	c.doCacheLevels(client)
	c.doCacheFilters(client)
	c.doCacheIntegrations(client)
	c.doCacheRepositories(client)
	c.mutex.Unlock()
}

var AliasCache = &AliasCacher{
	mutex:        sync.Mutex{},
	Tiers:        make(map[string]Tier),
	Lifecycles:   make(map[string]Lifecycle),
	Teams:        make(map[string]Team),
	Categories:   make(map[string]Category),
	Levels:       make(map[string]Level),
	Filters:      make(map[string]Filter),
	Integrations: make(map[string]Integration),
	Repositories: make(map[string]Repository),
}
