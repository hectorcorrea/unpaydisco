A proof of concept to create a simple discovery layer on top of the
[Unpaywall](https://unpaywall.org/) data.

The idea is to import the Unpaywall data into Solr and allow users to run basic 
searches to find out what open access articles are available. Only open access articles
(as defined by [best_oa_location](http://unpaywall.org/data-format)) are indexed into Solr .

## Prerequisites
To use this program you need to have a running version of Solr and a copy of the Unpaywall data. 

A sample data file with 100 journal articles is provided under the `./data` folder to get started.

The default `settings.json` file assumes that you are running Solr on the default port (8983) and 
that you have a Solr core named `unpaydisco` where the data will be imported to. You can tweak 
these settings by editting the settings file.


## The Code
```
# Get the code 
git clone https://github.com/hectorcorrea/unpaydisco.git
cd unpaydisco
go get
go build

# Create the Solr core
solr create -c unpaydisco

# Import the sample data
./unpaydisco -settings settings.json -import ./data/first_100.json

# To run the discovery interface (then browse to http://localhost:9001)
./unpaydisco -settings settings.json
```

