-- Table: public.bc_participant

-- DROP TABLE public.bc_participant;

CREATE TABLE public.bc_participant
(
    bcp_id uuid NOT NULL,
    bcp_email text COLLATE pg_catalog."default" NOT NULL,
    bcp_lastname text COLLATE pg_catalog."default" NOT NULL,
    bcp_firstname text COLLATE pg_catalog."default" NOT NULL,
    bcp_hash text COLLATE pg_catalog."default",
    CONSTRAINT bc_participant_pkey PRIMARY KEY (bcp_id)
)

TABLESPACE pg_default;

ALTER TABLE public.bc_participant
    OWNER to blockchain_admin;
COMMENT ON TABLE public.bc_participant
    IS 'Participant table for blockchain network';