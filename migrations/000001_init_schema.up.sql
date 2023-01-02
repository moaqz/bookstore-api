-- 1. CATEGORIES
CREATE TABLE categories(
  id SERIAL PRIMARY KEY, 
  name VARCHAR(40),
  description TEXT DEFAULT NULL
);

-- 2. AUTHORS
CREATE TABLE authors (
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  avatar TEXT NOT NULL,
  github_name VARCHAR(50) DEFAULT NULL,
  instagram_name VARCHAR(50) DEFAULT NULL,
  linkedin_name VARCHAR(50) DEFAULT NULL,
  twitter_name VARCHAR(50) DEFAULT NULL
);

-- 3. BOOKS
CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(40),
  subtitle TEXT,
  about_the_book TEXT,
  page_count INTEGER,
  price NUMERIC(2),
  image TEXT,
  published_at DATE,
  language VARCHAR(40),

  author_id INTEGER,
  category_id INTEGER
);

ALTER TABLE books ADD FOREIGN KEY (author_id) REFERENCES authors (id);
ALTER TABLE books ADD FOREIGN KEY (category_id) REFERENCES categories (id);