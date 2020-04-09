-- PROCEDURE: public.proc_create_or_update_participant(uuid, text, text, text, text)

-- DROP PROCEDURE public.proc_create_or_update_participant(uuid, text, text, text, text);

CREATE OR REPLACE PROCEDURE public.proc_create_or_update_participant(
	in_id uuid,
	in_email text,
	in_lastname text,
	in_firstname text,
	in_hash text)
LANGUAGE 'plpgsql'

AS $BODY$DECLARE

participant_row_count integer;

BEGIN

SELECT INTO participant_row_count count(*)
  FROM public.bc_participant 
 WHERE bcp_email = in_email;

IF participant_row_count = 0 THEN
	INSERT INTO public.bc_participant(bcp_id, bcp_email, bcp_lastname, bcp_firstname, bcp_hash)
		VALUES (in_id, in_email, in_lastname, in_firstname, in_hash);
		
	RAISE NOTICE 'Successfully created new participant';
END IF;

IF participant_row_count > 0 THEN

	UPDATE public.bc_participant
	   SET bcp_lastname = in_lastname, bcp_firstname = in_firstname
	 WHERE bcp_email = in_email;

	RAISE NOTICE 'Successfully updated new participant(s)';
END IF;

END$BODY$;
