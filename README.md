# Comparisons between Ruby, PHP, and Go.

I'm putting together a little set of comparison demos to highlight performance differences between Ruby, PHP, and Go. My goal is to highlight where Go is a more optimized for handling certain tasks. 

There are 3 sets of code, solving 3 problems:

* textout
* calc
* password

## textout

* Grabs the content of a Wordpress Blog from mysql
* Writes it out to static HTML files

The individual tasks and test scripts can take 1 argument:

* __loopcount__ - the number of times to create the same set of documents 

Data comes from a Wordpress Database.  

## calc

* Grabs the content of a World Flight database from mysql
* Loops through all records and uses the Haversine formula to compute the distance between cities
* Writes it out to static HTML files

The individual tasks and test scripts can take 1 argument:

* __max__ - the number of records to retrieve from the database. 

Data comes from openflights.org.  Get a copy of the data from them to test: http://openflights.org/data.html

## password

* Accepts a password
* Tests it against password validation rules
* Includes a dictionary test

The individual tasks and test scripts can take 2 argument:

* __limit__ - the number of passwords to attempt.
* __method__ - 2 options *bruteforce* and *hash*. *Hash* is significantly faster in all three languages.  