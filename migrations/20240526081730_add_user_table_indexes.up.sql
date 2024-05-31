CREATE UNIQUE INDEX user_username ON users(username);
CREATE INDEX user_email_and_is_admin ON users(email, is_admin);