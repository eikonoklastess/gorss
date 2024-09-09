CREATE TABLE IF NOT EXISTS channels (
    id INTEGER PRIMARY KEY AUTOINCREMENT,   -- Unique ID for each channel
    title TEXT NOT NULL,                    -- Title of the RSS channel
    description TEXT,                       -- Description of the RSS channel
    link TEXT NOT NULL,                     -- URL of the RSS feed
    last_updated DATETIME,                  -- Last time the feed was updated
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP -- Timestamp when the channel was added
);
