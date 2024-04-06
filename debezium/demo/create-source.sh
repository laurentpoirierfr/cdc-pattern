# Create connector
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:9090/connectors/ -d @pg-source-config.json

# Get connectors
curl -H "Accept:application/json" localhost:9090/connectors/


# Delete 
# curl -X DELETE localhost:9090/connectors/pg-demo-source