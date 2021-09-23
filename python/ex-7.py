
import pandas
df = pandas.read_csv('employee-det.csv')
for row in df.values:
    print(str(row[0])+' '+row[1]+' '+str(row[2]))