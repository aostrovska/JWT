async function request(url) {
	var body = {
			"login" : document.getElementById("login").value,
			"pass" : document.getElementById("password").value
		}
		let res = await fetch(url, {
		method: 'POST', // *GET, POST, PUT, DELETE, etc.
		mode: 'cors', // no-cors, *cors, same-origin
		cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
		credentials: 'same-origin', // include, *same-origin, omit
		headers: {
		  'Content-Type': 'application/json'
		  // 'Content-Type': 'application/x-www-form-urlencoded',
		},
		redirect: 'follow', // manual, *follow, error
		referrerPolicy: 'no-referrer', // no-referrer, *client
		body: JSON.stringify(body)// body data type must match "Content-Type" header
	  });
	  let inf = await res;
		console.log(inf);
	
	let jso = await res.text();
		document.getElementById("resp").innerHTML = jso;

	
}

function log() {
		request("http://localhost:8080/");
}

function dat() {
	//	request(document.getElementById("url").value, document.getElementById("body").value);
}
