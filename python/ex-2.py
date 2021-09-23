file = open("words.txt", "rt")
data = file.read()
words = data.split()

print('total number of words in text file :', len(words))
searchWord = input("input word to search: ")
if searchWord in words:
    print('Word found ',searchWord);
else:
    print('Not found')
