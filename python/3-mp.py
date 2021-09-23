import json

file = open("chat.json", "r")
data = file.read()
chatResponses = {}
chatResponses = json.loads(data)

print('Hi.. How i can help you?')

while True:
    cmd = input('\n').lower()
    if cmd == 'exit' or cmd == 'quit':
        print('Take care Bye!')
        break
    if len(cmd.split()) < 2:
        print('Please type atleast two words')
        continue

    if cmd in chatResponses:
        print(chatResponses[cmd]['reply'])
    else:
        indexes = {}
        for key in chatResponses:
            matches = len(set(cmd.split()).intersection(key.split()))
            indexes[key] = matches
        if len(indexes.keys()) == 0:
            print("Command not found! please try again")
        else:
            highestIndex = 0
            for key in indexes:
                if indexes[key] > highestIndex:
                    highestIndex = indexes[key]
            totalIndexes = list(indexes.values()).count(highestIndex)
            if totalIndexes == 1:
                for key in indexes:
                    if indexes[key] == highestIndex:
                        print(chatResponses[key]['reply'])
            else:
                print('Did you mean any of these?')
                for key in indexes:
                    if indexes[key] == highestIndex:
                        print(key)
            print()


