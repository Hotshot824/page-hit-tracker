CREATE TABLE count (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL CHECK(LENGTH(url) = 32),
    count INTEGER,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
);