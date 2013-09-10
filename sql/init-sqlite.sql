ATTACH "salty.db" as salty;

CREATE TABLE
    IF NOT EXISTS
    salty.user (
        id          INTEGER PRIMARY KEY AUTOINCREMENT,
        name        TEXT,
        password    TEXT,
        email       TEXT,
        join_date   INTEGER,
        last_login  INTEGER,
        balance     INTEGER
    );

CREATE TABLE
    IF NOT EXISTS
    salty.event (
        id      INTEGER PRIMARY KEY AUTOINCREMENT,
        status  TEXT,
        created TEXT
    );

CREATE TABLE
    IF NOT EXISTS
    salty.stream (
        id      INTEGER PRIMARY KEY AUTOINCREMENT,
        name    TEXT,
        url     TEXT
    );

CREATE TABLE
    IF NOT EXISTS
    salty.entrant (
        name    TEXT,
        stream  INTEGER,
        FOREIGN KEY(stream) REFERENCES "salty.stream"(id)
    );

CREATE TABLE
    IF NOT EXISTS
    salty.bet (
        user    INTEGER,
        event   INTEGER,
        entrant INTEGER,
        amount  INTEGER,
        rank    INTEGER,
        FOREIGN KEY(user) REFERENCES "salty.user"(id),
        FOREIGN KEY(event) REFERENCES "salty.event"(id),
        FOREIGN KEY(entrant) REFERENCES "salty.entrant"(id)
    );
 
CREATE TABLE
    IF NOT EXISTS
    salty.participant (
        event   INTEGER,
        entrant INTEGER,
        rank    INTEGER,
        FOREIGN KEY(event) REFERENCES "salty.event"(id),
        FOREIGN KEY(entrant) REFERENCES "salty.entrant"(id)
    );
