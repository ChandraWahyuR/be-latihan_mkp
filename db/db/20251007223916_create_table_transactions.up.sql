DO $$ BEGIN
    CREATE TYPE status_transaction AS ENUM('pending', 'expire', 'failed', 'success');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS transactions(
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    user_id VARCHAR(50),
    movie_id VARCHAR(50),
    studio_seat_id VARCHAR(50),
    status status_transaction,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT fk_users_t FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT fk_movies_t FOREIGN KEY (movie_id)
        REFERENCES movies(id) ON DELETE CASCADE,

    CONSTRAINT fk_seats_t FOREIGN KEY (studio_seat_id)
        REFERENCES studio_seats(id) ON DELETE CASCADE
);