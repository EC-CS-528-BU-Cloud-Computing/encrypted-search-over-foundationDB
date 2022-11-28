# Encrypted Search Over FoundationDB

** **
## Team:
- Dasha Smolina (dsmolina@bu.edu)
- Jin San Yoon (jinsaaan@bu.edu)
- Stephen Sweet (ssweet@bu.edu)
- Tian Tan (tiant@bu.edu)
- Thomas Shen (zhaojun@bu.edu)

## Mentor:
- Dan Lambright (dlambrig@gmail.com)



## 1. Vision and Goals Of The Project:

FoundationDB (FDB) is a distributed key-store with strong ACID guarantees (Atomicity, Consistency, Isolation, Distributed). It is widely used to store users' documents in the cloud. The users want to store and query their documents in encrypted form since they do not trust the server. Preventing leakage of documents and query results is the purpose of the project.

The core vision of this project is to develop a search algorithm that can search documents on the cloud (in FDB) without leakage.

The project is divided into the following steps:
* Write Java/Go interface to FDB, which can:
    * Load files (the dataset to search) into FDB - done once.
    * Load index into FDB - done once.
    * Update index with new files - done frequently
* Learn and apply the two encrypted search algorithm to our project, which can
    * Query for token (in encrypted), return files - done frequently

The idea system can:
* make a query request to FDB by random bits (encrypted token)
* get a response from FDB by random bits (encrypted files)

If the project is successful, the open source project will hopefully become a safe service for users of FDB.

## 2. Users/Personas Of The Project:

Users of our project are people who want to store their data/files in the cloud (FDB to be exact) without being leaked during server-side searches. Our users need a search scheme that is efficient while protecting search access patterns.

Specifically, when a user perform a search for a keyword over the files within a folder, the server learns only that a search happens in that folder, but does not learn which documents match the search, the number of documents that match, or other information about the keyword. 

** **

## 3.   Scope and Features Of The Project:

The project will have a Java interface to FDB. The project will implement the Clusion encrypted search algorithm with Java and FDB.The project will have a Go interface to FDB. The project will implement the Dory encrypted search algorithm with Go and FDB.

The project scope includes testing the two encryption search algorithms on our local machines. We will also test the algorithms on MOC. We will have an analysis comparing the two algorithms in terms of their space, speed, and security. We will compare them using varying dataset sizes, number of parallel FDB queries, storage schemes, cluster sizes, and query complexities. 

** **

## 4. Solution Concept
Encripted search is becoming more necessary for security in this day and age.<br />
FDB works as follows: <br />
Client loads the data into the FDB database which is then encrypted and is readily available for parsing. <br />
<pre>  
1.    user loads a file (ex. apple)   ______________________________________
      Client -------------------->  |      Foundation Data Base (FDB)      |
                                    |                                      |
                                    |                                      |
                                    |______________________________________|
2.    The file gets encrypted and FDB becomes indexable
                                     ______________________________________
      Client -------------------->  |      Foundation Data Base (FDB)      |
                                    |           a  p                       |
                                    |            p l         l e           |
                                    |______________________________________|      
3. The file is stored in the FDB
</pre>
 Our job is to: 
 1) Load Files into FDB.
 2) Update Files (delete, upload)
 3) Search Files
 
 We then have to evaluate the algorithm for its:
 * Speed
 * Space
 * Security (leakage)

## 5. Acceptance criteria

By the end of the project, we as a team expect to test the encrypted search algorithms on MOC and deliver the following results:
* Compare speed of queries in regards to:
   * Size of dataset
   * Clusion vs Dory
   * Number of Parallel FBD Queries
   * Take size of cluster and query complexity into account

## 6.  Release Planning:

1) Load files (the dataset to search) into FDB 
2) Load index into FDB 
3) Update index with new files
4) Query for token, return files 

** **

## 7. Implement

* Java Implementation
  * Environment: Windows, JDK8, FDB-Java-API 5.2.5

** **

## General comments

Remember that you can always add features at the end of the semester, but you can't go back in time and gain back time you spent on features that you couldn't complete.

** **

For more help on markdown, see
https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet

