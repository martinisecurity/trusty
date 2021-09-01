BEGIN;

create or replace function create_constraint_if_not_exists (
    s_name text, t_name text, c_name text, constraint_sql text
)
returns void AS
$$
begin
    -- Look for our constraint
    if not exists (select constraint_name
                   from information_schema.constraint_column_usage
                   where table_schema = s_name and table_name = t_name and constraint_name = c_name) then
        execute constraint_sql;
    end if;
end;
$$ language 'plpgsql';

--
-- USERS
--
CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL,
    extern_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    provider character varying(16) COLLATE pg_catalog."default" NOT NULL,
    login character varying(64) COLLATE pg_catalog."default" NOT NULL,
    name character varying(64) COLLATE pg_catalog."default" NOT NULL,
    email character varying(160) COLLATE pg_catalog."default" NOT NULL,
    company character varying(64) COLLATE pg_catalog."default" NOT NULL,
    avatar_url character varying(256) COLLATE pg_catalog."default" NOT NULL,
    access_token text COLLATE pg_catalog."default" NOT NULL,
    refresh_token text COLLATE pg_catalog."default" NOT NULL,
    token_expires_at timestamp with time zone,
    login_count integer,
    last_login_at timestamp with time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_provider_extern_id UNIQUE (provider, extern_id),
    CONSTRAINT users_provider_login UNIQUE (provider, login),
    CONSTRAINT users_provider_email UNIQUE (provider, email)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.users
    OWNER to postgres;

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_provider_email
    ON public.users USING btree
    (provider,email COLLATE pg_catalog."default")
    ;

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_provider_login
    ON public.users USING btree
    (provider,login COLLATE pg_catalog."default")
    ;

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_extern_id
    ON public.users USING btree
    (provider,extern_id)
    ;

SELECT create_constraint_if_not_exists(
    'public',
    'users',
    'unique_users_provider_login',
    'ALTER TABLE public.users ADD CONSTRAINT unique_users_provider_login UNIQUE USING INDEX idx_users_provider_login;');

SELECT create_constraint_if_not_exists(
    'public',
    'users',
    'unique_users_provider_email',
    'ALTER TABLE public.users ADD CONSTRAINT unique_users_provider_email UNIQUE USING INDEX idx_users_provider_email;');

--
-- ORGANIZATIONS
--
CREATE TABLE IF NOT EXISTS public.orgs
(
    id bigint NOT NULL,
    extern_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    reg_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    provider character varying(16) COLLATE pg_catalog."default" NOT NULL,
    login character varying(64) COLLATE pg_catalog."default" NOT NULL,
    name character varying(64) COLLATE pg_catalog."default" NOT NULL,
    email character varying(160) COLLATE pg_catalog."default" NOT NULL,
    billing_email character varying(160) COLLATE pg_catalog."default" NOT NULL,
    company character varying(64) COLLATE pg_catalog."default" NOT NULL,
    location character varying(64) COLLATE pg_catalog."default" NOT NULL,
    avatar_url character varying(256) COLLATE pg_catalog."default" NOT NULL,
    html_url character varying(256) COLLATE pg_catalog."default" NOT NULL,
    type character varying(16) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    street_address character varying(256) COLLATE pg_catalog."default" NOT NULL,
    city character varying(32) COLLATE pg_catalog."default" NOT NULL,
    postal_code character varying(16) COLLATE pg_catalog."default" NOT NULL,
    region character varying(16) COLLATE pg_catalog."default" NOT NULL,
    country character varying(16) COLLATE pg_catalog."default" NOT NULL,
    phone character varying(32) COLLATE pg_catalog."default" NOT NULL,
    approver_email character varying(160) COLLATE pg_catalog."default" NOT NULL,
    approver_name character varying(64) COLLATE pg_catalog."default" NOT NULL,
    status character varying(32) COLLATE pg_catalog."default" NOT NULL,
    expires_at timestamp with time zone,
    CONSTRAINT orgs_pkey PRIMARY KEY (id),
    CONSTRAINT orgs_provider_extern_id UNIQUE (provider, extern_id),
    CONSTRAINT orgs_provider_login UNIQUE (provider, login)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.orgs
    OWNER to postgres;

CREATE INDEX IF NOT EXISTS idx_orgs_provider
    ON public.orgs USING btree
    (provider);

CREATE INDEX IF NOT EXISTS idx_orgs_email
    ON public.orgs USING btree
    (email);

CREATE INDEX IF NOT EXISTS idx_orgs_phone
    ON public.orgs USING btree
    (phone);

--
-- Org Members
--

CREATE TABLE IF NOT EXISTS public.orgmembers
(
    id bigint NOT NULL,
    org_id bigint NOT NULL REFERENCES public.orgs ON DELETE RESTRICT,
    user_id bigint NOT NULL REFERENCES public.users ON DELETE RESTRICT,
    role character varying(64) COLLATE pg_catalog."default" NOT NULL,
    source character varying(16) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT orgmembers_pkey PRIMARY KEY (id),
    CONSTRAINT membership UNIQUE (org_id, user_id)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.orgmembers
    OWNER to postgres;

CREATE INDEX IF NOT EXISTS idx_orgmembers_team_id
    ON public.orgmembers USING btree
    (org_id ASC NULLS LAST);

CREATE INDEX IF NOT EXISTS idx_orgmembers_user_id
    ON public.orgmembers USING btree
    (user_id ASC NULLS LAST);

--
-- REPOS
--
CREATE TABLE IF NOT EXISTS public.repos
(
    id bigint NOT NULL,
    org_id bigint NOT NULL,
    extern_id bigint NOT NULL,
    provider character varying(16) COLLATE pg_catalog."default" NOT NULL,
    name character varying(64) COLLATE pg_catalog."default" NOT NULL,
    email character varying(160) COLLATE pg_catalog."default" NOT NULL,
    company character varying(64) COLLATE pg_catalog."default" NULL,
    avatar_url character varying(256) COLLATE pg_catalog."default" NULL,
    type character varying(16) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT repos_pkey PRIMARY KEY (id),
    CONSTRAINT repos_provider_id UNIQUE (extern_id, provider),
    CONSTRAINT repos_org_name UNIQUE (org_id, name)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.repos
    OWNER to postgres;

CREATE INDEX IF NOT EXISTS idx_repos_orgid
    ON public.repos USING btree
    (org_id)
    ;

CREATE INDEX IF NOT EXISTS idx_repos_provider
    ON public.repos USING btree
    (provider);

--
-- FRN
--
CREATE TABLE IF NOT EXISTS public.fcc_frn
(
    filer_id bigint NOT NULL,
    json text COLLATE pg_catalog."default" NULL,
    updated_at timestamp with time zone,
    CONSTRAINT fcc_frn_pkey PRIMARY KEY (filer_id)
)
WITH (
    OIDS = FALSE
);

CREATE INDEX IF NOT EXISTS idx_fcc_frn_updated_at
    ON public.fcc_frn USING btree
    (updated_at);

--
-- FRN Contact
--
CREATE TABLE IF NOT EXISTS public.fcc_contact
(
    id bigint NOT NULL,
    frn character varying(16) COLLATE pg_catalog."default" NOT NULL,
    json text COLLATE pg_catalog."default" NOT NULL,
    updated_at timestamp with time zone,
    CONSTRAINT fcc_contact_pkey PRIMARY KEY (id),
    CONSTRAINT fcc_contact_frn UNIQUE (frn)
)
WITH (
    OIDS = FALSE
);

CREATE INDEX IF NOT EXISTS idx_fcc_contact_updated_at
    ON public.fcc_contact USING btree
    (updated_at);

CREATE UNIQUE INDEX IF NOT EXISTS idx_fcc_contact_frn
    ON public.fcc_contact USING btree
    (frn);

--
-- Org Tokens
--
CREATE TABLE IF NOT EXISTS public.orgtokens
(
    id bigint NOT NULL,
    org_id bigint NOT NULL,
    requestor_id bigint NOT NULL,
    approver_email character varying(160) COLLATE pg_catalog."default" NOT NULL,
    token character varying(16) COLLATE pg_catalog."default" NOT NULL,
    code character varying(6) COLLATE pg_catalog."default" NOT NULL,
    used boolean NOT NULL,
    created_at timestamp with time zone,
    expires_at timestamp with time zone,
    used_at timestamp with time zone,
    CONSTRAINT orgtokens_pkey PRIMARY KEY (id),
    CONSTRAINT orgtokens_token_code UNIQUE (token, code)
)
WITH (
    OIDS = FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_orgtokens_token_code
    ON public.orgtokens USING btree
    (token,code);

CREATE INDEX IF NOT EXISTS idx_orgtokens_org_id
    ON public.orgtokens USING btree
    (org_id);

--
-- API Keys
--
CREATE TABLE IF NOT EXISTS public.apikeys
(
    id bigint NOT NULL,
    org_id bigint NOT NULL,
    key character varying(64) COLLATE pg_catalog."default" NOT NULL,
    enrollment boolean NOT NULL,
    management boolean NOT NULL,
    billing boolean NOT NULL,
    created_at timestamp with time zone,
    expires_at timestamp with time zone,
    used_at timestamp with time zone,
    CONSTRAINT apikeys_pkey PRIMARY KEY (id),
    CONSTRAINT apikeys_key UNIQUE (key)
)
WITH (
    OIDS = FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_apikeys_key
    ON public.apikeys USING btree
    (key);
--
--  Subscriptions
--
CREATE TABLE IF NOT EXISTS public.subscriptions
(
    id bigint NOT NULL,
    external_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    user_id bigint NOT NULL,
    customer_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    price_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    price_amount bigint NOT NULL,
    price_currency character varying(32) COLLATE pg_catalog."default" NOT NULL,
    payment_method_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone,
    expires_at timestamp with time zone,
    last_paid_at timestamp with time zone,
    status character varying(32) COLLATE pg_catalog."default" NOT NULL
)
WITH (
    OIDS = FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_subscriptions_id_user_id
    ON public.subscriptions USING btree
    (id,user_id);

CREATE UNIQUE INDEX IF NOT EXISTS idx_subscriptions_external_id
    ON public.subscriptions USING btree
    (external_id);

--
--
--
COMMIT;