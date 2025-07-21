package database

import (
	"database/sql"
	"fmt"
	"log"

	// Switched to a CGO-free driver for easier compilation and portability.
	_ "modernc.org/sqlite"
)

const dbPath = "./project_management.db" // Path to your SQLite database file

// getSchema returns the SQL statements to create the database schema.
// No changes needed here, the schema is well-defined.
func getSchema() string {
	return `
-- A central table for all users (clients and developers) to handle authentication.
-- The 'role' column distinguishes between user types.
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL, -- Always store hashed passwords
    role TEXT NOT NULL CHECK(role IN ('client', 'developer')),
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- Stores profile information specific to clients.
-- This table has a one-to-one relationship with the 'users' table.
CREATE TABLE IF NOT EXISTS clients (
    user_id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    company TEXT,
    contact_number TEXT,
    address TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Stores profile information specific to developers.
CREATE TABLE IF NOT EXISTS developers (
    user_id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    experience_years INTEGER DEFAULT 0,
    linkedin_url TEXT,
    -- Using a clear status is better than a simple 'yes'/'no' string
    availability_status TEXT DEFAULT 'available' CHECK(availability_status IN ('available', 'unavailable')),
    address TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- A master table for all possible skills to avoid duplication.
CREATE TABLE IF NOT EXISTS skills (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- A linking table to create a many-to-many relationship between developers and skills.
-- This lets a developer have multiple skills, and a skill can be had by multiple developers.
CREATE TABLE IF NOT EXISTS developer_skills (
    developer_user_id INTEGER NOT NULL,
    skill_id INTEGER NOT NULL,
    PRIMARY KEY (developer_user_id, skill_id),
    FOREIGN KEY (developer_user_id) REFERENCES developers (user_id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE CASCADE
);

-- Stores details for each project posted by a client.
CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    summary TEXT,
    quotation_price REAL, -- REAL is used for floating-point numbers in SQLite
    created_date TEXT,
    deadline TEXT,
    priority TEXT CHECK(priority IN ('Low', 'Medium', 'High')),
    status TEXT DEFAULT 'Pending' CHECK(status IN ('Pending', 'In Progress', 'Completed', 'Cancelled')),
    repo_url TEXT,
    design_tool TEXT,
    notes TEXT,
    FOREIGN KEY (client_user_id) REFERENCES clients (user_id) ON DELETE CASCADE
);

-- A linking table for the many-to-many relationship between projects and their required skills.
CREATE TABLE IF NOT EXISTS project_required_skills (
    project_id INTEGER NOT NULL,
    skill_id INTEGER NOT NULL,
    PRIMARY KEY (project_id, skill_id),
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills (id) ON DELETE CASCADE
);

-- A table to store multiple wireframe URLs for a single project.
-- This is more normalized than storing a JSON array in the projects table.
CREATE TABLE IF NOT EXISTS project_wireframes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id INTEGER NOT NULL,
    url TEXT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

-- A table to assign developers to projects.
-- This creates a many-to-many relationship between projects and developers.
CREATE TABLE IF NOT EXISTS project_assignments (
    project_id INTEGER NOT NULL,
    developer_user_id INTEGER NOT NULL,
    assignment_date TEXT DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (project_id, developer_user_id),
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    FOREIGN KEY (developer_user_id) REFERENCES developers (user_id) ON DELETE CASCADE
);
`
}

// InitDB initializes the SQLite database.
func InitDB() (*sql.DB, error) {
	// Use "sqlite" as the driver name for modernc.org/sqlite
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Ping the database to ensure connection is established
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Execute the schema creation SQL
	if _, err = db.Exec(getSchema()); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Database initialized successfully at", dbPath)
	return db, nil
}

// The conflicting main function has been removed.
