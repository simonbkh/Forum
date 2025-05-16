function updateReactionDisplay(reactionContainer, isActive, countChange) {
    var icon = reactionContainer.querySelector('.cmt-reaction, .post-reaction')
    var count = reactionContainer.querySelector('.reaction-count')

    if (icon && count) {
        icon.style.color = isActive ? '#2196F3' : '' // Set color to blue if active
        count.textContent = parseInt(count.textContent) + countChange // Update the reaction count
    }
}

function toggleReaction(contentType, contentId, reactionType) {
    var container = document.getElementById(contentId + (contentType === 'posts' ? '-container-for-post' : '-comment-reaction'))
    var reactionDivs = container ? container.querySelectorAll('.reaction, .comment-reaction') : null

    if (!container || !reactionDivs || reactionDivs.length < 2) {
        console.log('Container or reactions not found for:', contentId)
        return
    }

    var likeContainer = reactionDivs[0]
    var dislikeContainer = reactionDivs[1]

    var activeLike = likeContainer.querySelector('.post-reaction, .cmt-reaction').style.color === 'rgb(33, 150, 243)'
    var activeDislike = dislikeContainer.querySelector('.post-reaction, .cmt-reaction').style.color === 'rgb(33, 150, 243)'

    var action = (reactionType === 'like' && activeLike) || (reactionType === 'dislike' && activeDislike) ? 'remove' : 'add'

    var xhr = new XMLHttpRequest()
    xhr.open('POST', `/api/${contentType}/${contentId}/reaction`, true)
    xhr.setRequestHeader('Content-Type', 'application/json')

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var response = JSON.parse(xhr.responseText)

                if (!response.success) {
                    window.location.href = '/login'
                    return
                }

                if (reactionType === 'like') {
                    if (activeDislike) updateReactionDisplay(dislikeContainer, false, -1)
                    updateReactionDisplay(likeContainer, !activeLike, activeLike ? -1 : 1)
                } else {
                    if (activeLike) updateReactionDisplay(likeContainer, false, -1)
                    updateReactionDisplay(dislikeContainer, !activeDislike, activeDislike ? -1 : 1)
                }
            } else {
                console.error('Error with server response:', xhr.status)
                if (xhr.status != 400) {
                    window.location.replace("http://localhost:8080/login")
                } else {
                    document.body.innerHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Error Page</title>
  <style>
    .error-container {
      margin: 50px;
      position: relative;
      width: 700px;
      padding: 40px;
      background: #fff;
      border-radius: 20px;
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
      text-align: center;
      overflow: hidden;
      animation: slideIn 0.8s ease-out;
    }

    /* Animated Gradient Line */
    .error-container::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 5px;
      background: linear-gradient(90deg, rgba(0, 0, 0, .2), rgba(0, 0, 0, .2), #c9d6ff);
      animation: gradientMove 3s linear infinite;
    }


    body {
      display: flex;
      justify-content: center;
      align-items: center;
      min-height: 100vh;
      margin: 0;
      position: relative;
    }

    .error-code {
      font-size: 150px;
      font-weight: 800;
      margin: 0;
      color: #6A5ACD;
      background: -webkit-linear-gradient(#7494ec, #8a8c91);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
    }

    .error-message {
      font-size: 24px;
      color: #555;
      margin: 20px 0;
    }

    .btn {
      display: inline-block;
      padding: 12px 25px;
      margin: 10px 5px;
      font-size: 16px;
      font-weight: 600;
      text-decoration: none;
      color: #fff;
      border-radius: 50px;
      transition: all 0.3s ease;
    }

    .btn-primary {
      background:#7494ec;
    }

    .btn-primary:hover {
      background: #0056b3;
      transform: scale(1.05);
    }

    .btn-secondary {
      background: #6c757d;
    }

    .btn-secondary:hover {
      background: #5a6268;
      transform: scale(1.05);
    }

    /* Mobile Responsiveness */
    @media (max-width: 768px) {
      .error-container {
        width: 90%;
        padding: 20px;
      }

      .error-code {
        font-size: 100px;
      }

      .error-message {
        font-size: 18px;
      }
    }
  </style>
</head>
<body style="background: #fff;">
  <div class="error-container">
    <div class="error-code">400</div>
    <div class="error-message">BadRequest!!</div>
    <a href="/" class="btn btn-primary">Go to Home</a>
  </div>
</body>
</html>
`
                }
            }
        }
    }

    xhr.onerror = function () {
        console.error('Error making request')
    }

    var requestData = {

        type: reactionType,
        action: action,
        SessionToken: document.cookie
    }

    xhr.send(JSON.stringify(requestData))
}

document.addEventListener('DOMContentLoaded', function () {
    function setupReactionHandlers(container, contentType) {
        var contentId = contentType === 'posts' ? container.dataset.id : container.id.split('-')[0]
        var reactions = container.querySelectorAll('.post-reaction, .cmt-reaction')

        if (reactions.length >= 2) {
            reactions[0].onclick = function () { toggleReaction(contentType, contentId, 'like') }
            reactions[1].onclick = function () { toggleReaction(contentType, contentId, 'dislike') }
        }
    }

    document.querySelectorAll('.likes-section').forEach(container => setupReactionHandlers(container, 'posts'))
    document.querySelectorAll('[id$="-comment-reaction"]').forEach(container => setupReactionHandlers(container, 'comments'))
})
