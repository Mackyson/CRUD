window.addEventListener("DOMContentLoaded", ()=>{
	list = document.getElementById("list")
	axios({
		method: "GET",
		url: "http://localhost:8080/api/users/"
	}).then((r) => {
		users = r.data
		for (user of users) {
			var line = document.createElement("tr")

			var name = document.createElement("td")
			var password = document.createElement("td")
			var edit = document.createElement("td")
			var destroy = document.createElement("td")

			name.innerHTML = user["name"]
			password.innerHTML = user["password"]
			edit.innerHTML = "<button>Edit</button>"
			edit.addEventListener("click",updateUser(user["name"]))
			destroy.innerHTML = "<button>Delete</button>"
			destroy.addEventListener("click",deleteUser(user["name"]))

			line.appendChild(name)
			line.appendChild(password)
			line.appendChild(edit)
			line.appendChild(destroy)

			list.appendChild(line)

		}
	}).catch(err => {
		alert(err)
	})
})

function updateUser(name){
	return function (){
		axios({
			method: "PUT",
			url: "http://localhost:8080/api/change/" + name,
			data: {
				name: name,
			}
		}).then(response => {
			alert(response.data)
		}).catch(err => {
			alert(err)
		})
	}
}

function deleteUser(name){
	return function (){
		axios({
			method: "DELETE",
			url: "http://localhost:8080/api/delete/" + name,
		}).then(response => {
			alert(response.data)
			location.reload()
		}).catch(err => {
			alert(err)
		})
	}
}
