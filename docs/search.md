# Search API: execute a search query and get back search results

## Basics

Hoe internals of the search works can be understood in the way data it stored and how it is retrieved

- Getting Data In : Analyze the document for the index
	- tokenization (split on specific tokens like space)
	- lowercasing  (lowercase the document)
	- stop wording (words that can be ignored)
	- stemming (normalize the words in document. Mostly done internally )
	- indexing (prepare a score using TF * IDF)
	- forming a inverted index 
- Getting Data Out
	- Boolean query (performs boolean query on inverted index)
	- Sorting by relevance
		- Initialize priority queue
		- Iterate over matching docs
			- calculate ```TF * IDF``` score
			- add to priority queue
			- pop off lowest scoring docs
	- Aggregation
		- Initialize aggregator
		- Iterate
			- Add interesting results to aggregator
		- return results from aggregator




- Can be of below types:
	- Multi Index: Search for document in all indices or in some specific indices.
	- Multi type: Search for all documents in an index across all or specific types.
	- URI search: various param can be sent with uri for specific operation
- Single Document API: 
- Multi Document API: Multi Get, Bulk , Delete by Query , Update by Query , Reindex API

- Aggregation API: If you want to collect data and and work on the data set. Help to build complex summary data.
- Index API: Perform operation on index level : used to manage all aspects of index like settings, alises, mappings etc.
- Cluster API: Getting information of the cluster and its nodes

## Query DSL 
DSL based on JSON to define queries. It has 2 types of clauses:

- Leaf Query clauses: Look for particular value in (match,)
- Compound Query clauses: Combines leaf query clause with others to make a compund query
