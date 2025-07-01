CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(30) NOT NULL UNIQUE,
                                     password_hash TEXT NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS habits (
                                      id SERIAL PRIMARY KEY,
                                      user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                      name TEXT NOT NULL UNIQUE,
                                      description TEXT,
                                      created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS habit_events (
                                            id SERIAL PRIMARY KEY,
                                            habit_id INTEGER NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
                                            description TEXT,
                                            occurred_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_habits_user_id ON habits(user_id);
CREATE INDEX IF NOT EXISTS idx_events_habit_id_occurred_at ON habit_events(habit_id, occurred_at DESC);
