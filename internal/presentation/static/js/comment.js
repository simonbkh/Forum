
let comment_bottone = document.querySelectorAll("#chat")
comment_bottone.forEach(chat => {
    chat.addEventListener("click", comment)
})

async function comment(e) {
    let post_id = e.target.getAttribute("data-id")
    let comments_container = document.getElementById(`comments-container-for-${post_id}`)
    let posts = document.getElementById(`${post_id}-container-for-post`)
    let post = document.getElementById(`post-${post_id}`)
    let cmt = document.getElementById(`${post_id}`)
    let cancel_bottone = document.getElementById(`cansel-for-${post_id}`)
    let send = document.getElementById(`send-Comment-for-${post_id}`)
    let comment = document.getElementById(`commantair-for-${post_id}`)
    send.style.display = "none"
    post.style.display = "none"
    comments_container.style.display = "block"
    let data = await getcomment(post_id)
    displaycomment(data)
    function displaycomment(s) {
        let all = cmt.querySelectorAll(".commants")
        all.forEach(ele => ele.remove())
        if (s) {
            s.forEach(element => {
                let div = document.createElement("div")
                div.setAttribute("class", "commants")
                div.innerHTML = `
                        <p><small><i>👤 ${element.Username}</i></small></p>
                        <p><b>${element.Cont}</b></p>
                        <p><small><i>${element.Date}</i></small></p>
                        `
                cmt.append(div)
            });
        } else {
            let div = document.createElement("div")
            div.setAttribute("class", "commants")
            div.innerHTML = `
                        <h4>Not a comment</h4>
                        `
            cmt.append(div)
        }
    }


    comments_container.append(cmt)
    posts.append(comments_container)
    cancel_bottone.addEventListener("click", (e) => {

        comments_container.style.display = "none"
        post.style.display = "block"
    })

    comment.addEventListener("input", () => {

        if (comment.value === "") {
            send.style.display = "none"
        } else {
            send.style.display = "block"
        }
    })

    send.addEventListener("click", async () => {
        let post_id = send.getAttribute("data-post")
        let user_id = send.getAttribute("data-user")
        if (comment.value !== "") {
            try {
                const response = await fetch('http://localhost:8052/newcomment', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ post: `${post_id}`, id: `${user_id}`, comment: comment.value, date: new Date() })
                });

            } catch (error) {
                console.log('Error:', error);
                response.textContent = 'An error occurred while sending the request.';
            }
            let s = await getcomment(post_id)
            comment.value = ""
            send.style.display = "none"
            displaycomment(s)
        }

    })

}


async function getcomment(post_id) {
    try {
        const response = await fetch('http://localhost:8052/getcomment', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ post: `${post_id}` })
        });
        data = await response.json();
    } catch (error) {
        console.log('Error:', error);
        response.textContent = 'An error occurred while sending the request.';
    }
    return data
}

