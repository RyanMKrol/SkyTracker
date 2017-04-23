$( document ).ready(function() {

  var query = location.search.substr(1);
  var result = {};

  query.split("&").forEach(function(part) {
    var item = part.split("=");

    switch(item[0]){
      case 'email' :
        $("#EmailAddress").val(item[1]);
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
    }
  });

});

$( "form" ).on( "submit", function( event ) {

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

  jsonRaw["emailAddress"] = $("#EmailAddress").val()
  jsonRaw["budget"] = $( "#budgetSlider" ).slider( "value" );
  jsonRaw["tripMinLen"] = $( "#tripLengthSlider" ).slider( "values", 0 );
  jsonRaw["tripMaxLen"] = $( "#tripLengthSlider" ).slider( "values", 1 );
  jsonRaw["months"] = monthArray;
  jsonRaw["airports"] = airportArray;
  jsonRaw["salt"] = getUrlParam("token");

  var jsonData = JSON.stringify(jsonRaw);

  $.ajax({
    type: "POST",
    url: "./../Scripts/addUser.php",
    data: {_data: jsonData},
    success: function(data){
      alert("things were done right");
      console.log(data);
    }
  });

  return false;
});

// token extraction found at: https://www.sitepoint.com/url-parameters-jquery/
//  modified a bit for if the token isn't in the url
function getUrlParam(name){
	var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
  if (results != undefined) {
    console.log(results);
    return results[1];
  } else {
    return 0;
  }
}
