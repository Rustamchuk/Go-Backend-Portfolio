CREATE TABLE actors
(
    id         serial       not null unique,
    name       varchar(255) not null,
    gender     varchar(255) not null,
    birth_date date         not null
);

CREATE TABLE movies
(
    id           serial       not null unique,
    title        varchar(150) not null,
    description  text,
    release_date DATE         NOT NULL,
    rating       DECIMAL(2, 1) CHECK (rating >= 0 AND rating <= 10)
);

CREATE TABLE movie_actors
(
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,
    PRIMARY KEY (movie_id, actor_id),
    FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE,
    FOREIGN KEY (actor_id) REFERENCES actors (id) ON DELETE CASCADE
);