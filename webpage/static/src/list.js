window.addEventListener("DOMContentLoaded", ()=>{
	list = document.getElementById("list")
	axios({
		method: "GET",
		url: "http://localhost:8080/api/users/"
	}).then((r) => {
		users = r.data
		console.log(users)
		for (user of users) {
			elm = document.createElement("li")
			elm.textContent = user["name"] + ":" + user["password"]
			list.appendChild(elm)
		}
	}).catch(err => {
		alert(err)
	})
})
