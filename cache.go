package opslevel

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type Cacher struct {
	mutex        sync.Mutex
	Tiers        map[string]Tier
	Lifecycles   map[string]Lifecycle
	Systems      map[string]System
	Teams        map[string]Team
	Categories   map[string]Category
	Levels       map[string]Level
	Filters      map[string]Filter
	Integrations map[string]Integration
	Repositories map[string]Repository
	InfraSchemas map[string]InfrastructureResourceSchema
}

func (cacher *Cacher) TryGetTier(alias string) (*Tier, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Tiers[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetLifecycle(alias string) (*Lifecycle, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Lifecycles[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetSystems(alias string) (*System, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Systems[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetTeam(alias string) (*Team, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Teams[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetCategory(alias string) (*Category, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Categories[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetLevel(alias string) (*Level, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Levels[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetFilter(alias string) (*Filter, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Filters[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetIntegration(alias string) (*Integration, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Integrations[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetRepository(alias string) (*Repository, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.Repositories[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) TryGetInfrastructureSchema(alias string) (*InfrastructureResourceSchema, bool) {
	cacher.mutex.Lock()
	defer cacher.mutex.Unlock()
	if v, ok := cacher.InfraSchemas[alias]; ok {
		return &v, ok
	}
	return nil, false
}

func (cacher *Cacher) doCacheTiers(client *Client) {
	log.Debug().Msg("Caching 'Tier' lookup table from API ...")

	data, dataErr := client.ListTiers()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Tier' from API - REASON: %s", dataErr.Error())
	}
	for _, item := range data {
		cacher.Tiers[item.Alias] = item
	}
}

func (cacher *Cacher) doCacheLifecycles(client *Client) {
	log.Debug().Msg("Caching 'Lifecycle' lookup table from API ...")

	data, dataErr := client.ListLifecycles()
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Lifecycle' from API - REASON: %s", dataErr.Error())
	}
	for _, item := range data {
		cacher.Lifecycles[item.Alias] = item
	}
}

func (cacher *Cacher) doCacheSystems(client *Client) {
	log.Debug().Msg("Caching 'Systems' lookup table from API ...")

	data, dataErr := client.ListSystems(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Systems' from API - REASON: %s", dataErr.Error())
	}
	if data == nil {
		return
	}

	for _, item := range data.Nodes {
		for _, alias := range item.Aliases {
			cacher.Systems[alias] = item
		}
	}
}

func (cacher *Cacher) doCacheTeams(client *Client) {
	log.Debug().Msg("Caching 'Team' lookup table from API ...")

	data, dataErr := client.ListTeams(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Team' from API - REASON: %s", dataErr.Error())
	}
	if data == nil {
		return
	}

	for _, item := range data.Nodes {
		for _, alias := range item.Aliases {
			cacher.Teams[alias] = item
		}
	}
}

func (cacher *Cacher) doCacheCategories(client *Client) {
	log.Debug().Msg("Caching 'Category' lookup table from API ...")

	data, dataErr := client.ListCategories(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Category' from API - REASON: %s", dataErr.Error())
	}
	if data == nil {
		return
	}

	for _, item := range data.Nodes {
		cacher.Categories[item.Alias()] = item
	}
}

func (cacher *Cacher) doCacheLevels(client *Client) {
	log.Debug().Msg("Caching 'Level' lookup table from API ...")

	data, dataErr := client.ListLevels(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Level' from API - REASON: %s", dataErr.Error())
	}

	for _, item := range data.Nodes {
		cacher.Levels[item.Alias] = item
	}
}

func (cacher *Cacher) doCacheFilters(client *Client) {
	log.Debug().Msg("Caching 'Filter' lookup table from API ...")

	data, dataErr := client.ListFilters(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Filter' from API - REASON: %s", dataErr.Error())
	}
	if data == nil {
		return
	}

	for _, item := range data.Nodes {
		cacher.Filters[item.Alias()] = item
	}
}

func (cacher *Cacher) doCacheIntegrations(client *Client) {
	log.Debug().Msg("Caching 'Integration' lookup table from API ...")

	data, dataErr := client.ListIntegrations(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Integration' from API - REASON: %s", dataErr.Error())
	}
	if data == nil {
		return
	}

	for _, item := range data.Nodes {
		cacher.Integrations[item.Alias()] = item
	}
}

func (cacher *Cacher) doCacheRepositories(client *Client) {
	log.Debug().Msg("Caching 'Repository' lookup table from API ...")

	data, dataErr := client.ListRepositories(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'Repository' from API - REASON: %s", dataErr.Error())
	}
	if data == nil {
		return
	}

	for _, item := range data.Nodes {
		cacher.Repositories[item.DefaultAlias] = item
	}
}

func (cacher *Cacher) doCacheInfraSchemas(client *Client) {
	log.Debug().Msg("Caching 'InfrastructureSchema' lookup table from API ...")

	data, dataErr := client.ListInfrastructureSchemas(nil)
	if dataErr != nil {
		log.Warn().Msgf("===> Failed to list all 'InfrastructureSchema' from API - REASON: %s", dataErr.Error())
	}
	if data == nil {
		return
	}

	for _, item := range data.Nodes {
		cacher.InfraSchemas[item.Type] = item
	}
}

func (cacher *Cacher) CacheTiers(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheTiers(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheLifecycles(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheLifecycles(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheSystems(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheSystems(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheTeams(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheTeams(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheCategories(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheCategories(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheLevels(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheLevels(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheFilters(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheFilters(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheIntegrations(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheIntegrations(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheRepositories(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheRepositories(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheInfraSchemas(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheInfraSchemas(client)
	cacher.mutex.Unlock()
}

func (cacher *Cacher) CacheAll(client *Client) {
	cacher.mutex.Lock()
	cacher.doCacheTiers(client)
	cacher.doCacheLifecycles(client)
	cacher.doCacheSystems(client)
	cacher.doCacheTeams(client)
	cacher.doCacheCategories(client)
	cacher.doCacheLevels(client)
	cacher.doCacheFilters(client)
	cacher.doCacheIntegrations(client)
	cacher.doCacheRepositories(client)
	cacher.doCacheInfraSchemas(client)
	cacher.mutex.Unlock()
}

var Cache = &Cacher{
	mutex:        sync.Mutex{},
	Tiers:        make(map[string]Tier),
	Lifecycles:   make(map[string]Lifecycle),
	Systems:      make(map[string]System),
	Teams:        make(map[string]Team),
	Categories:   make(map[string]Category),
	Levels:       make(map[string]Level),
	Filters:      make(map[string]Filter),
	Integrations: make(map[string]Integration),
	Repositories: make(map[string]Repository),
	InfraSchemas: make(map[string]InfrastructureResourceSchema),
}
