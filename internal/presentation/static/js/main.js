document.querySelectorAll('.post').forEach(post => {
    const postId = post.getAttribute('post-id');
    const ratings = post.querySelectorAll('.reaction');
    // console.log(rating.length)  
    const likeReaction = ratings[0];
    ratings.forEach(rating => {
        const button = rating.querySelector('.post-reaction');
        const count = rating.querySelector('.reaction-count');
        button.addEventListener("click", async () => {
            const likeOrDislike = likeReaction === rating ? "like" : "dislike";
            const response = await fetch(`/api/posts/${postId}/${likeOrDislike}`);
            const body = await response.json(); 
            console.log(body)
        });

    });
});


// function react() {
//     const rating = document.querySelector('.reaction');
//     const likeReaction = rating
//     const like =likeReaction.querySelector('.reaction-count')
//     like.innerText= parseInt(like. innerText)+1
//     rating.className="reaction selected"
//     console.log(rating)
// }