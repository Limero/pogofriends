let index = 0

window.addEventListener('focus', updateFriendCode)
document.getElementById('next-code').addEventListener('click', updateFriendCode)

function updateFriendCode() {
    let friendCode = friendCodes[index]
    navigator.clipboard.writeText(friendCode).then(
        () => {
            document.getElementById('friend-code').innerHTML = friendCode
            index++
        },
    )
}
