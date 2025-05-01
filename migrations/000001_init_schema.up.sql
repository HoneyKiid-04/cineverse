DO $$
BEGIN
    -- Check and create role_enum type if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
        CREATE TYPE role_enum AS ENUM ('moderator', 'user');
    END IF;

    -- Check and create content_type_enum type if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'content_type_enum') THEN
        CREATE TYPE content_type_enum AS ENUM ('movie', 'show');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role role_enum NOT NULL DEFAULT 'user'
);

CREATE TABLE IF NOT EXISTS contents (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    type content_type_enum NOT NULL,
    release_date TIMESTAMP WITH TIME ZONE,
    duration INTEGER,
    rating REAL,
    genres VARCHAR(255),
    director VARCHAR(255),
    cast_members TEXT,
    poster_url TEXT,
    seasons INTEGER,
    episodes INTEGER
);

-- Insert default users if they don't exist
DO $$
BEGIN
    -- Insert admin user if not exists
    IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin') THEN
        INSERT INTO users (username, email, password, role)
        VALUES (
            'admin',
            'admin@cineverse.com',
            -- Password: admin123 (bcrypt hashed)
            '$2a$12$fNJHRd01RaqxgfyLfcci9OCISbA5.LX7AGy94oCzSlYPZKjj1IByy',
            'moderator'
        );
    END IF;

    -- Insert default user if not exists
    IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'user') THEN
        INSERT INTO users (username, email, password, role)
        VALUES (
            'user',
            'user@cineverse.com',
            -- Password: user123 (bcrypt hashed)
            '$2a$12$p92M9Uc7pXm/gOPE/j.DX.fAG2NCGlOun2nBLqLpH4Eb5y2AXPPVG',
            'user'
        );
    END IF;
END$$;