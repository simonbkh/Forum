
let comment_bottone = document.querySelectorAll("#chat")
let n = 0
let err = document.createElement("p")
comment_bottone.forEach(chat => {
    len_comment()
    async function len_comment() {
        let post_id = chat.getAttribute("data-id")
        commentlen(post_id)
        function postlendisplay(len) {
            let existingLenElement = chat.parentNode.querySelector('.len');
            if (existingLenElement) {
                existingLenElement.remove();
            }
            let i = document.createElement("p")
            i.textContent = `${len}`
            i.setAttribute("class", "len")
            chat.parentNode.append(i)
        }
        async function commentlen(postid) {
            fetch('/commentlen', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ post: `${postid}` })
            })
                .then(async response => {
                    if (!response.ok) {
                        return response.text().then(text => {
                            try {
                                const jsonResponse = JSON.parse(text);
                                console.log(jsonResponse);
                            } catch (e) {
                                document.body.innerHTML = text;
                            }
                        });
                    } else {
                        return response.json();
                    }
                })
                .then(data => {
                    postlendisplay(data)
                })
                .catch(error => {
                    console.log('Error:', error);
                })

        }

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
        let err = document.getElementById(`error-for-${post_id}`)
        if (comment.value.length < 1) {
            send.style.display = "none"
        }
        post.style.display = "none"
        comments_container.style.display = "block"
        getcomment(post_id)
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
                    let reaction = document.createElement("div")
                    reaction.setAttribute('id', `${element.ID}-comment-reaction`)
                    reaction.innerHTML = `<div class="comment-reaction">
                            <span class="cmt-reaction material-icons reacted" >thumb_up</span>
                            <span class="reaction-count">${element.Reactions[0]}</span>
                        </div>
                        <div class="comment-reaction">
                            <span class="cmt-reaction material-icons reacted">thumb_down</span>
                            <span class="reaction-count">${element.Reactions[1]}</span>
                        </div>`

                    div.setAttribute("class", "commants")
                    let username_italic = document.createElement("i")
                    let small = document.createElement("small")
                    let p = document.createElement("p")
                    username_italic.textContent = `👤 ${element.Username}`
                    small.append(username_italic)
                    p.append(small)
                    div.append(p)
                    let b = document.createElement("strong")
                    let parag_comment = document.createElement("pre")
                    parag_comment.style.lineHeight = "1.5"
                    parag_comment.style.whiteSpace = "pre-line"
                    b.textContent = element.Cont
                    parag_comment.append(b)
                    div.append(parag_comment)
                    let date_italic = document.createElement("i")
                    let small_date = document.createElement("small")
                    let parag_date = document.createElement("p")
                    date_italic.textContent = Dateformat(element.Date)
                    small_date.append(date_italic)
                    parag_date.append(small_date)
                    div.append(parag_date)
                    div.append(reaction)
                    cmt.append(div)
                    const reactions = div.querySelectorAll('.cmt-reaction')

                    if (element.State == 1) {
                        reactions[0].setAttribute("style", 'color: #2196F3')
                    } else if (element.State == 2) {
                        reactions[1].setAttribute("style", 'color: #2196F3')
                    }
                    reactions[0].addEventListener('click', () => toggleReaction('comments', element.ID, 'like'));
                    reactions[1].addEventListener('click', () => toggleReaction('comments', element.ID, 'dislike'));
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
        let temp
        comment.addEventListener("input", () => {
           temp = comment.value.replaceAll(" ", "")
            if (comment.value === "" || temp === "" || temp.replaceAll("\n", "") === "") {
                err.style.display = "none"
                send.style.display = "none"
            } else if (comment.value.length > 1000) {
                err.style.display = "block"
                err.textContent = `${1000 - comment.value.length}`
                send.style.display = "none"
            } else {
                err.style.display = "none"
                send.style.display = "block"
            }
        })

        send.addEventListener("click", () => {
            let post_id = send.getAttribute("data-post")
            if (comment.value !== "" && n === 0) {
                n++
                fetch('/newcomment', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ post: `${post_id}`, comment: comment.value, date: new Date() })
                })
                    .then(async response => {

                        if (!response.ok) {
                            return response.text().then(text => {
                                try {
                                    const jsonResponse = JSON.parse(text);
                                    if (jsonResponse && !jsonResponse.success) {

                                        window.location.href = "/login"
                                        n = 0
                                    }
                                } catch (e) {
                                    document.body.innerHTML = text;
                                }
                            });
                        } else {
                            return response.json();
                        }
                    })
                    .then(data => {

                        if (data && data.success) {
                            comment.value = ""
                            send.style.display = "none"
                            n = 0
                            getcomment(post_id)
                        }
                    })

                    .catch(error => {

                        console.log(error);
                    })
            }

        })
        function clear(elem) {
            let all = cmt.querySelectorAll(elem)
            all.forEach(ele => ele.remove())

        }


        async function getcomment(post_id) {
            fetch('/getcomment', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ post: `${post_id}` })
            })
                .then(async response => {
                    if (!response.ok) {
                        return response.text().then(text => {
                            try {
                                const jsonResponse = JSON.parse(text);
                                console.log(jsonResponse);
                            } catch (e) {
                                document.body.innerHTML = text;
                            }
                        });
                    } else {
                        return response.json();
                    }
                })
                .then(data => {
                    displaycomment(data)
                })
                .catch(error => {
                    console.log('Error:', error);
                })

        }
    }
})


const Dateformat = (timestamp) => {
    let pastDate = new Date(timestamp);
    pastDate.setHours(pastDate.getHours() - 1)
    let now = new Date();
    pastDate.setHours(pastDate.getHours() + 1)
    let seconds = Math.floor((now - pastDate) / 1000);
    if (seconds < 60) {
        return `${seconds} seconds`;
    } else if (seconds < 3600) {
        return `${Math.floor((seconds / 60))} minutes`;
    } else if (seconds < 86400) {
        return `${Math.floor(seconds / 3600)} heures`;
    } else {
        return `${Math.floor(seconds / 86400)} jours`;
    }
}