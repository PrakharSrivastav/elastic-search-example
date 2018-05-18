# elastic-search-example

Elastic search. Basic usage and examples 


## Intention

The intention of this document is to quickly setup elasticsearch environment for developers. Their are instructions for setting up initial data and performing simple search requests using commandline. More language specific examples will be added soon.

## Setup and installation

- 1 Install docker image: ```docker pull docker.elastic.co/elasticsearch/elasticsearch-oss:6.0.1```
- 2 Run docker image: ```docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch-oss:6.0.1```
- 3 Check the elasticsearch is running ```curl -XGET localhost:9200```
- 4 We need to upload some data on elastic search. Download it using the command ```wget https://raw.githubusercontent.com/elastic/elasticsearch/master/docs/src/test/resources/accounts.json``` 
- 5 Import the data to the server. ```curl -H "Content-Type: application/json" -XPOST "localhost:9200/bank/_doc/_bulk?pretty&refresh" --data-binary "@accounts.json"```
- 6 Check if the data is imported properly ```curl -H "Content-Type: application/json" -XPOST "localhost:9200/bank/doc/_bulk?pretty&refresh" --data-binary "@accounts.json"```
 	The output should look like:
  
health | status | index | uuid | pri  | rep  | docs.count | docs.deleted | store.size | pri.store.size
:---   | :----- | :---- | :--- | :--- | :--- | :---       | :---         | :---       | :---  
yellow | open   | bank  | xGJ3Vcmw | 5 | 1 | 0 | 0 | 1.1kb | 1.1kb


