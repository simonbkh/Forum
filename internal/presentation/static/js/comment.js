export function func() {
    let bot = document.querySelector("#chat")
    let cancel = document.querySelector("#cansel")
    let comm = document.getElementById("all")
    let comment = document.getElementById("commant")
    let send = document.getElementById("send")
    let tt = document.getElementById("tt")
    comm.style.display = "none"
    bot.addEventListener("click", async (e) => {
        tt.style.display = "none"
        comm.style.display = "inline-block"
        let data;
        try {
            const response = await fetch('http://localhost:8081/getcomment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ post: "{{.User_id}}" })
            });
            data = await response.json();
        } catch (error) {
            console.log('Error:', error);
            response.textContent = 'An error occurred while sending the request.';
        }
        data.forEach(element => {
            let div = document.createElement("div")
            div.setAttribute("class", "commants")
            div.innerHTML = `
                        <p><b>${element.Username}</b></p>
                            <p>${element.Cont}</p>
                            <p><b>${element.Date}</b></p>
                        `
            comm.append(div)
        });
        console.log(data);
    })
    cancel.addEventListener("click", () => {
        comm.style.display = "none"
        tt.style.display = "inline-block"
    })
    send.addEventListener("click", async () => {
        try {
            const response = await fetch('http://localhost:8081/newcomment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ post: "{{.Post_id}}", id: "{{.User_id}}", comment: comment.value, date: "{{.Date}}" })
            });

        } catch (error) {
            console.log('Error:', error);
            response.textContent = 'An error occurred while sending the request.';
        }
    })
}