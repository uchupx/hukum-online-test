CREATE TABLE members (
  id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);
CREATE TABLE gatherings (
  id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  creator VARCHAR(255) NOT NULL,
  location VARCHAR(255) NOT NULL,
  scheduled_at TIMESTAMP NOT NULL,
  name VARCHAR(255) NOT NULL,
  type VARCHAR(255) NOT NULL
);
CREATE TABLE attendees (
  member_id BIGINT NOT NULL,
  gathering_id BIGINT NOT NULL,
  PRIMARY KEY (member_id, gathering_id),
  FOREIGN KEY (member_id) REFERENCES members(id),
  FOREIGN KEY (gathering_id) REFERENCES gatherings(id)
);
CREATE TABLE invitations (
  id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  member_id BIGINT NOT NULL,
  gathering_id BIGINT NOT NULL,
  status VARCHAR(255) NOT NULL,
  FOREIGN KEY (member_id) REFERENCES members(id),
  FOREIGN KEY (gathering_id) REFERENCES gatherings(id)
);