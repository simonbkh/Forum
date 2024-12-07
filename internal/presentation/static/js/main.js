// document.getElementsByClassName('post').forEach(post => {
//     const postId = post.getAttribute('post-id');
//     const rating = post.querySelectorAll('.reaction');
//     console.log(rating.length)
//     likeReaction = rating[0];
//     rating.forEach(rating => {
//         const button = rating.querySelector('.post-reaction');
//         const count = rating.querySelector('.reaction-count');
//         console.log(button);
//         console.log(count);
//     });
// });


function react() {
    const rating = document.querySelector('.reaction');
    const likeReaction = rating
    const like =likeReaction.querySelector('.reaction-count')
    like.innerText= parseInt(like.innerText)+1
    rating.className="reaction selected"
    console.log(rating)
}