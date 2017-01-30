DROP TABLE DestinationAirports;

CREATE TABLE DestinationAirports (
  DestAirportID int NOT NULL AUTO_INCREMENT,
  DestAirportName varchar(255) NOT NULL UNIQUE,
  DestAirportCode varchar(255) NOT NULL UNIQUE,
  DestCountry varchar(255) NOT NULL,
  DestCity varchar(255) NOT NULL,
  PRIMARY KEY(DestAirportID)
);

-- France
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Paris Charles de Gaulle','CDG','France','Paris');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Paris Orly','ORY','France','Paris');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Nice','NCE','France','Nice');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Lyon St Exupéry','LYS','France','Lyon');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Marseille','MRS','France','Marseille');

-- Spain
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Adolfo Suárez Madrid-Barajas','MAD','Spain','Madrid');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Barcelona-El Prat','BCN','Spain','Barcelona');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Palma De Mallorca','PMI','Spain','Palma De Mallorca');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Málaga-Costa Del Sol','AGP','Spain','Málaga');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Gran Canaria','LPA','Spain','Gran Canaria');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Ibiza','IBZ','Spain','Ibiza');

-- Germany
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Frankfurt','FRA','Germany','Frankfurt');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Munich','MUC','Germany','Munich');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Düsseldorf','DUS','Germany','Düsseldorf');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Hamburg','HAM','Germany','Hamburg');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Stuttgart','STR','Germany','Stuttgart');

-- Italy
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Rome Leonardo da Vinci-Fiumicino','FCO','Italy','Rome');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Milan Malpensa','MXP','Italy','Milan');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Venice Marco Polo','VCE','Italy','Venice');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Naples','NAP','Italy','Naples');

-- Portugal
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Lisbon','LIS','Portugal','Lisbon');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Porto','OPO','Portugal','Porto');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Faro','FAO','Portugal','Faro');

-- Ireland
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Dublin','DUB','Ireland','Dublin');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Belfast International','BFS','Ireland','Belfast');

-- Belgium
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Brussels-Zaventem','BRU','Belgium','Brussels');

-- Amsterdam
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Amsterdam Airport Schiphol','AMS','The Netherlands','Amsterdam');

-- Switzerland
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Zurich','ZRH','Switzerland','Zurich');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Geneva','GVA','Switzerland','Geneva');

-- Austria
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Vienna','VIE','Austria','Vienna');

-- Czech Republic
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Prague','PRG','Czech Republic','Prague');

-- Poland
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Warsaw Chopin Airport','WAW','Poland','Warsaw');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Kraków Airport','KRK','Poland','Kraków');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Gdańsk Lech Wałęsa Airport','GDN','Poland','Gdańsk');

-- Denmark
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Copenhagen Airport','CPH','Denmark','Copenhagen');

-- Finland
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Helsinki Airport','HEL','Finland','Helsinki');

-- Iceland
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Keflavík International Airport','KEF','Iceland','Reykjavík');

-- Norway
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Oslo Airport, Gardermoen','OSL','Norway','Oslo');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Bergen Airport Flesland','GBO','Norway','Bergen');

-- Sweden
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Stockholm Arlanda Airport','ARN','Sweden','Stockholm');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Göteborg Landvetter Airport','GOT','Sweden','Gothenburg');

-- Hungary
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Budapest','BUD','Hungary','Budapest');

-- Croatia
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Franjo Tuđman Airport','ZAG','Croatia','Zagreb');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Split Airport','SPU','Croatia','Split');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Dubrovnik Airport','DBV','Croatia','Dubrovnik');

-- Croatia
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Athens','ATH','Greece','Athens');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Heraklion','HER','Greece','Heraklion');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Rhodes','RHO','Greece','Rhodes');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Corfu','CFU','Greece','Corfu');
INSERT INTO DestinationAirports (DestAirportName, DestAirportCode, DestCountry, DestCity) VALUES ('Kos','KGS','Greece','Kos');

SELECT * FROM DestinationAirports;
