// init
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
			outStr += "<td id='button" + x + "_" + y +"'  readonly onclick='navButtonPressed(" + x + " , " + y + ")'>?</td>";
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
	sendPost("navbutton", ["xpos", "ypos"], [x, y]);

}

// set the board to reflect the player's existing position
function positionPlayer()
{
	sendPost("navloaded", null, null);
	sendPost("getOtherPlayers", null, null);
}

// send a post request to the server
function sendPost(actionType, keys, values)
{
	// clear the message board
	setNavMessage("");
	// if given a single key and value, make them lists
	if(Array.isArray(keys) == false)
	{
		keys = [keys];
		values = [values];
	}
	// make a querystring
	var outStr = "?actiontype=" + actionType;
	var i = 0;
	for(i = 0; i < keys.length; i++)
	{
		outStr += "&" + keys[i] + "=" + values[i];
	}
	// send it
	var xhttp = new XMLHttpRequest();
	xhttp.open("POST", "https://tradewars-se201.herokuapp.com/" + outStr, true)
	xhttp.onreadystatechange = function(e) {
		if (xhttp.readyState === 4) {
			if (xhttp.status === 200) {
				eval(xhttp.response);
			}
		}
	}
	xhttp.send();  
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

// set the message in the nav page message board
function setNavMessage(str) {
	document.getElementById("navMessageBoard").innerHTML = str;
}

// update which tile shows the player's location
function updatePlayerLoc(xLoc, yLoc) {
	var x = 0;
	var y = 0;
	for(y = 0; y < 10; y++)
	{
		for(x = 0; x < 10; x++)
		{
			var elementStr = "button" + x + "_" + y;
			if(xLoc === x && yLoc === y) {
				document.getElementById(elementStr).innerHTML = "X";
			}
			else {
				document.getElementById(elementStr).innerHTML = ".";
			}
			
		}
	}
}

function updateOtherPlayersLoc(xloc, yLoc){
	console.log('you got to update other players')
	for(y = 0; y < 10; y++){
		for(x = 0; x <10; x++){
			if (x == xLoc && y == yLoc){
				var elementStr = "button" + x + "_" + y;
				document.getElementById(elementStr).innerHTML = "Y";
			}
		}
	}
}

const ws = new WebSocket('ws://' + window.location.host + '/ws');
ws.onopen = function open() {
	ws.send('Works!');
}

ws.onopen = function incoming(data) {
	console.log(data);
}