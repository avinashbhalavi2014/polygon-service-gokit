CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS areas (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(26) NOT NULL,
    polygon GEOMETRY(Polygon, 4326) NOT NULL
);

CREATE INDEX polygon_geom_idx on areas USING GIST (polygon); 