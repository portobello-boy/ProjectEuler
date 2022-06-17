"""
Each character on a computer is assigned a unique code and the preferred standard is ASCII (American Standard Code for Information Interchange). For example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.

A modern encryption method is to take a text file, convert the bytes to ASCII, then XOR each byte with a given value, taken from a secret key. The advantage with the XOR function is that using the same encryption key on the cipher text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.

For unbreakable encryption, the key is the same length as the plain text message, and the key is made up of random bytes. The user would keep the encrypted message and the encryption key in different locations, and without both "halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key. If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the message. The balance for this method is using a sufficiently long password key for security, but short enough to be memorable.

Your task has been made easy, as the encryption key consists of three lower case characters. Using p059_cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes, and the knowledge that the plain text must contain common English words, decrypt the message and find the sum of the ASCII values in the original text.
"""

FILE = "cipher.txt"
COMMON_WORDS = ['the', 'and', 'that', 'have']

def readChars(filename):
    chars = []
    with open(filename) as f:
        for line in f:
            chars += line.strip().split(',')
        chars = [int(c) for c in chars]
    return chars

def containsCommonWords(chars):
    countWords = 0
    for word in COMMON_WORDS:
        found = False
        for i in range(len(chars) - len(word)):
            if ''.join(chars[i:i+len(word)]) == word:
                found = True
        if found:
            countWords += 1
    if countWords == len(COMMON_WORDS):
        return True
    return False

def xor(chars, key):
    xorList = []
    for i in range(len(chars)):
        xorList.append(chr(chars[i] ^ key[i%3]))
    return xorList

def decrypt(chars, key):
    decryption = xor(chars, key)
    if containsCommonWords(decryption):
        print("{} - {}".format(key, ''.join(decryption)))
        print(sum([ord(c) for c in decryption]))

chars = readChars(FILE)

for a in range(ord('a'), ord('z') + 1):
    for b in range(ord('a'), ord('z') + 1):
        for c in range(ord('a'), ord('z') + 1):
            key = [a,b,c]
            decrypt(chars, key)