CREATE TABLE users
(
    id      serial       not null unique,
    name    varchar(255) not null,
    balance INT          not null,
);

CREATE TABLE quests
(
    id   serial       not null unique,
    name varchar(255) not null,
    cost INT          not null,
);

CREATE TABLE complete_quests
(
    user_id  INT NOT NULL,
    quest_id INT NOT NULL,
);