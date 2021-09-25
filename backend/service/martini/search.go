package martini

import (
	"net/http"

	"github.com/go-phorce/dolly/rest"
	"github.com/go-phorce/dolly/xhttp/marshal"
	"github.com/go-phorce/dolly/xlog"
	"github.com/juju/errors"
	"github.com/martinisecurity/trusty/api"
	v1 "github.com/martinisecurity/trusty/api/v1"
	"github.com/martinisecurity/trusty/backend/db/orgsdb/model"
	"github.com/rvflash/opencorporates"
)

// SearchCorpsHandler syncs and returns user's orgs
func (s *Service) SearchCorpsHandler() rest.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ rest.Params) {
		name := api.GetQueryString(r.URL, "name")
		jurisdiction := api.GetQueryString(r.URL, "jurisdiction")

		res := v1.SearchOpenCorporatesResponse{
			Companies: make([]v1.OpenCorporatesCompany, 0, 100),
		}

		if name != "" {
			api := opencorporates.API()
			it := api.Companies(name, jurisdiction)
			for {
				company, err := it.Next()
				if err != nil {
					if err != opencorporates.EOF {
						logger.KV(xlog.ERROR, "err", errors.Details(err))
					}
					break
				}

				if !company.DissolutionDate.Time.IsZero() {
					continue
				}

				res.Companies = append(res.Companies, v1.OpenCorporatesCompany{
					Name:         company.Name,
					Kind:         company.Kind,
					Number:       company.Number,
					CountryCode:  company.CountryCode,
					Jurisdiction: company.Jurisdiction,
					CreationDate: company.CreationDate.Time,
					Street:       company.Address.Street,
					City:         company.Address.City,
					Region:       company.Address.Region,
					PostalCode:   company.Address.PostalCode,
					Country:      company.Address.Country,
				})
			}
		}
		marshal.WriteJSON(w, r, res)
	}
}

// SearchOrgsHandler returns Orgs
func (s *Service) SearchOrgsHandler() rest.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ rest.Params) {
		regID := api.GetQueryString(r.URL, "reg_id")
		frn := api.GetQueryString(r.URL, "frn")
		ctx := r.Context()
		res := v1.OrgsResponse{}

		orgs := map[uint64]*model.Organization{}

		if frn != "" {
			org, err := s.db.GetOrgByExternalID(ctx, v1.ProviderMartini, frn)
			if err == nil {
				orgs[org.ID] = org
			}
		}
		if regID != "" {
			org, err := s.db.GetOrgByRegistrationID(ctx, v1.ProviderMartini, regID)
			if err == nil {
				orgs[org.ID] = org
			}
		}

		for _, o := range orgs {
			res.Orgs = append(res.Orgs, *o.ToDto())
		}

		marshal.WriteJSON(w, r, res)
	}
}
