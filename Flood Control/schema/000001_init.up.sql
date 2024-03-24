CREATE TABLE flood
(
    id      SERIAL PRIMARY KEY,
    user_id BIGINT    NOT NULL,
    time    TIMESTAMP NOT NULL
);