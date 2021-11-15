# Xtern-Matcher

## Running

[Pre-compiled binaries are available for all major operating systems.](https://github.com/mhoc/xtern-matcher/releases/latest)

After downloading:

```
$ chmod +x DOWNLOADED_FILE
$ mv DOWNLOAD_FILE xternmatcher
$ ./xternmatcher -h
```

If you see help output, you're good to go.

## Compiling

To compile from scratch, at least `go1.11` is required.

```
$ git clone https://github.com/mhoc/xtern-matcher
$ go build
$ ./xtern-matcher -h
```

## Running

Input is provided via CSV files. These files can be exported from Google Sheets, or provided by any other source that can create CSVs. Two are required; one for company preferences and the other for student preferences.

The company preferences sheet should be formatted like the following. Column names should not be provided; we simply order each column like `Company Name`, `Desired Number of Interns`, then their rankings. If one of the student names is empty, or the value "Choose Your Candidate", it is ignored.

```
Google,3,Alice,Bob,Charlie
```

Students looks similar, lacking the number in the second column. If one of the company names is empty or "Choose a Company", it is ignored.

```
Alice,Google,Facebook,
Bob,Microsoft,,
Charlie,Yelp,Google,Square
```

Once we have these two files, simply pass them to the matcher:

```
$ ./xternmatcher -alg simple -in-companies ./companies.csv -in-students ./students.csv -out csv -pivot students
```

There are options available to change the output format and pivot view of the data; check out the `-h` for more info.
