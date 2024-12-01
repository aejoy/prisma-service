CREATE TABLE photo
(
    id        VARCHAR(16)  NOT NULL,
    creator   VARCHAR(128) NOT NULL,
    "to"        VARCHAR(128)          DEFAULT '',
    url       VARCHAR(128) NOT NULL,
    blur_hash VARCHAR(128) NOT NULL,
    height    SMALLINT     NOT NULL,
    width     SMALLINT     NOT NULL,
    size      INT          NOT NULL,
    published TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated   TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    archived  TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '90 days'
);