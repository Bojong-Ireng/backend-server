DROP TABLE IF EXISTS login_credentials;
DROP TABLE IF EXISTS google_credentials;
DROP TABLE IF EXISTS walk_track_histories;
DROP TABLE IF EXISTS bus_track_histories;
DROP TABLE IF EXISTS buses;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    user_id CHAR(36),
    username VARCHAR(50),
    total_points INT8,
    registered_on DATE,
    phone_number VARCHAR(20),
    PRIMARY KEY (user_id)
);

CREATE TABLE buses (
    bus_id ChAR(36),
    bus_name VARCHAR(50),
    company_name VARCHAR(50) NOT NULL,
    PRIMARY KEY (bus_id)
);

CREATE TABLE walk_track_histories (
    walk_history_id SERIAL,
    path JSON,
    distance_traveled FLOAT4,
    duration INTERVAL,
    timestamp TIMESTAMP,
    user_id CHAR(36),
    PRIMARY KEY (walk_history_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE bus_track_histories (
    bus_history_id SERIAL,
    path JSON,
    distance_traveled FLOAT4,
    duration INTERVAL,
    timestamp TIMESTAMP,
    user_id CHAR(36),
    bus_id CHAR(36),
    PRIMARY KEY (bus_history_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (bus_id) REFERENCES buses(bus_id)
);

CREATE TABLE google_credentials (
    user_id CHAR(36),
    email VARCHAR(50) NOT NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE login_credentials (
    user_id CHAR(36),
    email VARCHAR(50) NOT NULL,
    password CHAR(60) NOT NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);