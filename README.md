# Xtern-Matcher

Installing and running (`go1.11` required):

```
$ go build
$ ./xtern-matcher -h
$ ./xtern-matcher -alg simple -in-companies ./companies.csv -in-students ./students.csv -out csv -pivot students
```

Input files must be CSVs formated like the following for companies:

```
{companyName},{desiredNumberOfInterns},{studentNameRank0},...,{studentNameRankN}
```

And like this for students:

```
{studentName},{companyNameRank0},...,{companyNameRankN}
```

