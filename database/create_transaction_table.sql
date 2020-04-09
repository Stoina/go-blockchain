-- Table: public.bc_transaction

-- DROP TABLE public.bc_transaction;

CREATE TABLE public.bc_transaction
(
    bct_id uuid NOT NULL,
    bct_block_id uuid NULL,
    bct_transmitter_id uuid NOT NULL,
    bct_receiver_id uuid NOT NULL,
    bct_creditvalue numeric NOT NULL,
    bct_hash text COLLATE pg_catalog."default",
    CONSTRAINT bc_transaction_pkey PRIMARY KEY (bct_id),
    CONSTRAINT bc_block_fkey FOREIGN KEY (bct_block_id) REFERENCES public.bc_block (bcb_id),
    CONSTRAINT bc_transmitter_fkey FOREIGN KEY (bct_transmitter_id) REFERENCES public.bc_participant (bcp_id),
    CONSTRAINT bc_receiver_fkey FOREIGN KEY (bct_receiver_id) REFERENCES public.bc_participant (bcp_id)
)

TABLESPACE pg_default;

ALTER TABLE public.bc_transaction
    OWNER to blockchain_admin;
COMMENT ON TABLE public.bc_transaction
    IS 'Transaction table for blockchain network';