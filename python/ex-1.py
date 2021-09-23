import pandas
df = pandas.read_csv('employee-det.csv')
print(df['Employee id'].values)
a = df['Employee id'].values.tolist()
result = dict((i, a.count(i)) for i in a)
totalDuplicates = 0;
duplicates = []
for keys in result.keys():
    if result[keys] > 1:
        duplicates.append(keys)
        totalDuplicates+=1

print(result);
print('duplicates employee ids ',duplicates)
print('total duplicates ',totalDuplicates)