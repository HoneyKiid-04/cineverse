ALTER TABLE users
ADD COLUMN bio TEXT,
ADD COLUMN avatar_url TEXT;

-- Update existing users with default values
UPDATE users
SET bio = '',
    avatar_url = ''
WHERE bio IS NULL
   OR avatar_url IS NULL;