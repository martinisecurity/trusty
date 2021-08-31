package pgsql

import (
	"context"

	"github.com/ekspand/trusty/internal/db"
	"github.com/ekspand/trusty/internal/db/orgsdb/model"
	"github.com/juju/errors"
)

// UpdateOrg inserts or updates Organization
func (p *Provider) UpdateOrg(ctx context.Context, org *model.Organization) (*model.Organization, error) {
	id, err := p.NextID()
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = db.Validate(org)
	if err != nil {
		return nil, errors.Trace(err)
	}

	logger.Debugf("extern_id=%s, login=%s", org.ExternalID, org.Login)

	res := new(model.Organization)

	err = p.db.QueryRowContext(ctx, `
			INSERT INTO orgs(id,extern_id,reg_id,provider,login,name,email,billing_email,company,location,avatar_url,html_url,type,created_at,updated_at,street_address,city,postal_code,region,country,phone,approver_email,approver_name,status,expires_at)
				VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25)
			ON CONFLICT (provider,login)
			DO UPDATE
				SET name=$6,email=$7,billing_email=$8,company=$9,location=$10,avatar_url=$11,html_url=$12,type=$13,created_at=$14,updated_at=$15,
				street_address=$16,city=$17,postal_code=$18,region=$19,country=$20,phone=$21
			RETURNING id,extern_id,reg_id,provider,login,name,email,billing_email,company,location,avatar_url,html_url,type,created_at,updated_at,street_address,city,postal_code,region,country,phone,approver_email,approver_name,status,expires_at
			;`, id,
		org.ExternalID, org.RegistrationID,
		org.Provider, org.Login, org.Name, org.Email, org.BillingEmail, org.Company, org.Location,
		org.AvatarURL, org.URL,
		org.Type,
		org.CreatedAt.UTC(),
		org.UpdatedAt.UTC(),
		org.Street, org.City, org.PostalCode, org.Region, org.Country, org.Phone,
		org.ApproverEmail, org.ApproverName, org.Status, org.ExpiresAt.UTC(),
	).Scan(&res.ID,
		&res.ExternalID,
		&res.RegistrationID,
		&res.Provider,
		&res.Login,
		&res.Name,
		&res.Email,
		&res.BillingEmail,
		&res.Company,
		&res.Location,
		&res.AvatarURL,
		&res.URL,
		&res.Type,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Street,
		&res.City,
		&res.PostalCode,
		&res.Region,
		&res.Country,
		&res.Phone,
		&res.ApproverEmail,
		&res.ApproverName,
		&res.Status,
		&res.ExpiresAt,
	)
	if err != nil {
		return nil, errors.Trace(err)
	}
	res.CreatedAt = res.CreatedAt.UTC()
	res.UpdatedAt = res.UpdatedAt.UTC()
	res.ExpiresAt = res.ExpiresAt.UTC()
	return res, nil
}

// UpdateOrgStatus updates the status and approver
func (p *Provider) UpdateOrgStatus(ctx context.Context, org *model.Organization) (*model.Organization, error) {
	err := db.Validate(org)
	if err != nil {
		return nil, errors.Trace(err)
	}

	logger.Debugf("id=%d, extern_id=%s, status=%s", org.ID, org.ExternalID, org.Status)

	res := new(model.Organization)

	err = p.db.QueryRowContext(ctx, `
			UPDATE orgs
				SET approver_email=$2,approver_name=$3,status=$4,expires_at=$5
			WHERE id = $1
			RETURNING id,extern_id,reg_id,provider,login,name,email,billing_email,company,location,avatar_url,html_url,type,created_at,updated_at,street_address,city,postal_code,region,country,phone,approver_email,approver_name,status,expires_at
			;`, org.ID, org.ApproverEmail, org.ApproverName, org.Status, org.ExpiresAt.UTC(),
	).Scan(&res.ID,
		&res.ExternalID,
		&res.RegistrationID,
		&res.Provider,
		&res.Login,
		&res.Name,
		&res.Email,
		&res.BillingEmail,
		&res.Company,
		&res.Location,
		&res.AvatarURL,
		&res.URL,
		&res.Type,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Street,
		&res.City,
		&res.PostalCode,
		&res.Region,
		&res.Country,
		&res.Phone,
		&res.ApproverEmail,
		&res.ApproverName,
		&res.Status,
		&res.ExpiresAt,
	)
	if err != nil {
		return nil, errors.Trace(err)
	}
	res.CreatedAt = res.CreatedAt.UTC()
	res.UpdatedAt = res.UpdatedAt.UTC()
	res.ExpiresAt = res.ExpiresAt.UTC()
	return res, nil
}

// GetOrg returns Organization
func (p *Provider) GetOrg(ctx context.Context, id uint64) (*model.Organization, error) {
	res := new(model.Organization)

	err := p.db.QueryRowContext(ctx,
		`SELECT id,extern_id,reg_id,provider,login,name,email,billing_email,company,location,avatar_url,html_url,type,created_at,updated_at,street_address,city,postal_code,region,country,phone,approver_email,approver_name,status,expires_at
		FROM orgs
		WHERE id=$1
		;`, id,
	).Scan(&res.ID,
		&res.ExternalID,
		&res.RegistrationID,
		&res.Provider,
		&res.Login,
		&res.Name,
		&res.Email,
		&res.BillingEmail,
		&res.Company,
		&res.Location,
		&res.AvatarURL,
		&res.URL,
		&res.Type,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Street,
		&res.City,
		&res.PostalCode,
		&res.Region,
		&res.Country,
		&res.Phone,
		&res.ApproverEmail,
		&res.ApproverName,
		&res.Status,
		&res.ExpiresAt,
	)
	if err != nil {
		return nil, errors.Trace(err)
	}
	res.CreatedAt = res.CreatedAt.UTC()
	res.UpdatedAt = res.UpdatedAt.UTC()
	res.ExpiresAt = res.ExpiresAt.UTC()
	return res, nil
}

// GetOrgByExternalID returns Organization by external ID
func (p *Provider) GetOrgByExternalID(ctx context.Context, provider, externalID string) (*model.Organization, error) {
	res := new(model.Organization)

	err := p.db.QueryRowContext(ctx,
		`SELECT id,extern_id,reg_id,provider,login,name,email,billing_email,company,location,avatar_url,html_url,type,created_at,updated_at,street_address,city,postal_code,region,country,phone,approver_email,approver_name,status,expires_at
		FROM orgs
		WHERE provider=$1 AND extern_id=$2
		;`, provider, externalID,
	).Scan(&res.ID,
		&res.ExternalID,
		&res.RegistrationID,
		&res.Provider,
		&res.Login,
		&res.Name,
		&res.Email,
		&res.BillingEmail,
		&res.Company,
		&res.Location,
		&res.AvatarURL,
		&res.URL,
		&res.Type,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Street,
		&res.City,
		&res.PostalCode,
		&res.Region,
		&res.Country,
		&res.Phone,
		&res.ApproverEmail,
		&res.ApproverName,
		&res.Status,
		&res.ExpiresAt,
	)
	if err != nil {
		return nil, errors.Trace(err)
	}
	res.CreatedAt = res.CreatedAt.UTC()
	res.UpdatedAt = res.UpdatedAt.UTC()
	res.ExpiresAt = res.ExpiresAt.UTC()
	return res, nil
}

// RemoveOrg deletes org and all its members
func (p *Provider) RemoveOrg(ctx context.Context, id uint64) error {

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Trace(err)
	}

	_, err = p.db.ExecContext(ctx, `DELETE FROM orgmembers WHERE org_id=$1;`, id)
	if err != nil {
		logger.Errorf("err=[%s]", errors.Details(err))
		return errors.Trace(err)
	}
	_, err = p.db.ExecContext(ctx, `DELETE FROM orgs WHERE id=$1;`, id)
	if err != nil {
		logger.Errorf("err=[%s]", errors.Details(err))
		return errors.Trace(err)
	}
	_, err = p.db.ExecContext(ctx, `DELETE FROM subscriptions WHERE id=$1;`, id)
	if err != nil {
		logger.Errorf("err=[%s]", errors.Details(err))
		return errors.Trace(err)
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		return errors.Trace(err)
	}

	logger.Noticef("id=%d", id)

	return nil
}

// UpdateRepo inserts or updates Repository
func (p *Provider) UpdateRepo(ctx context.Context, repo *model.Repository) (*model.Repository, error) {
	id, err := p.NextID()
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = db.Validate(repo)
	if err != nil {
		return nil, errors.Trace(err)
	}

	res := new(model.Repository)

	err = p.db.QueryRowContext(ctx, `
			INSERT INTO repos(id,org_id,extern_id,provider,name,email,company,avatar_url,type,created_at,updated_at)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			ON CONFLICT (org_id, name)
			DO UPDATE
				SET avatar_url=$8, type=$9, updated_at=$11
			RETURNING id,org_id,extern_id,provider,name,email,company,avatar_url,type,created_at,updated_at
			;`, id, repo.OrgID, repo.ExternalID, repo.Provider, repo.Name, repo.Email, repo.Company, repo.AvatarURL, repo.Type,
		repo.CreatedAt, repo.UpdatedAt,
	).Scan(&res.ID,
		&res.OrgID,
		&res.ExternalID,
		&res.Provider,
		&res.Name,
		&res.Email,
		&res.Company,
		&res.AvatarURL,
		&res.Type,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return res, nil
}

// GetRepo returns Repository
func (p *Provider) GetRepo(ctx context.Context, id uint64) (*model.Repository, error) {
	res := new(model.Repository)

	err := p.db.QueryRowContext(ctx,
		`SELECT id,org_id,extern_id,provider,name,email,company,avatar_url,type,created_at,updated_at
		FROM repos
		WHERE id=$1
		;`, id,
	).Scan(&res.ID,
		&res.OrgID,
		&res.ExternalID,
		&res.Provider,
		&res.Name,
		&res.Email,
		&res.Company,
		&res.AvatarURL,
		&res.Type,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return res, nil
}

// AddOrgMember adds a user to Org
func (p *Provider) AddOrgMember(ctx context.Context, orgID, userID uint64, role, membershipSource string) (*model.OrgMembership, error) {
	id, err := p.NextID()
	if err != nil {
		return nil, errors.Trace(err)
	}

	if role == "" || len(role) > db.MaxLenForName {
		return nil, errors.Errorf("invalid role: %q", role)
	}

	member := new(model.OrgMembership)

	err = p.db.QueryRowContext(ctx, `
		INSERT INTO orgmembers(id,org_id,user_id,role,source)
			VALUES($1, $2, $3, $4, $5)
		ON CONFLICT ON CONSTRAINT membership
			DO UPDATE SET role=$4,source=$5
		RETURNING id,org_id,user_id,role,source
		;`, id, orgID, userID, role, membershipSource).
		Scan(
			&member.ID,
			&member.OrgID,
			&member.UserID,
			&member.Role,
			&member.Source,
		)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return member, nil
}

// GetOrgMembers returns list of membership info
func (p *Provider) GetOrgMembers(ctx context.Context, orgID uint64) ([]*model.OrgMemberInfo, error) {

	res, err := p.db.QueryContext(ctx, `
		SELECT
			orgmembers.id,
			orgmembers.org_id,
			orgs.name,
			orgmembers.user_id,
			orgmembers.role,
			orgmembers.source,
			users.name,
			users.email
		FROM
			orgmembers
		LEFT JOIN users ON users.ID = orgmembers.user_id
		LEFT JOIN orgs ON orgs.ID = orgmembers.org_id
		WHERE
		      orgmembers.org_id = $1
		;
		`, orgID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer res.Close()

	list := make([]*model.OrgMemberInfo, 0, 100)

	for res.Next() {
		m := new(model.OrgMemberInfo)
		err = res.Scan(
			&m.MembershipID,
			&m.OrgID,
			&m.OrgName,
			&m.UserID,
			&m.Role,
			&m.Source,
			&m.Name,
			&m.Email,
		)
		if err != nil {
			return nil, errors.Trace(err)
		}

		list = append(list, m)
	}

	return list, nil
}

// RemoveOrgMembers removes users from the org
func (p *Provider) RemoveOrgMembers(ctx context.Context, orgID uint64, all bool) ([]*model.OrgMembership, error) {
	var sql string
	if all {
		sql = `DELETE FROM orgmembers
				WHERE org_id=$1
				RETURNING id,org_id,user_id,role,source;`
	} else {
		sql = `DELETE FROM members
				WHERE org_id=$1 AND source NOT IN ('github')
				RETURNING id,org_id,user_id,role,source;`
	}

	deleted := make([]*model.OrgMembership, 0, 100)
	res, err := p.db.QueryContext(ctx, sql, orgID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer res.Close()

	for res.Next() {
		m := new(model.OrgMembership)
		err = res.Scan(
			&m.ID,
			&m.OrgID,
			&m.UserID,
			&m.Role,
			&m.Source,
		)
		if err != nil {
			return nil, errors.Trace(err)
		}

		deleted = append(deleted, m)
	}

	return deleted, nil
}

// RemoveOrgMember remove users from the org
func (p *Provider) RemoveOrgMember(ctx context.Context, orgID, memberID uint64) (*model.OrgMembership, error) {
	m := new(model.OrgMembership)

	err := p.db.QueryRowContext(ctx,
		`DELETE FROM orgmembers
			WHERE org_id=$1 AND user_id=$2
			RETURNING id,org_id,user_id,role,source;`,
		orgID, memberID).
		Scan(
			&m.ID,
			&m.OrgID,
			&m.UserID,
			&m.Role,
			&m.Source,
		)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return m, nil
}

// GetUserMemberships returns list of membership info
func (p *Provider) GetUserMemberships(ctx context.Context, userID uint64) ([]*model.OrgMemberInfo, error) {
	res, err := p.db.QueryContext(ctx, `
		SELECT
			orgmembers.id,
			orgmembers.org_id,
			orgs.name,
			orgmembers.user_id,
			orgmembers.role,
			orgmembers.source,
			users.name,
			users.email
		FROM
			orgmembers
		LEFT JOIN users ON users.ID = orgmembers.user_id
		LEFT JOIN orgs ON orgs.ID = orgmembers.org_id
		WHERE orgmembers.user_id = $1
		;
		`, userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer res.Close()

	list := make([]*model.OrgMemberInfo, 0, 100)

	for res.Next() {
		m := new(model.OrgMemberInfo)
		err = res.Scan(
			&m.MembershipID,
			&m.OrgID,
			&m.OrgName,
			&m.UserID,
			&m.Role,
			&m.Source,
			&m.Name,
			&m.Email,
		)
		if err != nil {
			return nil, errors.Trace(err)
		}

		list = append(list, m)
	}

	return list, nil
}

// GetUserOrgs returns list of orgs
func (p *Provider) GetUserOrgs(ctx context.Context, userID uint64) ([]*model.Organization, error) {
	q, err := p.db.QueryContext(ctx, `
			SELECT
				orgs.id,
				orgs.extern_id,
				orgs.reg_id,
				orgs.provider,
				orgs.login,
				orgs.name,
				orgs.email,
				orgs.billing_email,
				orgs.company,
				orgs.location,
				orgs.avatar_url,
				orgs.html_url,
				orgs.type,
				orgs.created_at,
				orgs.updated_at,
				orgs.street_address,
				orgs.city,
				orgs.postal_code,
				orgs.region,
				orgs.country,
				orgs.phone,
				orgs.approver_email,
				orgs.approver_name,
				orgs.status,
				orgs.expires_at
			FROM
				orgmembers
			LEFT JOIN users ON users.ID = orgmembers.user_id
			LEFT JOIN orgs ON orgs.ID = orgmembers.org_id
			WHERE orgmembers.user_id = $1
			;
			`, userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer q.Close()

	list := make([]*model.Organization, 0, 100)

	for q.Next() {
		res := new(model.Organization)
		err = q.Scan(
			&res.ID,
			&res.ExternalID,
			&res.RegistrationID,
			&res.Provider,
			&res.Login,
			&res.Name,
			&res.Email,
			&res.BillingEmail,
			&res.Company,
			&res.Location,
			&res.AvatarURL,
			&res.URL,
			&res.Type,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.Street,
			&res.City,
			&res.PostalCode,
			&res.Region,
			&res.Country,
			&res.Phone,
			&res.ApproverEmail,
			&res.ApproverName,
			&res.Status,
			&res.ExpiresAt,
		)
		if err != nil {
			return nil, errors.Trace(err)
		}
		res.CreatedAt = res.CreatedAt.UTC()
		res.UpdatedAt = res.UpdatedAt.UTC()
		res.ExpiresAt = res.ExpiresAt.UTC()
		list = append(list, res)
	}

	return list, nil
}
