
let comment_bottone = document.querySelectorAll("#chat")
let len

comment_bottone.forEach(chat => {
    len_comment()
    async function len_comment() {
        let post_id = chat.getAttribute("data-id")
        len = await commentlen(post_id)
        let existingLenElement = chat.parentNode.querySelector('.len');
        if (existingLenElement) {
            existingLenElement.remove();
        }
        let i = document.createElement("sup")
        i.textContent = `${len} 🗨`
        i.setAttribute("class", "len")
        chat.parentNode.append(i)
    }
    chat.addEventListener("click", comment)
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
            clear(".commants")
            clear("i")
            if (s) {
                let i = document.createElement("i")
                i.textContent = ` ${s.length} Comment`
                i.setAttribute("class", "mutch")
                cmt.append(i)
                s.forEach(element => {
                    let div = document.createElement("div")
                    div.setAttribute("class", "commants")
                    let username_italic = document.createElement("i")
                    let small = document.createElement("small")
                    let p = document.createElement("p")
                    username_italic.textContent = `👤 ${element.Username}`
                    small.append(username_italic)
                    p.append(small)
                    div.append(p)
                    let b = document.createElement("strong")
                    let parag_comment = document.createElement("p")
                    b.textContent = element.Cont
                    parag_comment.append(b)
                    div.append(parag_comment)
                    let date_italic = document.createElement("i")
                    let small_date = document.createElement("small")
                    let parag_date = document.createElement("p")
                    date_italic.textContent = element.Date
                    small_date.append(date_italic)
                    parag_date.append(small_date)
                    div.append(parag_date)
                    cmt.append(div)
                });
            } else {
                let i = document.createElement("i")
                i.textContent = ` 0 Comment`
                i.setAttribute("class", "mutch")
                cmt.append(i)
                let div = document.createElement("div")
                div.setAttribute("class", "commants")
                let H = document.createElement("h4")
                H.textContent = `No comments yet`
                div.append(H)
                let h = document.createElement("em")
                h.textContent = ` Be the first to comment ✏️`
                div.append(h)
                cmt.append(div)
            }
        }


        comments_container.append(cmt)
        posts.append(comments_container)
        cancel_bottone.addEventListener("click", (e) => {
            comments_container.style.display = "none"
            post.style.display = "block"
            len_comment()

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
                    const response = await fetch('http://localhost:8080/newcomment', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({ post: `${post_id}`, id: `${user_id}`, comment: comment.value, date: new Date() })
                    });
                        data = await response.json();
                    if (!data.success) {
                        window.location.href = "/login"
                    }
               

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
        function clear(elem) {
            let all = cmt.querySelectorAll(elem)
            all.forEach(ele => ele.remove())

        }
    }
})




async function getcomment(post_id) {
    try {
        const response = await fetch('http://localhost:8080/getcomment', {
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

async function commentlen(postid) {
    try {
        const response = await fetch('http://localhost:8080/commentlen', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ post: `${postid}` })
        });
        data = await response.json();
    } catch (error) {

        console.log('Error:', error);
        response.textContent = 'An error occurred while sending the request.';
    }
    return data
}