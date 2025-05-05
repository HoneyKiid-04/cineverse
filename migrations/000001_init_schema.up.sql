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
    IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'moderator') THEN
        INSERT INTO users (username, email, password, role)
        VALUES (
            'moderator',
            'moderator@cineverse.com',
            -- Password: moder123 (bcrypt hashed)
            '$2a$12$3vxCywoqAheKfSoM.UPVd.1O3/Y2YP7teRd1FjKI24Dhfsl4LhFRq',
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

-- Insert sample contents


-- Insert sample contents
INSERT INTO contents (title, description, type, release_date, duration, rating, genres, director, cast_members, poster_url, seasons, episodes) 
VALUES 
    ('The Shawshank Redemption', 'Two imprisoned men bond over a number of years.', 'movie', '1994-09-23T00:00:00Z', 142, 9.3, 'Drama', 'Frank Darabont', 'Tim Robbins, Morgan Freeman', 'https://image.tmdb.org/t/p/original/q6y0Go1tsGEsmtFryDOJo3dEmqu.jpg', NULL, NULL),
    ('Breaking Bad', 'A high school chemistry teacher turned drug lord.', 'show', '2008-01-20T00:00:00Z', 45, 9.5, 'Crime, Drama', 'Vince Gilligan', 'Bryan Cranston, Aaron Paul', 'https://image.tmdb.org/t/p/original/ggFHVNu6YYI5L9pCfOacjizRGt.jpg', 5, 62),
    ('The Dark Knight', 'Batman faces his greatest challenge.', 'movie', '2008-07-18T00:00:00Z', 152, 9.0, 'Action, Crime, Drama', 'Christopher Nolan', 'Christian Bale, Heath Ledger', 'https://image.tmdb.org/t/p/original/qJ2tW6WMUDux911r6m7haRef0WH.jpg', NULL, NULL),
    ('Stranger Things', 'Supernatural events in a small town.', 'show', '2016-07-15T00:00:00Z', 50, 8.7, 'Drama, Fantasy, Horror', 'The Duffer Brothers', 'Millie Bobby Brown, Finn Wolfhard', 'https://image.tmdb.org/t/p/original/49WJfeN0moxb9IPfGn8AIqMGskD.jpg', 4, 34),
    ('Inception', 'A thief enters dreams to steal secrets.', 'movie', '2010-07-16T00:00:00Z', 148, 8.8, 'Action, Sci-Fi', 'Christopher Nolan', 'Leonardo DiCaprio, Joseph Gordon-Levitt', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSmkq35zEZtIgZWAHKWdGD_IRMUx-c9EOgrcQ&s', NULL, NULL),
    ('Game of Thrones', 'Noble families fight for control.', 'show', '2011-04-17T00:00:00Z', 60, 9.3, 'Action, Adventure, Drama', 'Various', 'Emilia Clarke, Kit Harington', 'https://image.tmdb.org/t/p/original/u3bZgnGQ9T01sWNhyveQz0wH0Hl.jpg', 8, 73),
    ('Pulp Fiction', 'Interconnected stories of criminals in LA.', 'movie', '1994-10-14T00:00:00Z', 154, 8.9, 'Crime, Drama', 'Quentin Tarantino', 'John Travolta, Uma Thurman', 'https://image.tmdb.org/t/p/original/d5iIlFn5s0ImszYzBPb8JPIfbXD.jpg', NULL, NULL),
    ('The Mandalorian', 'A lone bounty hunter''s adventures.', 'show', '2019-11-12T00:00:00Z', 40, 8.8, 'Action, Adventure, Sci-Fi', 'Jon Favreau', 'Pedro Pascal, Grogu', 'https://image.tmdb.org/t/p/original/sWgBv7LV2PRoQgkxwlibdGXKz1S.jpg', 3, 24),
    ('Interstellar', 'Space explorers search for a new home.', 'movie', '2014-11-07T00:00:00Z', 169, 8.6, 'Adventure, Drama, Sci-Fi', 'Christopher Nolan', 'Matthew McConaughey, Anne Hathaway', 'https://image.tmdb.org/t/p/original/gEU2QniE6E77NI6lCU6MxlNBvIx.jpg', NULL, NULL),
    ('The Office', 'Daily lives of office employees.', 'show', '2005-03-24T00:00:00Z', 22, 8.9, 'Comedy', 'Various', 'Steve Carell, Rainn Wilson', 'https://image.tmdb.org/t/p/original/qWnJzyZhyy74gjpSjIXWmuk0ifX.jpg', 9, 201),
    ('The Matrix', 'A computer programmer discovers reality.', 'movie', '1999-03-31T00:00:00Z', 136, 8.7, 'Action, Sci-Fi', 'The Wachowskis', 'Keanu Reeves, Laurence Fishburne', 'https://image.tmdb.org/t/p/original/f89U3ADr1oiB1s9GkdPOEpXUk5H.jpg', NULL, NULL),
    ('Black Mirror', 'Dark anthology about technology.', 'show', '2011-12-04T00:00:00Z', 60, 8.8, 'Drama, Sci-Fi, Thriller', 'Charlie Brooker', 'Various', 'https://image.tmdb.org/t/p/original/7PRddO7z7mcPi21nZTCMGShAyy1.jpg', 6, 27),
    ('Gladiator', 'A Roman general seeks justice.', 'movie', '2000-05-05T00:00:00Z', 155, 8.5, 'Action, Adventure, Drama', 'Ridley Scott', 'Russell Crowe, Joaquin Phoenix', 'https://image.tmdb.org/t/p/original/ty8TGRuvJLPUmAR1H1nRIsgwvim.jpg', NULL, NULL),
    ('Chernobyl', 'The 1986 nuclear disaster.', 'show', '2019-05-06T00:00:00Z', 60, 9.4, 'Drama, History, Thriller', 'Craig Mazin', 'Jared Harris, Stellan Skarsgård', 'https://image.tmdb.org/t/p/original/hlLXt2tOPT6RRnjiUmoxyG1LTFi.jpg', 1, 5),
    ('The Godfather', 'The Corleone crime family saga.', 'movie', '1972-03-24T00:00:00Z', 175, 9.2, 'Crime, Drama', 'Francis Ford Coppola', 'Marlon Brando, Al Pacino', 'https://image.tmdb.org/t/p/original/3bhkrj58Vtu7enYsRolD1fZdja1.jpg', NULL, NULL),
    ('Fargo', 'Crime anthology series.', 'show', '2014-04-15T00:00:00Z', 53, 8.9, 'Crime, Drama, Thriller', 'Noah Hawley', 'Various', 'https://myserencam.wordpress.com/wp-content/uploads/2021/01/fargo-896x1024-1.jpg', 5, 41),
    ('Forrest Gump', 'Life journey of a simple man.', 'movie', '1994-07-06T00:00:00Z', 142, 8.8, 'Drama, Romance', 'Robert Zemeckis', 'Tom Hanks, Robin Wright', 'https://image.tmdb.org/t/p/original/arw2vcBveWOVZr6pxd9XTd1TdQa.jpg', NULL, NULL),
    ('True Detective', 'Anthology crime series.', 'show', '2014-01-12T00:00:00Z', 55, 8.9, 'Crime, Drama, Mystery', 'Nic Pizzolatto', 'Matthew McConaughey, Woody Harrelson', 'https://m.media-amazon.com/images/M/MV5BYjgwYzA1NWMtNDYyZi00ZGQyLWI5NTktMDYwZjE2OTIwZWEwXkEyXkFqcGc@._V1_FMjpg_UX1000_.jpg', 3, 24),
    ('Fight Club', 'An insomniac office worker forms a club.', 'movie', '1999-10-15T00:00:00Z', 139, 8.8, 'Drama', 'David Fincher', 'Brad Pitt, Edward Norton', 'https://image.tmdb.org/t/p/original/pB8BM7pdSp6B6Ih7QZ4DrQ3PmJK.jpg', NULL, NULL),
    ('The Crown', 'British royal family drama.', 'show', '2016-11-04T00:00:00Z', 58, 8.7, 'Biography, Drama, History', 'Peter Morgan', 'Claire Foy, Olivia Colman', 'https://image.tmdb.org/t/p/original/6eehp9I54syN3x753XMqjKz8M3F.jpg', 6, 60);