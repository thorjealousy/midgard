
-- +migrate Up

CREATE TABLE events (
    time            TIMESTAMPTZ     NOT NULL,
    id              BIGSERIAL       NOT NULL,
    height          BIGINT          NOT NULL,
    type            VARCHAR         NOT NULL,
    status          VARCHAR,
    PRIMARY KEY (id, time)
);

CREATE TABLE pools_history (
    time            TIMESTAMPTZ     NOT NULL,
    id              BIGSERIAL       NOT NULL,
    event_id        BIGINT          NOT NULL,
    event_type      VARCHAR         NOT NULL,
    pool            VARCHAR         NOT NULL,
    asset_amount    BIGINT          NOT NULL,
    rune_amount     BIGINT          NOT NULL,
    units           BIGINT,
    status          SMALLINT        NOT NULL,
    PRIMARY KEY (id, time)
);
CREATE INDEX event_id_pools_history_idx ON pools_history (event_id);
CREATE INDEX pool_pools_history_idx ON pools_history (pool);

CREATE TYPE swap_direction as enum('buy', 'sell');
CREATE TABLE swaps (
    time            TIMESTAMPTZ     NOT NULL,
    event_id        BIGINT          NOT NULL,
    from_address    VARCHAR         NOT NULL,
    to_address      VARCHAR         NOT NULL,
    pool            VARCHAR         NOT NULL,
    price_target    BIGINT,
    trade_slip      REAL,
    liquidity_fee   BIGINT,
    second_event_id BIGINT,
    direction       swap_direction  NOT NULL,
    PRIMARY KEY (event_id, time)
);
CREATE INDEX idx_swaps ON swaps (from_address, pool);

CREATE TYPE tx_direction as enum('in', 'out');
CREATE TABLE txs (
    time        TIMESTAMPTZ       NOT NULL,
    id SERIAL,
    tx_hash varchar not null,
    event_id bigint not null,
    direction tx_direction not null,
    chain varchar,
    from_address varchar,
    to_address varchar,
    memo varchar,
    primary key (id, time, event_id)
);

CREATE TABLE coins (
    time        TIMESTAMPTZ       NOT NULL,
    id SERIAL,
    tx_hash varchar not null,
    event_id bigint not null,
    chain varchar not null,
    symbol varchar not null,
    ticker varchar not null,
    amount bigint not null,
    primary key (id, time, event_id)
);

CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
SELECT create_hypertable('events', 'time');
SELECT create_hypertable('pools_history', 'time');
SELECT create_hypertable('swaps', 'time');
SELECT create_hypertable('txs', 'time');
SELECT create_hypertable('coins', 'time');

-- +migrate Down

DROP TABLE events;
DROP TABLE pools_history;
DROP TABLE swaps;
DROP TABLE txs;
DROP TABLE coins;

DROP TYPE tx_direction;
DROP TYPE swap_direction;
