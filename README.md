# elastic-search-example


## Basics concepts and terminologies:
- Open source search engine based on apache lucene
- Java, Cross platform
- Designed to take data from any source, analyze it and search through it.
- Features:
	- Near real time: Since any changes to the document needs propagation to all nodes and shards. It taks nearly a second to see the changes.
	- Clustered , Distributed across nodes. Highly scalable
	- Communication with search server is done via HTTP Rest api. ```curl -X <Http verb> <Node>:<Port>/<Index>/<Type>/<Id>```
	- Schema less Json document

### Document: 

- Basic unit of information that can be indexed.
- Consists of fields with a datatype.
- Similar to object in OOP, or a row in RDBMS.
- Expressed in JSON.
- No limit on the number of documents that can be stored in an index.
  
### Index :

- Collection of documents with similar feature (major catagory e.g movies, customer).
- Index is identified by a name. 
	- Name is used for performing indexing, search , update and delete.
	- Should be lowercase.
	- Can defina as many indexes as possible within a cluster.

### Type: Is a logical partition of an index. 

- Represents a class of documents within index that have several matching features.
- Consists of name and mapping.
- Similar to table within relational database.
- An index can have one or more types defined, each with their own mapping.
- MAPPING:
	- Similar to database schema.
	- describes fields and their datatypes.
	- also describes how the fields should be indexed.
	- DynamicMapping: means that its optional to define a mapping explicitly.

### Shards

- Index can be subdivided into shards.
- Useful if an index contains more data than hardware can support (eg 1TB data on 500GB disk)
- Default number of shard = 5.
- Allows to scale horizontally.
- Performance gains. Allows to distribute and parallelize operations.

### Replica

- Enables high avaliablity.
- Copy of a shard

## Api Conventions

### Api Types

#### Index APIs

- **Create** ```curl -XPUT /uri/<index_name> -d'{}'```
- **Delete** ```curl -XDELTE /uri/<index_name> ```

##### Mapping of queries to document:
Most of the times we need to define a minimal data structure for the documents being stored in an index. Mappings helps us to do that. Mapping is the process of defining how a document, and the fields that it contains are stored and indexed.

- Mapping Types: 
	- Meta
	- Field or properties
- Mapping Field datatypes:
	- Core datatypes:
	- Geo datatypes:
	- Complex datatypes:
	- Specialized datatypes (custom)

#### Document API: Operation on documents

- Get a document : GET /uri/id 
- Create a document : ```curl -XPUT /uri/index/type/id -d '{}'```
- Replacing a document : ```curl -XPUT /uri/index/type/id -d '{}'``` This will overwrite an existing doc.
- Updating a document: ```curl -XPOST /uri/index/type/id/_update -d '{}'```. Supply only updated values.
- Delete a document : ```curl -XDELETE /uri/index/type/id ```
	- Also possible to delete by query.

##### Bulk document apis

- To perform bulk operations _bulk apis can be used.
- A way to improve performance when working on bulk data.
- The bulkd actions are performed in sequential manner.
- If one of the actions fails, the subsequent actions will still be performed.

```curl -XPOST /uri/index/type/_bulk -d '{}'```

#### Search API: execute a search query and get back search results

- Can be of below types:
	- Multi Index: Search for document in all indices or in some specific indices.
	- Multi type: Search for all documents in an index across all or specific types.
	- URI search: various param can be sent with uri for specific operation
- Single Document API: 
- Multi Document API: Multi Get, Bulk , Delete by Query , Update by Query , Reindex API

- Aggregation API: If you want to collect data and and work on the data set. Help to build complex summary data.
- Index API: Perform operation on index level : used to manage all aspects of index like settings, alises, mappings etc.
- Cluster API: Getting information of the cluster and its nodes

#### Query DSL 
DSL based on JSON to define queries. It has 2 types of clauses:

- Leaf Query clauses: Look for particular value in (match,)
- Compound Query clauses: Combines leaf query clause with others to make a compund query


#### Modules is 

- Static
- Dynamic