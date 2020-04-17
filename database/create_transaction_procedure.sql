CREATE OR REPLACE PROCEDURE public.proc_create_transaction(
	in_id uuid,
	in_transmitter_id text,
	in_receiver_id text,
	in_creditvalue decimal,
    in_status int,
	in_hash text)
LANGUAGE 'plpgsql'

AS $BODY$DECLARE

transmitter_exists integer;
receiver_exists integer;

BEGIN

SELECT INTO transmitter_exists count(*)
  FROM public.bc_participant
 WHERE bcp_id = in_transmitter_id;

IF transmitter_exists == true AND receiver_exists == true THEN

	INSERT INTO public.bc_participant(bct_id, bct_transmitter_id, bct_receiver_id, bct_creditvalue, bct_status, bct_hash)
		VALUES (in_id, in_transmitter_id, in_receiver_id, in_creditvalue, in_status, in_hash);
		
	RAISE NOTICE 'Successfully created new transaction';

END IF;

RAISE NOTICE 'Transmitter or receiver does not exists!';

END$BODY$;
