console.log(document.cookie);

let likes = {}
let dislikes = {}

function toggleReaction(postId, type) {  
      
    const post = document.getElementById(`${postId}-container-for-post`)
    if (!post) return

    const reactions = post.querySelectorAll('.reaction')
    const likeReaction = reactions[0]
    const dislikeReaction = reactions[1]

    const likeIcon = likeReaction.querySelector('.post-reaction')
    const dislikeIcon = dislikeReaction.querySelector('.post-reaction')
    const likeCount = likeReaction.querySelector('.reaction-count')
    const dislikeCount = dislikeReaction.querySelector('.reaction-count')

    console.log(likes, dislikes);
    

    // Send to server using single endpoint
    fetch(`/api/posts/${postId}/reaction`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            type: type,
            action: type === 'like' ? (!likes[postId] ? 'add' : 'remove') : (!dislikes[postId] ? 'add' : 'remove'),
            SessionToken: document.cookie
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            console.log('Reaction added successfully');
            if (type === 'like') {
                if (likes[postId]) {
                    delete likes[postId]
                    likeIcon.style.color = ''
                    likeCount.textContent = parseInt(likeCount.textContent) - 1
                } else {
                    likes[postId] = true
                    likeIcon.style.color = '#2196F3'
                    likeCount.textContent = parseInt(likeCount.textContent) + 1
        
        
                    if (dislikes[postId]) {
                        delete dislikes[postId]
                        dislikeIcon.style.color = ''
                        dislikeCount.textContent = parseInt(dislikeCount.textContent) - 1
                    }
                }
            } else if (type === 'dislike') {
                if (dislikes[postId]) {
                    delete dislikes[postId]
                    dislikeIcon.style.color = ''
                    dislikeCount.textContent = parseInt(dislikeCount.textContent) - 1
                } else {
                    dislikes[postId] = true
                    dislikeIcon.style.color = '#F44336'
                    dislikeCount.textContent = parseInt(dislikeCount.textContent) + 1
        
                    if (likes[postId]) {
                        delete likes[postId]
                        likeIcon.style.color = ''
                        likeCount.textContent = parseInt(likeCount.textContent) - 1
                    }
                }
            }
        }else{
            window.location.href = '/login'
            return
        }
    })
    .catch(error => console.error('Error:', error))

    
}

// Add click event listeners to all reaction buttons
document.addEventListener('DOMContentLoaded', () => {
    const posts = document.querySelectorAll('.likes-section')
    
    
    posts.forEach(post => {
        const postId = post.id
        const reactions = post.querySelectorAll('.reaction')

        // Like button
        reactions[0].addEventListener('click', () => toggleReaction(postId, 'like'))

        // Dislike button
        reactions[1].addEventListener('click', () => toggleReaction(postId, 'dislike'))
    })
})