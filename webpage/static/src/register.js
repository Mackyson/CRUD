// import axios from "axios"

function createUser(){
	name = document.getElementById("name").value
	password = document.getElementById("password").value
	axios({
		method: "POST",
		url: "http://localhost:8080/api/signup",
		data: {
			name: name,
			password: password,
		}
	}).then(response => {
		alert(response.data)
	}).catch(err => {
		alert(err)
	})
}

window.addEventListener("DOMContentLoaded", ()=>{
	// document.getElementById("resign").addEventListener("click", deleteUser)
	document.getElementById("register").addEventListener("click", createUser)
})
