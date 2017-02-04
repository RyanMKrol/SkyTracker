<?php

//fucntion to create a .sql file for the table we want to store the data in
function createMySQLFile($srcAirport, $destAirport) {

  $myfile = fopen("../Data/${srcAirport}_${destAirport}.sql", "w+");

  fwrite($myfile, "DROP TABLE ${srcAirport}_${destAirport};\n\n");
  fwrite($myfile, "CREATE TABLE ${srcAirport}_${destAirport} (\n");
  fwrite($myfile, "\tTripID int NOT NULL AUTO_INCREMENT,\n");
  fwrite($myfile, "\tSourcePort varchar(255) NOT NULL,\n");
  fwrite($myfile, "\testPort varchar(255) NOT NULL,\n");
  fwrite($myfile, "\tDepartDate DATE NOT NULL UNIQUE,\n");
  fwrite($myfile, "\tReturnDate DATE NOT NULL UNIQUE,\n");
  fwrite($myfile, "\tPrice int NOT NULL,\n");
  fwrite($myfile, "\tPRIMARY KEY(TripID)\n");
  fwrite($myfile, ");\n");

  return $myfile;

}

function getData($srcAirport, $destAirport, $departYear, $departMonth, $returnYear, $returnMonth){

  global $apikey;
  global $httpSuccess;
  global $httpExcess;

  $call = "http://partners.api.skyscanner.net/apiservices/browsegrid/v1.0/GB/GBP/en-GB/$srcAirport/$destAirport/$departYear-$departMonth/$returnYear-$returnMonth?apiKey=$apikey";

  // initialist the api request
  $curl = curl_init($call);

  // returns the api request as a string
  curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

  // execute the api request
  $curl_response = curl_exec($curl);

  echo $call . "\n";

  //i think that we're going to build the table files in the thread, and then use a shell command to execute it
  //mysql -D"SkyTracker" -p"$password" < testSCRIPT.sql

  switch ($http_code = curl_getinfo($curl, CURLINFO_HTTP_CODE)) {

    case $httpSuccess:  # All's fine

      $data = json_decode($curl_response,true);
      return $data
      break;

    case $httpExcess:  # Using too much
      sleep(1);
      return getData($srcAirport, $destAirport, $departYear, $departMonth, $returnYear, $returnMonth);
      break;
      
    default:
      echo 'Unexpected HTTP code: ', $http_code, "\n";
  }
}

/*

$outboundDay = 0;
$inboundDay = 0;

foreach($data["Dates"] as $key => $val) { //foreach element in $arr

  foreach($val as $inKey => $inVal) { //foreach element in $arr
    if(isset($inVal['MinPrice'])){

      //this isn't right just yet, fix this.
      echo "flying out on day $outboundDay, and coming back on $inboundDay\n";

      echo $inVal['MinPrice'] . "\n";
      echo $inVal['QuoteDateTime'] . "\n";
    }

    $outboundDay++;
  }

  $inboundDay++;
  $outboundDay = 0;

}
echo "all good\n";

*/

?>
