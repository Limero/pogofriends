import re
import requests

subReddit = 'PokemonGoFriends'
redditThread = 'wj4i3y'

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64; rv:81.0) Gecko/20100101 Firefox/81.0',
}

s = requests.Session()
s.headers = headers

r = s.get('https://reddit.com/r/'+subReddit+'/comments/'+redditThread+'.json')
body = r.text

regex = r'([0-9]{4}\s[0-9]{4}\s[0-9]{4})'
friendCodes = set(re.findall(regex, body))
print("Found", len(friendCodes), "friendcodes")

jsConst = 'const friendCodes = ['
for friendCode in friendCodes:
    jsConst += '"'+friendCode+'",'
jsConst += ']'

textfile = open("friendcodes.js", "w")
a = textfile.write(jsConst)
textfile.close()
