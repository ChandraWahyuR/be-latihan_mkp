DO $$ BEGIN
    CREATE TYPE status_seat AS ENUM('avaible', 'locked', 'paid');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS studio_seats(
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    studio_schedule_id VARCHAR(50),
    seat_numbers VARCHAR(255),
    status status_seat,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT fk_studio_schedule FOREIGN KEY (studio_schedule_id)
        REFERENCES studio_schedules(id) ON DELETE CASCADE
);