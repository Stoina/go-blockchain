-- Table: public.bc_changehistory

-- DROP TABLE public.bc_changehistory;

CREATE TABLE public.bc_changehistory
(
    bcch_id integer NOT NULL DEFAULT nextval('bc_changehistory_bcch_id_seq'::regclass),
    bcch_objectid integer NOT NULL,
    bcch_changeid integer NOT NULL,
    bcch_description text COLLATE pg_catalog."default" NOT NULL,
    bcch_appname text COLLATE pg_catalog."default" NOT NULL,
    bcch_date timestamp with time zone,
    CONSTRAINT bc_changehistory_pkey PRIMARY KEY (bcch_id)
)

TABLESPACE pg_default;

ALTER TABLE public.bc_changehistory
    OWNER to postgres;