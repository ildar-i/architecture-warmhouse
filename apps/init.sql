-- Create the database if it doesn't exist
SELECT 'CREATE DATABASE smarthome'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'smarthome')\gexec

-- Connect to the database
\c smarthome;

-- Create the sensors table
CREATE TABLE IF NOT EXISTS sensors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    location VARCHAR(100) NOT NULL,
    value FLOAT DEFAULT 0,
    unit VARCHAR(20),
    status VARCHAR(20) NOT NULL DEFAULT 'inactive',
    last_updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Create indexes for common queries
CREATE INDEX IF NOT EXISTS idx_sensors_type ON sensors(type);
CREATE INDEX IF NOT EXISTS idx_sensors_location ON sensors(location);
CREATE INDEX IF NOT EXISTS idx_sensors_status ON sensors(status);

truncate sensors;

insert into sensors (name, type, location, value, unit)
values ('Батарея отопления', 'heat', '2124131231.423423423', random(), '1');

insert into sensors (name, type, location, value, unit)
values ('Кондиционер', 'heat', '2124131477.423423488', random(), '2');

insert into sensors (name, type, location, value, unit)
values ('Печка', 'heat', '54124131477.7634234358', random(), '3');