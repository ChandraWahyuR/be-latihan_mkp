CREATE TABLE IF NOT EXISTS studio_schedules(
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    movie_id VARCHAR(255),
    name_studio VARCHAR(255),  
    starting TIMESTAMPTZ DEFAULT NOW(),
    ending TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT fk_movies FOREIGN KEY (movie_id)
        REFERENCES movies(id) ON DELETE CASCADE
);