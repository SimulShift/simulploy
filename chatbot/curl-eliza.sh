
curl \
    --header 'Content-Type: application/json' \
    --data '{"sentence": "I feel happy."}' \
    --http2 \
    http://docker-chatbot-prod-1:8083/connectrpc.eliza.v1.ElizaService/Say

curl -sSL "https://github.com/fullstorydev/grpcurl/releases/download/v1.8.7/grpcurl_1.8.7_linux_x86_64.tar.gz" | tar -xz -C /usr/local/bin

curl \
    --header 'Content-Type: application/json' \
    --data '{"sentence": "I feel happy."}' \
    --http2-prior-knowledge \
    http://docker-chatbot-prod-1:8083/connectrpc.eliza.v1.ElizaService/Say

curl \
    --header 'Content-Type: application/json' \
    --data '{"sentence": "I feel happy."}' \
    --http2-prior-knowledge \
    http://localhost:8083/connectrpc.eliza.v1.ElizaService/Say