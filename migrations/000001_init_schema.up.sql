CREATE TYPE role_enum AS ENUM ('moderator', 'user');
CREATE TYPE content_type_enum AS ENUM ('movie', 'show');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role role_enum NOT NULL DEFAULT 'user'
);

CREATE TABLE contents (
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