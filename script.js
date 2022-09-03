let index = 0
updateFriendCode()

window.addEventListener('focus', function (event) {
    updateFriendCode()
});

function updateFriendCode() {
    let friendCode = friendCodes[index]
    document.getElementById('friend-code').innerHTML = friendCode
    putInClipboard(friendCode)
    index++
}

function putInClipboard(s) {
    navigator.clipboard.writeText(s)
}
