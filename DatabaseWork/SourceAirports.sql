DROP TABLE SourceAirports;

CREATE TABLE SourceAirports (
  SrcAirportID int NOT NULL AUTO_INCREMENT,
  SrcAirportName varchar(255) NOT NULL,
  SrcAirportCode varchar(255) NOT NULL UNIQUE,
  PRIMARY KEY(SrcAirportID)
);

INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('London Heathrow','LHR');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('London Gatwick','LGW');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('London Luton','LTN');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('Birmingham International','BHX');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('Manchester','MAN');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('Edinburgh','EDI');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('Glasgow','GLA');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('East Midlands','EMA');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('Bristol','BRS');
INSERT INTO SourceAirports (SrcAirportName, SrcAirportCode) VALUES ('Newcastle','NCL');

SELECT * FROM SourceAirports;
