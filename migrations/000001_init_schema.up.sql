-- 1. CATEGORIES
CREATE TABLE categories(
  id SERIAL PRIMARY KEY, 
  name VARCHAR(40),
  description TEXT DEFAULT NULL
);

-- 2. BOOKS
CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(40),
  subtitle TEXT,
  about_the_book TEXT,
  page_count INTEGER,
  price NUMERIC,
  image TEXT,
  language VARCHAR(40),
  author_name VARCHAR(40),
  author_avatar TEXT,

  category_id INTEGER REFERENCES categories(id)
);

-- 3. USERS
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(128) NOT NULL,
  is_staff BOOLEAN NOT NULL
);

-- 4. WHISLIST
CREATE TABLE wishlist (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  book_id INTEGER REFERENCES books(id),
  quantity INTEGER NOT NULL,
  added_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 5. CART
CREATE TABLE cart (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  book_id INTEGER REFERENCES books(id),
  quantity INTEGER NOT NULL,
  added_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 6. ORDERS
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  book_id INTEGER REFERENCES books(id),
  quantity INTEGER NOT NULL,
  price NUMERIC NOT NULL,
  order_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
