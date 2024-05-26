CREATE INDEX merchant_category ON merchant(category);
CREATE INDEX merchant_location ON merchant USING GIST(geom);
CREATE INDEX merchant_created_at ON merchant(created_at);