"""
Write a  program to sort a list of tuples using Lambda.
Original list of tuples:
[('English', 88), ('Science', 90), ('Maths', 97), ('Social sciences', 82)]
Sorting the List of Tuples:
[('Social sciences', 82), ('English', 88), ('Science', 90), ('Maths', 97)]

"""

marks = [('English', 99), ('Science', 90), ('Maths', 97), ('Social sciences', 82)]

sort = sorted(marks, key=lambda x: x[1])

print(sort)
