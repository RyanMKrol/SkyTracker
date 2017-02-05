<?php

  //globals and includes
  include 'credentials.php';
  include 'httpCodes.php';

?>
<?php
  //used to separate functions from main code

  //this function will be used to setup my source and destination airport arrays
  function arraySetup($conn, $table){

    $result = $conn->query("SELECT * FROM $table;");

    $resultArray = array();

    while($row = mysqli_fetch_array($result)){
      $resultArray[] = $row;
    }

    return $resultArray;
  }

  //function to programatically create my .sql files
  function createMySQLFile($srcAirport, $destAirport) {

    $myfile = fopen("../Data/${srcAirport}_${destAirport}.sql", "w+");

    echo "Creating ../Data/${srcAirport}_${destAirport}.sql\n";

    fwrite($myfile, "DROP TABLE ${srcAirport}_${destAirport};\n\n");
    fwrite($myfile, "CREATE TABLE ${srcAirport}_${destAirport} (\n");
    fwrite($myfile, "\tTripID int NOT NULL AUTO_INCREMENT,\n");
    fwrite($myfile, "\tSourcePort varchar(255) NOT NULL,\n");
    fwrite($myfile, "\tDestPort varchar(255) NOT NULL,\n");
    fwrite($myfile, "\tDepartDate DATE NOT NULL,\n");
    fwrite($myfile, "\tReturnDate DATE NOT NULL,\n");
    fwrite($myfile, "\tPrice int NOT NULL,\n");
    fwrite($myfile, "\tPRIMARY KEY(TripID),\n");
    fwrite($myfile, "\tCONSTRAINT uc_date_pair UNIQUE (DepartDate, ReturnDate)\n");
    fwrite($myfile, ");\n\n");

    return $myfile;

  }

  //function to pull data from SkyScanner's API
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

    //there are several possible HTTP response code, I'll be using the following
    switch ($http_code = curl_getinfo($curl, CURLINFO_HTTP_CODE)) {

      case $httpSuccess:  # All's fine

        $data = json_decode($curl_response,true);
        return $data;
        break;

      case $httpExcess:  # Using too much
        sleep(1);
        echo '*********************************************************************** TOO MANY CALLS, LETS WAIT AND TRY AGAIN';
        return getData($srcAirport, $destAirport, $departYear, $departMonth, $returnYear, $returnMonth);
        break;

      default:
        echo '************************************************************************Unexpected HTTP code: ', $http_code, "\n";
        echo "$call\n";
    }
  }

  //function responsible for parsing data and writing valid sql commands to the right file
  function writeData($data, $file, $srcAirport, $destAirport, $departYear, $departMonth, $returnYear, $returnMonth){

    $outboundDay = 0;
    $inboundDay = 0;

    foreach($data["Dates"] as $key => $val) { //foreach element in $arr

      foreach($val as $inKey => $inVal) { //foreach element in $arr
        if(isset($inVal['MinPrice'])){

          $minPrice = $inVal['MinPrice'];

          fwrite($file, "INSERT INTO ${srcAirport}_${destAirport} (SourcePort, DestPort, DepartDate, ReturnDate, Price) VALUES ('$srcAirport', '$destAirport', '$departYear-$departMonth-$outboundDay', '$returnYear-$returnMonth-$inboundDay', $minPrice);\n");
        }

        $outboundDay++;
      }

      $inboundDay++;
      $outboundDay = 0;

    }
  }

?>
<?php

  //main

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $sourcesArr = arraySetup($conn, 'SourceAirports');
  $destinationsArr = arraySetup($conn, 'DestinationAirports');

  //creates date from the current date
  $srcDate = date_create();
  $destDate = date_create();

  $monthsNeeded = 2;

  if($srcDate->format('d') >= 14){
    $monthsNeeded = 3;
  }

  foreach($sourcesArr as $srcAirport) { //foreach element in $arr
    foreach($destinationsArr as $destAirport) { //foreach element in $arr

      $src = $srcAirport["SrcAirportCode"];
      $dest = $destAirport["DestAirportCode"];

      //found in functions.php
      $mysqlFile = createMySQLFile($src, $dest);

      //we can depart within the next 6 months
      foreach(range(1,6) as $j){

        //and we can come back within #monthsNeeded months of the set-off date
        foreach(range(1,$monthsNeeded) as $i){

          $departYear = $srcDate->format('Y');
          $departMonth = $srcDate->format('m');
          $returnYear = $destDate->format('Y');
          $returnMonth = $destDate->format('m');

          $data = getData($src, $dest, $departYear, $departMonth, $returnYear, $returnMonth);

          if(!is_null($data)){

              //write the data to the sql file
              writeData($data, $mysqlFile, $src, $dest, $departYear, $departMonth, $returnYear, $returnMonth);

          } else {
              //have to decide on functionality for this later
          }

          date_add($destDate,date_interval_create_from_date_string("1 month"));

        }
        date_add($srcDate,date_interval_create_from_date_string("1 month"));

        $destDate = date_create();
        date_add($destDate,date_interval_create_from_date_string("$j months"));
      }

      $srcDate = date_create();
      $destDate = date_create();

      //update the database
      exec("mysql -u root -p\"$password\" -f \"SkyTracker\" < ${src}_${dest}.sql");

      //close file now we're done with the src-dest pair
      fclose($mysqlFile);
    }
  }
?>
