
"""
Exercise - 2

2- Search and count the word available in file. 
"""

# read file content
file = open("words.txt", "rt")
data = file.read()

# split file data to list containing words
words = data.split()

print('total number of words in text file :', len(words))

# enter word to search
searchWord = input("input word to search: ")

# check if word is present in list
if searchWord in words:
    # if word is present then print word and total number of occurance
    print('Word found: ',searchWord," occurances: ",words.count(searchWord));
else:
    print('Word not found')
