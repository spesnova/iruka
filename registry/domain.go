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
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Domain{}, err
	}

	if opts.Hostname == "" {
		return schema.Domain{}, errors.New("hostname parameter is required, but missing")
	}

	id := uuid.NewUUID()
	currentTime := time.Now()
	domain := schema.Domain{
		ID:        id,
		AppID:     app.ID,
		Hostname:  opts.Hostname,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	j, err := marshal(domain)

	if err != nil {
		return schema.Domain{}, err
	}

	key := path.Join(r.keyPrefix, domainPrefix, domain.ID.String())
	_, err = r.etcd.Create(key, string(j), 0)

	if err != nil {
		return schema.Domain{}, err
	}

	return domain, nil
}

func (r *Registry) DestroyDomain(identity string) (schema.Domain, error) {
	domain, err := r.Domain(identity)

	if err != nil {
		return schema.Domain{}, err
	}

	key := path.Join(r.keyPrefix, domainPrefix, domain.ID.String())
	_, err = r.etcd.Delete(key, true)

	if err != nil {
		return schema.Domain{}, errors.New("Failed to delete domain: " + domain.ID.String())
	}

	return domain, nil
}

func (r *Registry) Domain(identity string) (schema.Domain, error) {
	var domain schema.Domain

	domains, err := r.Domains()

	if err != nil {
		return domain, err
	}

	if uuid.Parse(identity) == nil {
		for _, domain := range domains {
			if domain.Hostname == identity {
				return domain, nil
			}
		}
	} else {
		for _, domain := range domains {
			if uuid.Equal(domain.ID, uuid.Parse(identity)) {
				return domain, nil
			}
		}
	}

	return domain, errors.New("No such domain: " + identity)
}

func (r *Registry) DomainFilteredByApp(appIdentity, identity string) (schema.Domain, error) {
	domain, err := r.Domain(identity)

	if err != nil {
		return schema.Domain{}, err
	}

	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Domain{}, err
	}

	if uuid.Equal(domain.AppID, app.ID) {
		return domain, nil
	}

	return domain, errors.New("No such domain: " + identity)
}

func (r *Registry) Domains() ([]schema.Domain, error) {
	key := path.Join(r.keyPrefix, domainPrefix)
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

func (r *Registry) DomainsFilteredByApp(appIdentity string) ([]schema.Domain, error) {
	var domains []schema.Domain

	app, err := r.App(appIdentity)

	if err != nil {
		return nil, err
	}

	ds, err := r.Domains()

	if err != nil {
		return nil, err
	}

	if ds == nil {
		return nil, nil
	}

	for _, d := range ds {
		if uuid.Equal(d.AppID, app.ID) {
			domains = append(domains, d)
		}
	}

	return domains, nil
}
