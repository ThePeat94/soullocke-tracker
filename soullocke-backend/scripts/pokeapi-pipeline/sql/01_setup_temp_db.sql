DO
$do$
    BEGIN
        IF EXISTS (
            SELECT FROM pg_catalog.pg_roles
            WHERE  rolname = 'ash') THEN

            RAISE NOTICE 'Role "ash" already exists. Skipping.';
        ELSE
            CREATE ROLE ash;
        END IF;
    END
$do$;

CREATE DATABASE poke_restore;