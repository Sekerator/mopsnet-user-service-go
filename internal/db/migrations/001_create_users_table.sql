-- Write your migrate up statements here

    CREATE EXTENSION IF NOT EXISTS "pgcrypto";

    CREATE TABLE users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        username VARCHAR(127) NOT NULL UNIQUE,
        password_hash VARCHAR(127) NOT NULL,
        auth_token VARCHAR(127),
        email VARCHAR(63),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

    CREATE OR REPLACE FUNCTION update_updated_at()
        RETURNS TRIGGER AS $$
    BEGIN
        NEW.updated_at = CURRENT_TIMESTAMP;
        RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;

    CREATE TRIGGER set_updated_at
        BEFORE UPDATE ON users
        FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

---- create above / drop below ----

---- create above / drop below ----

    DROP TRIGGER IF EXISTS set_updated_at ON users;
    DROP FUNCTION IF EXISTS update_updated_at;
    DROP TABLE IF EXISTS users;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
