let index = -1
updateFriendCode()

window.addEventListener('focus', function (event) {
    updateFriendCode()
})

document.getElementById('next-code').addEventListener('click', updateFriendCode)

function updateFriendCode() {
    index++
    let friendCode = friendCodes[index]
    document.getElementById('friend-code').innerHTML = friendCode
    putInClipboard(friendCode)
}

function putInClipboard(s) {
    navigator.clipboard.writeText(s)
}
