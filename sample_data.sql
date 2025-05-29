-- Sample data for order-inventory-app
-- Creates a superuser account with username 'Admin' and password 'admin123'

INSERT INTO users (username, password, created_at, updated_at) VALUES (
  'Admin',
  '$2a$10$0GydG2B5nlf8Hb1IaB7A8u3ObqI2zWX0eGkXObMPeC9oy8nnQ9mBe',
  NOW(),
  NOW()
);

-- Note: password is bcrypt hashed version of 'admin123'
