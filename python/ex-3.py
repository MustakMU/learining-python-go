file = open("words.txt", "r")
data = file.read()
words = data.split(' ')
st = ''
st = [st + i for i in words]
print('sr' ,st)
wordToReplace = input("Enter word to replace ")
replacementWord = input("Enter replacement word ")
count = data.count(wordToReplace)
print('Found {} words in file '.format(count))
if count > 0:
    data = data.replace(wordToReplace, replacementWord)
    #file.truncate(0)
    print(data)
    with open("words.txt", "w") as w:
        w.write(data)
    #file.write(data)

    print('Replacemt done') 
else:
    print("No word found")