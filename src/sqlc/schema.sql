CREATE TABLE users (
  id         varchar(24) PRIMARY KEY,
  name       text NOT NULL,
  email      text NOT NULL UNIQUE,
  password   text NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
CREATE TABLE items (
  id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name        text NOT NULL,
  description text,
  price       bigint NOT NULL,
  created_at  timestamp NOT NULL DEFAULT now(),
  updated_at  timestamp NOT NULL DEFAULT now()
);
CREATE TABLE orders (
  id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id    varchar(24) REFERENCES users(id),
  item_id    uuid REFERENCES items(id),
  quantity   int NOT NULL,
  created_at timestamp NOT NULL DEFAULT now()
);
