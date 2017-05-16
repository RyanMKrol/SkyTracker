<?php
  //globals and includes
  include './../Go/src/Credentials/credentials.php';

?>
<?php
  //main

  // Create connection
  $conn = new mysqli($server, $user, $password, $database);

  // Check connection
  if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
  }

  $emailAddress = $argv[1];

  $title = "Welcome to SkyTracker!";

  $sql = "SELECT * FROM Users WHERE UserEmailAddress = '$emailAddress';";
  $result = mysqli_query($conn, $sql);
  $result = mysqli_fetch_assoc($result);

  $salt = $result['UserSalt'];

  $emailBody = "";
  $emailBody .= "<!doctype html>";
  $emailBody .= "<html lang=\"en\">";
  $emailBody .= "<head>";
  $emailBody .= "<meta charset=\"UTF-8\">";
  $emailBody .= "</head>";
  $emailBody .= "<body style = \"font-family: Georgia; background-color: #eee; color:#111111\">";
  $emailBody .= "<br>";
  $emailBody .= "<div style = \"width: 700px; margin:auto;\">";
  $emailBody .= "<h3>Welcome!</h3>";
  $emailBody .= "<p>It appears that you have signed up to receive flight price reports from SkyTracker, if this is not you, then please feel free to <a style = \"color:#5C596B;font-weight: normal;\" href = \"http://www.skytracker.co?unsubscribe&email=$emailAddress&token=$salt\">unsubscribe</a> ";
  $emailBody .= "or check us out at <a style = \"color:#5C596B;font-weight: normal;\" href = 'http://www.skytracker.co'>SkyTracker</a> and decide for yourself if you want to join! <br><br>If you <b>did</b> mean to sign up, please do us a favour and add us as a contact so we can get our reports straight to your inbox!<br><br>Thanks!</p>";
  $emailBody .= "</div>";
  $emailBody .= "<div style = \"width:700px; margin: auto;\">- the SkyTracker team.</div>";
  $emailBody .= "<br><div style = \"width: 700px; margin: auto;\">";
  $emailBody .= "<table style = \"width: 600px; margin: auto; border-bottom-style: solid; border-bottom-color: white; border-bottom-width: 2px;\">";
  $emailBody .= "</table><br>";
  $emailBody .= "<a href = \"https://github.com/RyanMKrol/SkyTracker\"><img style = \"display: block; margin: 0 auto;\" src = \"http://skytracker.co/Images/GitHub-Mark-32px.png\"></img></a><br>";
  $emailBody .= "<div style = \"font-size: 9pt; width: 500px; margin: auto;text-align:Center;\">";
  $emailBody .= "Our mailing address is: <br><a style = \"color:#5C596B;font-weight: normal;\" href = \"mailto:root@skytracker.co\">root@skytracker.co</a><br><br>";
  $emailBody .= "</div></div>";
  $emailBody .= "</body>";
  $emailBody .= "</html>";

  echo $emailBody . "\n";
  file_put_contents("temp.txt", $emailBody);

  exec("cat temp.txt | mail -a \"From: SkyTracker <no-reply@skytracker.co>\" -a \"Content-type: text/html\" -s \"$title\" $emailAddress");

  exec("rm temp.txt");

  $conn->close();

?>
