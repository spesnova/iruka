package registry

import (
	"errors"
	"path"
	"sort"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/spesnova/iruka/schema"
)

const (
	domainPrefix = "domains"
)

type Domains []schema.Domain

// imprement the sort interface
func (d Domains) Len() int {
	return len(d)
}

func (d Domains) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Domains) Less(i, j int) bool {
	return d[i].ID.String() < d[j].ID.String()
}

func (r *Registry) CreateDomain(appIdentity string, opts schema.DomainCreateOpts) (schema.Domain, error) {
	if opts.Hostname == "" {
		return schema.Domain{}, errors.New("hostname parameter is required, but missing")
	}

	id := uuid.NewUUID()
	currentTime := time.Now()
	domain := schema.Domain{
		ID:        id,
		Hostname:  opts.Hostname,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	j, err := marshal(domain)

	if err != nil {
		return schema.Domain{}, err
	}

	key := path.Join(r.keyPrefix, domainPrefix, appIdentity, domain.ID.String())
	_, err = r.etcd.Create(key, string(j), 0)

	if err != nil {
		return schema.Domain{}, err
	}

	return domain, nil
}

func (r *Registry) DestroyDomain(appIdentity, domainID string) (schema.Domain, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Domain{}, err
	}

	domain, err := r.Domain(appIdentity, domainID)

	if err != nil {
		return schema.Domain{}, err
	}

	key := path.Join(r.keyPrefix, domainPrefix, app.ID.String(), domainID)
	_, err = r.etcd.Delete(key, true)

	if err != nil {
		if isKeyNotFound(err) {
			err = nil
		}
		return schema.Domain{}, err
	}

	return domain, nil
}

func (r *Registry) Domain(appIdentity, domainID string) (schema.Domain, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Domain{}, err
	}

	key := path.Join(r.keyPrefix, domainPrefix, app.ID.String(), domainID)
	res, err := r.etcd.Get(key, false, true)

	if err != nil {
		if isKeyNotFound(err) {
			err = nil
		}

		return schema.Domain{}, err
	}

	var domain schema.Domain
	err = unmarshal(res.Node.Value, &domain)

	if err != nil {
		return schema.Domain{}, err
	}

	return domain, nil
}

func (r *Registry) Domains(appIdentity string) ([]schema.Domain, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return nil, err
	}

	key := path.Join(r.keyPrefix, domainPrefix, app.ID.String())
	res, err := r.etcd.Get(key, false, true)

	if err != nil {
		if isKeyNotFound(err) {
			err = nil
		}
		return nil, err
	}

	if len(res.Node.Nodes) == 0 {
		return nil, nil
	}

	var domains Domains

	for _, node := range res.Node.Nodes {
		var domain schema.Domain
		err = unmarshal(node.Value, &domain)

		if err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}

	sort.Sort(sort.Reverse(domains))

	return domains, nil
}

func (r *Registry) UpdateDomain(appIdentity string, domainID string, opts schema.DomainUpdateOpts) (schema.Domain, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Domain{}, err
	}

	domain, err := r.Domain(appIdentity, domainID)

	if err != nil {
		return schema.Domain{}, err
	}

	if opts.ID.String() == "" {
		return schema.Domain{}, errors.New("id parameter is required, but missing")
	}

	if opts.Hostname != "" {
		domain.Hostname = opts.Hostname
	}

	j, err := marshal(domain)

	if err != nil {
		return schema.Domain{}, err
	}

	key := path.Join(r.keyPrefix, domainPrefix, app.ID.String(), domainID)

	if _, err := r.etcd.Set(key, string(j), 0); err != nil {
		return schema.Domain{}, err
	}

	return domain, nil
}
