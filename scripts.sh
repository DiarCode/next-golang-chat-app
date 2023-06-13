docker swarm join --token SWMTKN-1-0r13g86yhd7lyq667xcnen5b1vceq595v75zss9z8youczc54n-34svqkypa9umg456fs1humsag 192.168.65.4:237

docker build -t auth-service ./backend/services/auth
docker build -t chat-service ./backend/services/chat
docker build -t posts-service ./backend/services/posts
docker build -t gateway ./backend/gateway
docker build -t frontend-client ./frontend