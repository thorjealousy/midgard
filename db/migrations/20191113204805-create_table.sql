
-- +migrate Up

CREATE TABLE events (
    time        TIMESTAMPTZ       not null,
    id integer not null,
    height integer not null,
    type varchar not null,
    status varchar not null,
    primary key (time, id)
);
CREATE TABLE stakes (
    time        TIMESTAMPTZ       NOT NULL,
    event_id integer not null ,
    chain integer not null,
    symbol varchar not null,
    ticker varchar not null,
    units integer,
    primary key (time, event_id)
);
CREATE TYPE tx_direction as enum('in', 'out');
CREATE TABLE txs (
    time        TIMESTAMPTZ       NOT NULL,
    tx_hash varchar not null,
    event_id integer not null,
    direction tx_direction not null,
    chain varchar,
    from_address varchar,
    to_address varchar,
    memo varchar,
    primary key (time, event_id, tx_hash)
);
CREATE TABLE coins (
    time        TIMESTAMPTZ       NOT NULL,
    tx_id integer not null,
    event_id integer not null,
    chain varchar not null,
    symbol varchar not null,
    ticker varchar not null,
    amount integer not null,
    primary key (time, tx_id, event_id)
);
CREATE TABLE gas (
    time        TIMESTAMPTZ       NOT NULL,
    event_id integer not null,
    chain varchar not null,
    symbol varchar not null,
    ticker varchar not null,
    amount integer not null,
    primary key (time, event_id)
);

SELECT create_hypertable('events', 'time');
SELECT create_hypertable('stakes', 'time');
SELECT create_hypertable('txs', 'time');
SELECT create_hypertable('coins', 'time');
SELECT create_hypertable('gas', 'time');

-- +migrate Down

DROP TABLE events;
DROP TABLE stakes;
DROP TABLE txs;
DROP TABLE coins;
DROP TABLE gas;

DROP TYPE tx_direction;