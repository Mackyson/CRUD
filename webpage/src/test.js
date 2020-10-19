// import axios from "axios"

function deleteUser(){
	name = document.getElementById("name").value
	// password = document.getElementById("password")
	axios({
		method: "DELETE",
		url: "http://localhost:8080/delete/" + name,
	}).then(response => {
		alert(response.data)
	}).catch(err => {
		alert(err)
	})
}


window.addEventListener("DOMContentLoaded", ()=>{
	document.getElementById("resign").addEventListener("click", deleteUser)
})
