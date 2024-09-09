CREATE TABLE IF NOT EXISTS channels (
    id INTEGER PRIMARY KEY AUTOINCREMENT,   -- Unique ID for each channel
    title TEXT NOT NULL,                    -- Title of the RSS channel
    description TEXT,                       -- Description of the RSS channel
    link TEXT NOT NULL,                     -- URL of the RSS feed
    last_updated DATETIME,                  -- Last time the feed was updated
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP -- Timestamp when the channel was added
);

CREATE TABLE IF NOT EXISTS posts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique ID for each item
    channel_id INTEGER NOT NULL,           -- ID of the RSS channel this item belongs to
    title TEXT NOT NULL,                   -- Title of the RSS item
    description TEXT,                      -- Description of the RSS item
    link TEXT NOT NULL,                    -- URL of the RSS item
    pub_date DATETIME,                     -- Publication date of the RSS item
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- When the item was added to the database
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE -- Foreign key to rss_channels table
);
