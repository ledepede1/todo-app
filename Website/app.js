async function GetTodoList() {
    const response = await fetch("http://localhost:8080/backend/getlist", {
        method: "GET",
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    });

    if (response.ok) {
        const iconUnchecked = "glyphicon glyphicon-unchecked";
        const iconChecked = "glyphicon glyphicon-check";
        const data = await response.json();

        data.forEach((element, index) => {
            console.log(element.id)
            const icon = element.checked ? iconChecked : iconUnchecked;
            console.log(icon)

            const listItem = document.createElement("div");
            listItem.innerHTML = `
            <div class="container" id="container${index}">
                <div class="wrapper">
                    <div class="item">
                        <p>${element.text}</p> 
                        <button onClick="{SaveTask(${element.id}, ${index})}">
                            <i id="icon${index}" class="${icon}"></i>
                        </button>
                        <button onClick="{DeleteTask(${element.id}, ${index})}">
                            <i id="remove${index}" class="glyphicon glyphicon-remove-circle"></i>
                        </button>
                    </div>
                </div>
            <div/>
            `;
            todoListContainer.appendChild(listItem);
        });
    }
}

async function CreateTask() {
    text = prompt("What should the task be?", "");
    const response = await fetch("http://localhost:8080/backend/addtask", {
        method: "POST",
        body: JSON.stringify({
            text: text
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    });
    location.reload();
}

async function DeleteTask(id, index) {
    const container = document.getElementById(`container${index}`);

    if (container) {
        container.remove();

        const response = await fetch("http://localhost:8080/backend/deletetask", {
        method: "POST",
        body: JSON.stringify({
            id: id
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    });
    }
}

async function SaveTask(id, index) {
    console.log(id)
    const response = await fetch("http://localhost:8080/backend/savetask", {
        method: "POST",
        body: JSON.stringify({
            id: id
        }),
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    });


    if (response.ok) {
        toggleIcon(index)
    }
}

function toggleIcon(index) {
    const iconElement = document.getElementById("icon" + index);

    if (iconElement) {
        const isChecked = iconElement.classList.contains("glyphicon-check");

        if (isChecked) {
            iconElement.classList.remove("glyphicon-check");
            iconElement.classList.add("glyphicon-unchecked");
        } else {
            iconElement.classList.remove("glyphicon-unchecked");
            iconElement.classList.add("glyphicon-check");
        }
    }
}