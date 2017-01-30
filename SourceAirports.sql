TRUNCATE TABLE SourceAirports;

CREATE TABLE SourceAirports (
  SourceAirportID int NOT NULL AUTO_INCREMENT,
  SourceAirportName varchar(255) NOT NULL,
  SourceAirportCode varchar(255) NOT NULL UNIQUE,
  PRIMARY KEY(SourceAirportID)
);

INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('London Heathrow','LHR');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('London Gatwick','LGW');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('London Luton','LTN');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('Birmingham International','BHX');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('Manchester','MAN');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('Edinburgh','EDI');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('Glasgow','GLA');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('East Midlands','EMA');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('Bristol','BRS');
INSERT INTO SourceAirports (SourceAirportName, SourceAirportCode) VALUES ('Newcastle','NCL');

SELECT * FROM SourceAirports;
