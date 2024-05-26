CREATE INDEX merchant_category ON merchants(category);
CREATE INDEX merchant_location ON merchants USING GIST(location);
CREATE INDEX merchant_created_at ON merchants(created_at);