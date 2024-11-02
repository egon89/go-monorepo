CREATE TABLE IF NOT EXISTS exchange (
    code TEXT,
    bid REAL,
    high REAL,
    low REAL,
    var_bid REAL,
    pct_change REAL,
    ask REAL,
    timestamp INTEGER,
    create_date DATETIME
);
