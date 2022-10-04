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

## Project Description Template

The purpose of this *Project Description* is to present the ideas proposed and decisions made during the preliminary envisioning and inception phase of the project. The goal is to analyze an initial concept proposal at a strategic level of detail and attain/compose an agreement between the project team members and the project customer (mentors and instructors) on the desired solution and overall project direction. You can also add images as you like.

This template proposal contains a number of sections, which you can edit/modify/add/delete/organize as you like.  Some key sections we’d like to have in the proposal are:

- Vision: An executive summary of the vision, goals, users, and general scope of the intended project.

- Solution Concept: the approach the project team will take to meet the business needs. This section also provides an overview of the architectural and technical designs made for implementing the project.

- Scope: the boundary of the solution defined by itemizing the intended features and functions in detail, determining what is out of scope, a release strategy and possibly the criteria by which the solution will be accepted by users and operations.

Project Proposal can be used during the follow-up analysis and design meetings to give context to efforts of more detailed technical specifications and plans. It provides a clear direction for the project team; outlines project goals, priorities, and constraints; and sets expectations.

** **

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

> This section describes the principal user roles of the project together with the key characteristics of these roles. This information will inform the design and the user scenarios. A complete set of roles helps in ensuring that high-level requirements can be identified in the product backlog.

> Again, the description should be specific enough that you can determine whether user A, performing action B, is a member of the set of users the project is designed for.

Users of our project are people who want to store their data/files in the cloud (FDB to be exact) without being leaked during server-side searches. Our users need a search scheme that is efficient while protecting search access patterns.

Specifically, when a user perform a search for a keyword over the files within a folder, the server learns only that a search happens in that folder, but does not learn which documents match the search, the number of documents that match, or other information about the keyword. 


** **

## 3.   Scope and Features Of The Project:

The project will have a Java interface to FDB. The project will implement the Clusion encrypted search algorithm with Java and FDB.The project will have a Go interface to FDB. The project will implement the Dory encrypted search algorithm with Go and FDB. 

The project scope includes testing the two encryption search algorithms on our local machines. We will also test the algorithms on MOC. We will have an analysis comparing the two algorithms in terms of their space, speed, and security. We will compare them using varying dataset sizes, number of parallel FDB queries, storage schemes, cluster sizes, and query complexities. 


** **

## 4. Solution Concept
Client loads the data into the FDB database which is then encripted and is readly avalible for parsing. <br />
                  ___________________ <br />
        (load)   {                    } <br />
client --------> {                    } <br />
                 {       FDB          } <br />
                 {                    } <br />
                 {                    } <br />
                 ---------------------- <br />
 Our job is to: 
 1) Load Files into FDB
 2) Update Files (delete, upload)
 3) Search Files
 
 We then have to evaluate the algorith for its:
 -Speed
 -Space
 -Security (leakage)

Design Implications and Discussion:

This section discusses the implications and reasons of the design decisions made during the global architecture design.

## 5. Acceptance criteria

By the end of the project, we as a team expect to test the encrypted search algorithms on MOC and deliver the following results:
* Compare speed of queries in regards to:
   * Size of dataset
   * Clusion vs Dory
   * Number of Parallel FBD Queries
   * Take size of cluster and query complexity into account

## 6.  Release Planning:

Load files (the dataset to search) into FDB 
Load index into FDB 
Update index with new files
Query for token, return files 

** **

## General comments

Remember that you can always add features at the end of the semester, but you can't go back in time and gain back time you spent on features that you couldn't complete.

** **

For more help on markdown, see
https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet

