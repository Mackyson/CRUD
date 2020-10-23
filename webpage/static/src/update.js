function updateUser(){
	axios({
		method: "PUT",
		url: "http://localhost:8080/api/change/" + name,
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
	document.getElementById("resign").addEventListener("click", updateUser)
	// document.getElementById("register").addEventListener("click", createUser)
})
