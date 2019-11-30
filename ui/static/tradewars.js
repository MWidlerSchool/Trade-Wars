// init
var xhttp = new XMLHttpRequest();
addButtons();

// hide inactive forms, show the active one
function setActiveForm(targetString)
{
	// hide everything
	$(".divPanel").hide();
	// show the thing with the matching ID
	$("#" + targetString).show();
}

// function to set the hundred buttons 
function addButtons()
{
	var x = 0;
	var y = 0;
	var outStr = "";
	for(y = 0; y < 10; y++)
	{
		outStr += "<tr>";
		for(x = 0; x < 10; x++)
		{
			outStr += "<td  readonly onclick='navButtonPressed(" + x + " , " + y + ")'>?</td>";
		}
		outStr += "</tr>";
	}
	document.getElementById("sectorBoard").innerHTML = outStr;
}

// read in player name from name page
function readInName()
{
	var playerName = document.getElementById("name-input").value;
	// test line; replace with actual code later
	console.log(playerName);
	document.cookie = "Callsign="+playerName
}

// receives presses from nav buttons
function navButtonPressed(x, y)
{
	console.log(x + ", " + y);
	var locStr = "actiontype=navbutton&xpos=" + x + "&ypos=" + y;
	xhttp.open("POST", "https://tradewars-se201.herokuapp.com/", true)
	xhttp.send(locStr); 
}

//hides main screen header when game starts
function hideMainHeader()
{
	var mainHeader = document.getElementById("startpage-header");
	
	mainHeader.style.display = "none"
}

function getCallsign() {
	var callHeader = document.getElementById("callsign");
	var callsign = document.cookie.match(Callsign)

	callHeader.innerHTML = callsign
}