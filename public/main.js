'using strict'

//Response Element
var eResponse;
//User Name
var name;

var adjectives = ["Special","Fun","Awkward","Radical","Super","Tubular","Typical","Smart","Awesome","Pink","Red","Caustic"];
var nouns = ["Pig","Man","Women","Child","Fish","Bottle","Cup","Tiger","Student","Tube","Foot","Car","Teacher","Professor"];
//name generator;
function getName(){
  var n1 = adjectives[Math.floor(Math.random()*adjectives.length)];
  var n2 = adjectives[Math.floor(Math.random()*adjectives.length)];
  var n3 = nouns[Math.floor(Math.random()*nouns.length)];

  return n1+"-"+n2+"-"+n3;
}

document.addEventListener("DOMContentLoaded", function(event) {
  eResponse = document.getElementsByClassName('response')[0];
  name = getName();
  console.log(name);
});

function connect(){
  httpGetAsync("/connect/"+name,function(response){
    eResponse.textContent=response;
    eResponse.className = "response connected";
});
}

function disconnect(){
  httpGetAsync("/disconnect/"+name,function(response){
    eResponse.textContent=response;
    eResponse.className = "response disconnected";
});
}

//Get Function
function httpGetAsync(theUrl, callback)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            callback(xmlHttp.responseText);
        if(xmlHttp.readyState == 4 && xmlHttp.status == 400)
            callback(xmlHttp.responseText);
    }
    xmlHttp.open("GET", theUrl, true); // true for asynchronous
    xmlHttp.send(null);
}
