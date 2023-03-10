function checkIfLiked(likeOrDislike) {
    console.log("test");
    emptyLike = document.getElementById('emptyLike');
    like = document.getElementById('like');

    if (likeOrDislike == 1) {
        like.style.display = "block"
        emptyLike.style.display = "none"
    } else {
        like.style.display = "none"
        emptyLike.style.display = "block"
    }
}