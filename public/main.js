'use strict';

//Page elements
var eResponse;
var eUsers;
var eButton;
var eMessage;
//User Name
var name;
//WebSocket
var ws;

var adjectives = [ "Special", "Fun", "Awkward", "Radical", "Super", "Tubular", "Typical", "Smart", "Awesome", "Pink", "Red", "Caustic" ];
var nouns = [ "Pig", "Man", "Women", "Child", "Fish", "Bottle", "Cup", "Tiger", "Student", "Tube", "Foot", "Car", "Teacher", "Professor" ];
//name generator;
function getName() {
  var n1 = adjectives[ Math.floor( Math.random() * adjectives.length ) ];
  var n2 = adjectives[ Math.floor( Math.random() * adjectives.length ) ];
  var n3 = nouns[ Math.floor( Math.random() * nouns.length ) ];

  return n1 + "-" + n2 + "-" + n3;
}

document.addEventListener( "DOMContentLoaded", function( event ) {
  eResponse = document.getElementsByClassName( 'response' )[ 0 ];
  eUsers = document.getElementsByClassName( 'users' )[ 0 ];
  eButton = document.getElementsByClassName( 'my-btn' )[ 0 ];
  eMessage = document.getElementsByClassName( 'message' )[ 0 ];
  name = getName();
  console.log( name );
} );

function websock() {
  var loc = window.location;
  var uri = 'ws:';

  //handle secure connection (TLS/WSS not implemented on server...)
  if ( loc.protocol === 'https:' ) {
    uri = 'wss:';
  }

  uri += '//' + loc.host;
  uri += loc.pathname + 'ws';

  ws = new WebSocket( uri )

  ws.onopen = function() {
    console.log( 'WebSocket Connected' )
    ws.send( "new-user;" + name );
  }

  ws.onclose = function() {
    disconnect();
  }

  ws.onmessage = function( evt ) {
    var out = document.getElementsByClassName( 'message' )[ 0 ];
    out.innerHTML += evt.data + "<br>";
    var msg = evt.data.split( ";" );
    console.log( msg );

    if ( msg[ 0 ] == "new-user" ) {
      var user = document.createElement( "LI" );
      user.className = "user";
      var text = document.createTextNode( msg[ 1 ] );
      user.appendChild( text );
      eUsers.appendChild( user );
    }
    if ( msg[ 0 ] == "disconnect" ) {
      var children = eUsers.children;
      for ( var i = 0; i < children.length; i++ ) {
        var child = children[ i ];
        if ( child.textContent == msg[ 1 ] ) {
          eUsers.removeChild( child );
        }
      }
    }
  }
}
//      <button class="button-disconnect pure-button" onclick="disconnect()">Disconnect</button>

var connect = function() {
  eResponse.textContent = "Connecting..."
  eResponse.className = "response connected";
  eButton.onclick = disconnect;
  eButton.textContent = "Disconnect";
  eButton.className = "my-btn button-disconnect pure-button";
  websock();
}

var disconnect = function() {
  eResponse.textContent = "Disconnecting...";
  eResponse.className = "response disconnected";
  eButton.onclick = connect;
  eButton.textContent = "Connect";
  eButton.className = "my-btn button-connect pure-button"
  var out = document.getElementsByClassName( 'message' )[ 0 ];
  out.innerHTML = "";
  while ( eUsers.hasChildNodes() ) {
    eUsers.removeChild( eUsers.lastChild );
  }
  ws.close();
}


//Get Function
/*
function httpGetAsync( theUrl, callback ) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if ( xmlHttp.readyState == 4 && xmlHttp.status == 200 )
            callback( xmlHttp.responseText );
        if ( xmlHttp.readyState == 4 && xmlHttp.status == 400 )
            callback( xmlHttp.responseText );
    }
    xmlHttp.open( "GET", theUrl, true ); // true for asynchronous
    xmlHttp.send( null );
}
*/
