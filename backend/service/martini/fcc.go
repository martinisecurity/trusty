package martini

import (
	"context"
	"net/http"

	"github.com/ekspand/trusty/api"
	v1 "github.com/ekspand/trusty/api/v1"
	"github.com/ekspand/trusty/internal/db"
	"github.com/ekspand/trusty/pkg/fcc"
	"github.com/go-phorce/dolly/rest"
	"github.com/go-phorce/dolly/xhttp/httperror"
	"github.com/go-phorce/dolly/xhttp/marshal"
	"github.com/go-phorce/dolly/xlog"
	"github.com/juju/errors"
)

// FccFrnHandler handles v1.PathForMartiniFccFrn
func (s *Service) FccFrnHandler() rest.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ rest.Params) {
		filerID := api.GetQueryString(r.URL, "filer_id")
		if filerID == "" {
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("missing filer_id parameter"))
			return
		}

		res, err := s.getFrnResponse(r.Context(), filerID)
		if err != nil {
			marshal.WriteJSON(w, r, httperror.WithUnexpected("request failed: "+err.Error()).WithCause(err))
			return
		}

		logger.KV(xlog.DEBUG, "filerID", filerID, "response", res)
		marshal.WriteJSON(w, r, res)
	}
}

// FccContactHandler handles v1.PathForMartiniFccContact
func (s *Service) FccContactHandler() rest.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ rest.Params) {
		frn := api.GetQueryString(r.URL, "frn")
		if frn == "" {
			marshal.WriteJSON(w, r, httperror.WithInvalidRequest("missing frn parameter"))
			return
		}

		res, err := s.getFccContact(r.Context(), frn)
		if err != nil {
			marshal.WriteJSON(w, r, httperror.WithUnexpected("request failed: "+err.Error()).WithCause(err))
			return
		}

		logger.KV(xlog.DEBUG, "frn", frn, "response", res)
		marshal.WriteJSON(w, r, res)
	}
}

func (s *Service) getFrnResponse(ctx context.Context, filerID string) (*v1.FccFrnResponse, error) {
	if filerID == "123456" {
		res := new(v1.FccFrnResponse)
		err := marshal.DecodeBytes([]byte(testFRN), res)
		if err != nil {
			return nil, errors.Annotate(err, "failed to decode testFRN")
		}
		return res, nil
	}

	id, err := db.ID(filerID)
	if err != nil {
		return nil, errors.Annotate(err, "invalid filer ID")
	}

	cached, err := s.db.GetFRNResponse(ctx, id)
	if err == nil {
		res := new(v1.FccFrnResponse)
		err = marshal.DecodeBytes([]byte(cached.Response), res)
		if err == nil {
			return res, nil
		}
	}

	fccClient := fcc.NewAPIClient(s.FccBaseURL)
	fQueryResults, err := fccClient.GetFiler499Results(filerID)
	if err != nil {
		return nil, errors.Annotate(err, "unable to query FCC")
	}
	res := &v1.FccFrnResponse{
		Filers: filer499ResultsToDTO(fQueryResults),
	}

	js, _ := marshal.EncodeBytes(marshal.DontPrettyPrint, res)
	_, err = s.db.UpdateFRNResponse(ctx, id, string(js))
	if err != nil {
		logger.Errorf("filerID=%q, err=[%s]", filerID, errors.Details(err))
	}

	return res, nil
}

func (s *Service) getFccContact(ctx context.Context, frn string) (*v1.FccContactResponse, error) {
	if frn == "99999999" {
		res := new(v1.FccContactResponse)
		marshal.DecodeBytes([]byte(testContact), res)
		return res, nil
	}

	cached, err := s.db.GetFccContactResponse(ctx, frn)
	if err == nil {
		res := new(v1.FccContactResponse)
		err = marshal.DecodeBytes([]byte(cached.Response), res)
		if err == nil {
			return res, nil
		}
	}

	fccClient := fcc.NewAPIClient(s.FccBaseURL)
	cQueryResults, err := fccClient.GetContactResults(frn)
	if err != nil {
		return nil, errors.Annotate(err, "unable to query FCC")
	}

	res := contactQueryResultsToDTO(cQueryResults)

	js, _ := marshal.EncodeBytes(marshal.DontPrettyPrint, res)
	_, err = s.db.UpdateFccContactResponse(ctx, frn, string(js))
	if err != nil {
		logger.Errorf("frn=%q, err=[%s]", frn, errors.Details(err))
	}
	return res, nil
}

// filer499ResultsToDTO converts to v1.FccFrnResponse
func filer499ResultsToDTO(fq *fcc.Filer499Results) []v1.Filer {
	filers := []v1.Filer{}

	for _, f := range fq.Filers {
		fDTO := v1.Filer{
			FilerID: f.Form499ID,
			FilerIDInfo: v1.FilerIDInfo{
				LegalName: f.FilerIDInfo.LegalName,
				FRN:       f.FilerIDInfo.FRN,
				HQAddress: v1.HQAdress{
					AddressLine: f.FilerIDInfo.HQAddress.AddressLine,
					City:        f.FilerIDInfo.HQAddress.City,
					State:       f.FilerIDInfo.HQAddress.State,
					ZipCode:     f.FilerIDInfo.HQAddress.ZipCode,
				},
			},
		}

		filers = append(filers, fDTO)
	}

	return filers
}

// contactQueryResultsToDTO converts to v1.FccContactResults
func contactQueryResultsToDTO(c *fcc.ContactResults) *v1.FccContactResponse {
	return &v1.FccContactResponse{
		FRN:                 c.FRN,
		RegistrationDate:    c.RegistrationDate,
		LastUpdated:         c.LastUpdated,
		BusinessName:        c.BusinessName,
		BusinessType:        c.BusinessType,
		ContactOrganization: c.ContactOrganization,
		ContactPosition:     c.ContactPosition,
		ContactName:         c.ContactName,
		ContactAddress:      c.ContactAddress,
		ContactEmail:        c.ContactEmail,
		ContactPhone:        c.ContactPhone,
		ContactFax:          c.ContactFax,
	}
}

const testFRN = `{
	"filers": [
			{
					"filer_id_info": {
							"customer_inquiries_telephone": "2057453970",
							"frn": "99999999",
							"hq_address": {
									"address_line": "241 APPLEGATE TRACE",
									"city": "PELHAM",
									"state": "AL",
									"zip_code": "35124"
							},
							"legal_name": "LOW LATENCY COMMUNICATIONS LLC"
					},
					"filer_id": "123456"
			}
	]
}`

const testContact = `{
	"business_name": "Low Latency Communications, LLC",
	"business_type": "Private Sector, Limited Liability Corporation",
	"contact_address": "241 Applegate Trace, Pelham, AL 35124-2945, United States",
	"contact_email": "denis+test@ekspand.com",
	"contact_fax": "",
	"contact_name": "Mr Matthew D Hardeman",
	"contact_organization": "Low Latency Communications, LLC",
	"contact_phone": "",
	"contact_position": "Secretary",
	"frn": "99999999",
	"last_updated": "",
	"registration_date": "09/29/2015 09:58:00 AM"
}`