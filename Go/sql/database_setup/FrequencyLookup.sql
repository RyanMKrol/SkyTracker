DROP TABLE FrequencyLookup;

CREATE TABLE FrequencyLookup (
  FLID int NOT NULL AUTO_INCREMENT,
  FLDays INT NOT NULL UNIQUE,
  PRIMARY KEY(FLID)
);

INSERT INTO FrequencyLookup (FLDays) VALUES (1);
INSERT INTO FrequencyLookup (FLDays) VALUES (7);
INSERT INTO FrequencyLookup (FLDays) VALUES (30);
