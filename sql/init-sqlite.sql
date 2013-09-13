ATTACH "salty.db" AS salty;

DROP TABLE IF EXISTS salty.user;
CREATE TABLE
    salty.user (
        id          INTEGER PRIMARY KEY AUTOINCREMENT,
        name        TEXT NOT NULL UNIQUE,
        email       TEXT NOT NULL UNIQUE,
        password    TEXT NOT NULL,
        join_date   INTEGER DEFAULT CURRENT_TIMESTAMP,
        last_login  INTEGER DEFAULT CURRENT_TIMESTAMP,
        balance     INTEGER
    );

DROP TABLE IF EXISTS salty.eventstatus;
CREATE TABLE
    salty.eventstatus (
        id      INTEGER PRIMARY KEY AUTOINCREMENT,
        code    INTEGER,
        status  TEXT
    );

INSERT INTO salty.eventstatus(code, status) VALUES (0, "Closed");
INSERT INTO salty.eventstatus(code, status) VALUES (1, "Open");
INSERT INTO salty.eventstatus(code, status) VALUES (2, "Finished");

DROP TABLE IF EXISTS salty.event;
CREATE TABLE
    salty.event (
        id      INTEGER PRIMARY KEY AUTOINCREMENT,
        status  INTEGER,
        created INTEGER DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY(status) REFERENCES "salty.eventstatus"(id)
    );

DROP TABLE IF EXISTS salty.stream;
CREATE TABLE
    salty.stream (
        id      INTEGER PRIMARY KEY AUTOINCREMENT,
        name    TEXT NOT NULL,
        url     TEXT NOT NULL UNIQUE
    );

DROP TABLE IF EXISTS salty.entrant;
CREATE TABLE
    salty.entrant (
        name    TEXT NOT NULL,
        stream  INTEGER,
        FOREIGN KEY(stream) REFERENCES "salty.stream"(id)
    );

DROP TABLE IF EXISTS salty.bet;
CREATE TABLE
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
 
DROP TABLE IF EXISTS salty.participant;
CREATE TABLE
    salty.participant (
        event   INTEGER,
        entrant INTEGER,
        rank    INTEGER,
        FOREIGN KEY(event) REFERENCES "salty.event"(id),
        FOREIGN KEY(entrant) REFERENCES "salty.entrant"(id)
    );
