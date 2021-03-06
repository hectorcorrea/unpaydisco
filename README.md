A proof of concept to create a simple discovery layer on top of the
[Unpaywall.org](https://unpaywall.org/) data.

The idea is to import the Unpaywall data into Solr and allow users to run basic
searches to find out what open access articles are available. Only open access articles
(as defined by [best_oa_location](http://unpaywall.org/data-format)) are indexed into Solr .


## Prerequisites
To use this program you need to have a running version of Solr and a copy of the Unpaywall data.

A sample data file with 100 journal articles is provided under the `./data` folder to get started.

The default `settings.json` file assumes that you are running Solr on the default port (8983) and that you have a Solr core named `unpaydisco` where the data will be imported to. You can tweak these settings by editting the settings file.


## Source Code
```
# Get the code
git clone https://github.com/hectorcorrea/unpaydisco.git
cd unpaydisco
go build -o unpaydisco ./cmd/web/.
```


## Solr core
To create the Solr core, add a catch all field (all_text) for searching,
and configure the copy-field directives to populate it issue the following commands:

```
solr create -c unpaydisco

curl -X POST -H 'Content-type:application/json' --data-binary '{
  "add-field":{
    "name":"all_text",
    "type":"text_en",
    "multiValued":true
  }
}' http://localhost:8983/solr/unpaydisco/schema

curl -X POST -H 'Content-type:application/json' --data-binary '{
  "add-copy-field":{ "source":"*_s", "dest":[ "all_text" ]},
  "add-copy-field":{ "source":"*_ss", "dest":[ "all_text" ]},
  "add-copy-field":{ "source":"*_txt_en", "dest":[ "all_text" ]}
}' http://localhost:8983/solr/unpaydisco/schema
```


# Running it
To import the sample data

```
./unpaydisco -settings settings.json -import ./data/first_100.json
```

To run the discovery interface

```
./unpaydisco -settings settings.json
```

and then then browse to http://localhost:9001. Once it's running it should look more or less like [this screenshot](https://github.com/hectorcorrea/unpaydisco/blob/master/misc/search_results.png).


## Unpaywall data
This code is mean to work with the [database snapshot](http://unpaywall.org/products/snapshot) that Unpaywall.org provides. Unpaywall.org provides a JSON file in which each line represents an article.

I've tested this program with the download file `unpaywall_snapshot_2018-09-24T232615.json` that contains a total of 99,940,229 articles, 24,977,220 of which were marked as Open Access. Importing the Open Access articles results in a Solr index just shy of 13 GB.