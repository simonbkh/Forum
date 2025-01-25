
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
    post.style.display = "none"
    comments_container.style.display = "inline-block"
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
                        <p><b>${element.Username}</b></p>
                        <p>${element.Cont}</p>
                        <p><b>${element.Date}</b></p>
                        `
                cmt.append(div)
            });
        } else {
            let div = document.createElement("div")
            div.setAttribute("class", "commants")
            div.innerHTML = `
                        <h3>Not a comment</h3>
                        `
            cmt.append(div)
        }
    }


    comments_container.append(cmt)
    posts.append(comments_container)
    cancel_bottone.addEventListener("click", (e) => {

        comments_container.style.display = "none"
        post.style.display = "inline-block"
    })
    send.addEventListener("click", async () => {
        let post_id = send.getAttribute("data-post")
        let user_id = send.getAttribute("data-user")

        try {
            const response = await fetch('http://localhost:8081/newcomment', {
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
        console.log("====Z", comment.value);
        comment.value = ""
        displaycomment(s)
    })

}


async function getcomment(post_id) {
    try {
        const response = await fetch('http://localhost:8081/getcomment', {
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

