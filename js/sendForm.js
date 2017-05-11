var ipAddress;
var captchaToken = undefined;

$( document ).ready(function() {

  var email;
  var token;
  var unsub = false;

  var query = location.search.substr(1);
  var result = {};

  query.split("&").forEach(function(part) {
    var item = part.split("=");

    switch(item[0]){
      case 'email' :
        $("#EmailAddress").val(item[1]);
        email = item[1];
        break;
      case 'source':
      case 'month' :
        $("input:checkbox[value=" + item[1] + "]").attr("checked", true);
        break;
      case 'tripMin' :
        $( "#tripLengthSlider" ).slider( "values", 0, item[1]);
        $( "#daysAmount" ).val(item[1] + " - " + $( "#tripLengthSlider" ).slider( "values", 1) + " Days");
        break;
      case 'tripMax' :
        $( "#tripLengthSlider" ).slider( "values", 1, item[1]);
        $( "#daysAmount" ).val($( "#tripLengthSlider" ).slider( "values", 0) + " - " + item[ 1 ] + " Days");
        break;
      case 'budget' :
        $( "#budgetSlider" ).slider( "option", "value", item[1]);
        $( "#budgetAmount" ).val( "£" + item[1] );
        break;
      case 'frequency' :
        $("input:radio[value=" + item[1] + "]").attr("checked", true);
        break;
      case 'token' :
        token = item[1];
        break;
      case 'unsubscribe' :
        unsub = true;
        break;
    }
  });

  if(unsub) {

    var jsonRaw = {};

    jsonRaw["token"] = token;
    jsonRaw["email"] = email;

    var jsonData = JSON.stringify(jsonRaw);

    console.log(jsonData);

    $.ajax({
      type: "POST",
      url: "./../Scripts/unsubscribe.php",
      data: {_data: jsonData},
      success: function(data){
        data = JSON.parse(data);
        alert(data['message']);
      },
    });
  }

});

$( "form" ).on( "submit", function( event ) {

  if(captchaToken == undefined){
    alert("Please fill out the Captcha before submitting");
    return false;
  }

  if($("input[name='SourceAirport']:checked").length == 0){
    alert("You must select at least one source airport!");
    return false;
  }

  if($("input[name='Month']:checked").length == 0){
    alert("You must select at least one month to track!");
    return false;
  }

  // now we've got an array for what's checked and what isn't
  var monthArray = {};
  $("input[name='Month']").each(function(){
    monthArray[$(this).val()] = $(this).is(":checked");
  });

  // now we've got an array for what's checked and what isn't
  var airportArray = {};
  $("input[name='SourceAirport']").each(function(){
    airportArray[$(this).val()] = $(this).is(":checked");
  });

  var jsonRaw = {};

  jsonRaw["emailAddress"] = $("#EmailAddress").val();
  jsonRaw["budget"] = $( "#budgetSlider" ).slider( "value" );
  jsonRaw["tripMinLen"] = $( "#tripLengthSlider" ).slider( "values", 0 );
  jsonRaw["tripMaxLen"] = $( "#tripLengthSlider" ).slider( "values", 1 );
  jsonRaw["months"] = monthArray;
  jsonRaw["airports"] = airportArray;
  jsonRaw["salt"] = getUrlParam("token");
  jsonRaw["frequency"] = $("input[name='Frequency']:checked").val();
  jsonRaw["captcha"] = captchaToken;

  var jsonData = JSON.stringify(jsonRaw);

  $.ajax({
    type: "POST",
    url: "./../Scripts/addUser.php",
    data: {_data: jsonData},
    success: function(data){
      data = JSON.parse(data);
      console.log(data);
      grecaptcha.reset();
      alert(data['message']);
      if(!data['success']){
        return false;
      }
    },
    error: function(response){
      alert("Sorry, something went wrong in the database.\n Try again later!");
    }
  });

  return false;
});

// token extraction found at: https://www.sitepoint.com/url-parameters-jquery/
//  modified a bit for if the token isn't in the url
function getUrlParam(name){
	var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
  if (results != undefined) {
    return results[1];
  } else {
    return 0;
  }
}

function captchaSet(token) {
  captchaToken = token;
}
