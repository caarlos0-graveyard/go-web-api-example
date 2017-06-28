CREATE SEQUENCE seq_beers;
CREATE TABLE beers(
  id bigint primary key default nextval('seq_beers'),
  name varchar(60) not null,
  price double precision,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
