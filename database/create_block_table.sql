-- Table: public.bc_block

-- DROP TABLE public.bc_block;

CREATE TABLE public.bc_block
(
    bcb_id uuid NOT NULL,
    bcb_gendate date NOT NULL,
    bcb_gentime time NOT NULL,
    bcb_hash text COLLATE pg_catalog."default",
    CONSTRAINT bc_block_pkey PRIMARY KEY (bcb_id)
)

TABLESPACE pg_default;

ALTER TABLE public.bc_block
    OWNER to blockchain_admin;
COMMENT ON TABLE public.bc_block
    IS 'Block table for blockchain network';