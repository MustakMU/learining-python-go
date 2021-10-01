
"""
Exercise - 1

1-Create employee data excel file with duplicate record  that contain the fallowing field
Emp id
Emp Name
Emp Salary
Write a python code to display duplicate employee record and count no of duplicate record available into a employee data file
"""


# using pandas library to parse csv file
import pandas

fileName = 'employee-det.csv'

df = pandas.read_csv(fileName)

# convert employee id coolumn to list
a = df['Employee id'].values.tolist()
# convert list to dictonary with key being employee id and value being total number of occurance of each key
result = dict((i, a.count(i)) for i in a)

totalDuplicates = 0;
duplicates = []

# iterate through the map
for keys in result.keys():
    # if number of occurance is more than one then the key is repeated more than one
    if result[keys] > 1:
        # store the duplicates and increase the total duplicates count
        duplicates.append(keys)
        totalDuplicates+=1

# print the result
print(result);
print('duplicates employee ids ',duplicates)
print('total duplicates ',totalDuplicates)